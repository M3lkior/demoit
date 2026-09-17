/*
Copyright 2018 Google LLC
Copyright 2022 David Gageot

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/dgageot/demoit/files"
	"github.com/dgageot/demoit/flags"
	"github.com/dgageot/demoit/handlers"
	"github.com/dgageot/demoit/livereload"
	"github.com/dgageot/demoit/shell"
	"github.com/gorilla/mux"
	"github.com/rjeczalik/notify"
)

func main() {
	flags.DevMode = flag.Bool("dev", false, "dev mode with live reload")
	flags.WebServerPort = flag.Int("port", 8888, "presentation port")
	flags.WebServerHost = flag.String("host", "localhost", "host to bind the presentation server")
	flags.ShellPort = flag.Int("shellport", 9999, "shell server port (terminal)")
	flags.Locale = flag.String("locale", "", "locale suffix (e.g. en, fr) to select demoit-<locale>.html")
	flag.Parse()
	if args := flag.Args(); len(args) > 0 {
		files.Root = args[0]
	}

	if err := handlers.VerifyConfiguration(); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/{id:[0-9]*}", handlers.Step).Methods("GET")
	r.HandleFunc("/last", handlers.LastStep).Methods("GET")
	r.PathPrefix("/sourceCode/").HandlerFunc(handlers.Code).Methods("GET")
	r.HandleFunc("/shell/", handlers.Shell).Methods("GET")
	r.HandleFunc("/shell/{folder}", handlers.Shell).Methods("GET")
	r.PathPrefix("/ping").HandlerFunc(handlers.Ping).Methods("HEAD", "GET")
	r.PathPrefix("/js/").HandlerFunc(handlers.Static).Methods("GET")
	r.PathPrefix("/fonts/").HandlerFunc(handlers.Static).Methods("GET")
	r.PathPrefix("/images/").HandlerFunc(handlers.Static).Methods("GET")
	r.PathPrefix("/media/").HandlerFunc(handlers.Static).Methods("GET")
	r.HandleFunc("/demoit.css", handlers.EngineCSS).Methods("GET")
	r.HandleFunc("/tailwind.css", handlers.Static).Methods("GET")
	r.HandleFunc("/style.css", handlers.Static).Methods("GET")
	r.HandleFunc("/favicon.ico", handlers.Static).Methods("GET")
	r.HandleFunc("/qrcode", handlers.QRCode).Methods("GET")
	r.HandleFunc("/pdf", handlers.ExportToPDF).Methods("GET")
	r.HandleFunc("/speakernotes", handlers.SpeakerNotes).Methods("GET")
	r.HandleFunc("/grid", handlers.Grid).Methods("GET")
	r.HandleFunc("/beta/vscode/{folder}", handlers.VSCode).Methods("GET")

	// Reverse Proxy Shell Server
	proxy := httputil.NewSingleHostReverseProxy(mustParseURL(fmt.Sprintf("http://127.0.0.1:%d", *flags.ShellPort)))
	r.PathPrefix("/tty").HandlerFunc(proxy.ServeHTTP)

	// Live Reload Server
	if *flags.DevMode {
		watchCSS(files.Root)

		lr := livereload.New(*flags.WebServerPort)
		lr.RegisterHandlers(r)

		events := make(chan notify.EventInfo, 1)
		if err := notify.Watch(files.Root+"/...", events, notify.All); err != nil {
			log.Fatal(err)
		}

		go func() {
			for event := range events {
				// TODO: Ignore files under .git
				// TODO: Debounce
				fmt.Println(event)
				lr.Reload(event.Path())
			}
		}()
	} else {
		fmt.Println(`"Dev Mode" to live reload your slides can be enabled with '--dev'`)
	}

	go func() {
		port := *flags.ShellPort
		log.Fatal(shell.ListenAndServe(port, "sh", "-c"))
	}()

	addr := flags.WebServerAddress()
	fmt.Println("Welcome to DemoIt. Please, open http://" + addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

// watchCSS starts `hack/css.sh <talk> --watch` beside the server, so that a
// Tailwind class written in a slide exists by the time live reload shows the
// page again. Without it the two halves of a change land at different times:
// the slide reloads immediately, the class it uses does not exist yet, and
// nothing on screen says why -- the browser simply applies a rule it does not
// have, which looks like the class having no effect.
//
// Dev mode only, and only when both halves are actually there. A binary put
// somewhere by `go install` has no repository around it, and a folder outside
// this tree has no tailwind.src.css to build; both are silent no-ops, because
// serving a deck has never required any of this and a warning on every start
// would be noise.
//
// A Ctrl+C at the terminal already reaches the child, which shares demoit's
// process group -- but only that. A `kill` on the server alone, or any other
// signal, would leave a watcher behind rebuilding a stylesheet for a talk
// nobody is showing, and the next run would start a second one. Hence the
// explicit handler: it is the difference between a watcher that stops the way
// the server does and one that accumulates.
func watchCSS(folder string) {
	const script = "hack/css.sh"

	if _, err := os.Stat(script); err != nil {
		return
	}
	if _, err := os.Stat(filepath.Join(folder, ".demoit", "tailwind.src.css")); err != nil {
		return
	}

	cmd := exec.Command(script, folder, "--watch")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		// Not fatal: a deck still renders with the stylesheet as it was
		// committed, which is exactly what a non-dev run gets.
		fmt.Println("Unable to watch the Tailwind stylesheet:", err)

		return
	}

	stopping := make(chan os.Signal, 1)
	signal.Notify(stopping, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-stopping
		_ = cmd.Process.Kill()

		// Re-raise on the default handler so demoit exits the way it would
		// have without us: installing a handler is what suppressed that.
		signal.Stop(stopping)
		if p, err := os.FindProcess(os.Getpid()); err == nil {
			_ = p.Signal(sig)
		}
	}()

	fmt.Println("Watching " + folder + "'s Tailwind stylesheet for new classes.")
}

func mustParseURL(rawURL string) *url.URL {
	url, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	return url
}
