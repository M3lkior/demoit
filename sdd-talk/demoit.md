---
layout: cover
class: responsive max center-align title
speakernotes: |
  S1 - Titre + toi. Qui je suis, pourquoi je parle de ça, la promesse de la
  demi-journée.
  Annoncer le format : 45 minutes d'introduction, puis atelier.
  Thèse unique à poser dès maintenant : l'IA a rendu le code abondant, le
  goulot s'est déplacé en amont (l'intention) et en aval (la vérification).
---

<h1 class="medium no-padding center-align top-align" style="color: white">Le développeur Product Engineer</h1>
<h3 style="color: white">Le développement piloté par les spécifications</h3>
<h5 style="color: white">Ludovic Dussart, Zatsit</h5>

---
layout: default
title: Je suis **Ludovic Dussart**
class: main responsive xlarge-height center-align
---
<div class="grid large-space">
  <div class="m6 s6 l6">
    <div class="round extra">
      <img style="block-size: 15rem; " src="/images/me.jpg">
    </div>
  </div>
  <div class="m6 s6 l6">
    <div class="center-align middle">
      <h4>Solutions architect <em>@zatsit</em></h4>
      <h4><em>AsyncAPI</em> maintainer</h4>
      <h4>Open Source <em>fanatic</em></h4>
    </div>
  </div>
  <div class="s12"><h2 class="left-align"><em>Où me trouver ?</em></h2></div>
  <div class="s1"><img class="circle " src="/images/twitter.jpg"/></div>
  <div class="s3 left-align"><h5>@ldussart</h5></div>
  <div class="s1"><img class="circle " src="/images/bsky.png"/></div>
  <div class="s3 left-align"><h5>ldussart.bsky.social</h5></div>
  <div class="s1"><img class="circle " src="/images/linkedin.png"/></div>
  <div class="s3 left-align"><h5>Ludovic Dussart</h5></div>
  <div class="s1"></div>
  <div class="s3"></div>
  <div class="s12"><h2 class="left-align"><em>Où nous trouver ?</em></h2></div>
  <div class="s1"><img class="circle " src="/images/web.png"/></div>
  <div class="s3 left-align"><h5>https://zatsit.fr</h5></div>
  <div class="s1"><img class="circle " src="/images/bsky.png"/></div>
  <div class="s3 left-align"><h5>zatsit.bsky.social</h5></div>
  <div class="s1"><img class="circle " src="/images/blog.png"/></div>
  <div class="s3 left-align"><h5>https://blog.zatsit.fr</h5></div>
</div>
<div class="medium-space"></div>
<div class="s12">
  <h6>« Engager notre <em>expertise numérique</em> au service de <em>l'impact des entreprises</em>, en créant un écosystème <em>durable</em>, <em>partenarial</em> et <em>positif</em> »</h6>
</div>


---
layout: default
title: Disclamer
---

<blockquote class="charniere">
  <h4>L'écosystème IA évolue bien trop vite pour que l'humain puisse le maitriser.</h4>
</blockquote>
---
layout: bare
class: transition
speakernotes: |
  Transition partie 0 - Un nouveau métier, et son problème. 7 minutes, S2 à S5.
  Enchaîner vite, 5 secondes à l'écran : la transition sert de repère à la
  salle, pas de temps de parole.
  Ce qu'on promet ici et qu'il faudra tenir : on part du personnage, on montre
  son problème, et le problème débouche sur le SDD en S5.
---

<h2>Le <b>Product Engineer<b></h2>
<h3>Nouveau métier ou évolution d'un métier existant ?</h3>

---
layout: default
title: "Product Engineer : le métier que le marché a inventé deux fois"
class: main responsive large-height
sources:
  - posthog.com/product-engineer
  - posthog.com/handbook/company/small-teams
  - productengineer.org
speakernotes: |
  - **PostHog** : la référence documentaire, avec un *Product Engineer Handbook* public. Leur définition : un développeur directement responsable de l'amélioration du produit, ayant un avis sur la roadmap, capable de prendre seul des décisions de design, faisant du support utilisateur directement, moins attaché aux fonctionnalités qu'il a ajoutées par le passé, conscient de la place de son travail dans la stratégie, et motivé avant tout par l'impact et les résultats. Modèle organisationnel associé : une petite équipe efficace doit avoir un seul leader, pouvoir livrer et décider de façon autonome avec un minimum de dépendances, et posséder sa propre mission, ses objectifs long terme, ses métriques clés et ses clients cibles. **Chiffre à mettre en slide : 187 personnes organisées en 47 équipes de 4,2 personnes en moyenne.**
  - **incident.io** : résume les product engineers comme des développeurs qui se soucient davantage des résultats et de l'impact que de l'implémentation exacte ou des outils utilisés.
  - **Jean-Michel Lemieux (ex-VP Engineering, Shopify)** : des ingénieurs qui ont soif d'utiliser la technologie pour court-circuiter les problèmes humains et utilisateurs.
  - **Product Engineer Manifesto** (productengineer.org / github.com/anttiviljami/product-engineer-manifesto, par Viljami Kuosmanen, ~2024) : c'est la responsabilité des builders de chercher d'abord à comprendre le problème avant de plonger dans les solutions, et de s'occuper des domaines design, technique et business en prenant une part active dans chacun.
  - **En France, le rôle existait dès 2024** : Le Talent Club a publié une interview d'une product engineer chez Flowie, arrivée par la voie PM. Son avertissement, à reprendre tel quel : beaucoup de gens qui veulent faire du produit veulent faire de la stratégie et prendre des décisions, alors qu'au quotidien c'est surtout de la delivery ; il ne faut pas devenir product engineer pour faire de la stratégie, ce ne sera pas la mission principale.
  Sources : PostHog Product Engineer Handbook (posthog.com/product-engineer) ;
  Product Engineer Manifesto (productengineer.org) ; Le Talent Club.
---

##### *2018 &ndash; 2024* &nbsp;&middot;&nbsp;

<p class="xlarge-text">Né dans des scale-ups produit, pour une raison
d'organisation : le <em>coût des relais</em> dépasse leur bénéfice quand
l'équipe est petite et que la boucle de feedback doit rester courte.</p>

<ul class="xlarge-text">
  <li><strong>PostHog</strong>: un développeur directement responsable de l'amélioration du produit, ayant un avis sur la roadmap, capable de prendre seul des décisions de design, faisant du support utilisateur directement, moins attaché aux fonctionnalités qu'il a ajoutées par le passé, conscient de la place de son travail dans la stratégie, et motivé avant tout par l'impact et les résultats. 187 personnes pour 47 équipes (4,2 personnes en moyenne)</li>
  <li><strong>incident.io</strong> : des devs attachés au résultat et à l'impact, pas à l'implémentation</li>
  <li><strong>JM. Lemieux</strong>, ex-VP Eng Shopify : des ingénieurs qui ont soif d'utiliser la technologie pour court-circuiter les problèmes humains et utilisateurs.</li>
  <li><strong>Product Engineer Manifesto</strong>, 2024 : c'est la responsabilité des builders de chercher d'abord à <em>comprendre le problème</em> avant de plonger dans les solutions, <em>et de s'occuper des domaines design, technique et business en prenant une part active dans chacun.</em></li>
</ul>

---
layout: default
title: "Product Engineer : le métier que le marché a inventé deux fois"
class: main responsive large-height
sources:
  - sfeir.com/concepts/product-engineer
  - turingcollege.com/blog/rise-of-the-ai-product-engineer
speakernotes: |
    Punchline : ce rôle n'a pas été créé par l'IA, il a été révélé par elle.

    - **Le "one-person band"** : le Chief AI Officer de Pendo décrit le glissement : traditionnellement un ingénieur livre, un product manager parle aux utilisateurs pour voir si c'était viable, et la boucle continue ; en se rapprochant du product engineer, on obtient une seule personne qui livre, itère, collecte le feedback et itère de nouveau, au lieu de répartir ça entre engineering, produit et design. Et parce qu'ils passent moins de temps à écrire du code grâce à Codex, Claude Code ou l'agent de leur choix, ce rôle se concentre sur le travail à plus forte valeur : la résolution créative de problèmes et les échanges avec les utilisateurs.
    - **L'argument économique (le plus solide)** : la ressource rare n'est plus l'exécution, c'est le jugement : savoir quoi construire, et savoir si ce qu'on a construit est bon. La rémunération monte pour ceux qui construisent concrètement et baisse pour ceux qui ne font que coordonner. Le rôle se situe à l'intersection de quatre disciplines : ingénierie logicielle/IA, product management, UX/UI et jugement sur la donnée.
---

:::grid{class="grid large-height"}

:::col{class="s12 center-align"} 
:::
:::col{class="s12 center-align"}
#### Ce rôle n'a pas été créé par l'IA. Il a été *révélé* par elle.
:::

:::col{class="s12"}

##### *2025 &ndash; 2026*

<p class="xlarge-text">Le même rôle passe de
<em>niche culturelle</em> à <em>norme</em>, parce que l'IA crée des goulots en amont des phases de production de codes, et plus seulement dans les petites équipes.</p>

<ul class="xlarge-text">
  <li>SFEIR positionne le product engineer comme le rôle central <em>de l'ère 10x</em>, qui ne se définit plus par une spécialité technique unique mais par sa capacité à coordonner toute la chaîne de production logicielle augmentée par l'IA.</li>
 
</ul>
:::


:::

---
layout: default
title: "Pourquoi maintenant : le mécanisme"
sources:
  - turingcollege.com/blog/rise-of-the-ai-product-engineer
  - sfeir.com/concepts/product-engineer
speakernotes: |
  S3 - La spécialisation était une réponse à la complexité. Quand l'IA absorbe
  une partie de la complexité d'exécution, la division du travail perd sa
  justification : l'IA élimine la distance entre les disciplines.
  Illustrer avec le "one-person band" du Chief AI Officer de Pendo : au lieu de
  répartir entre engineering, produit et design, une seule personne livre,
  itère, collecte le feedback et itère de nouveau.
  Puis l'argument économique, le plus solide : la ressource rare n'est plus
  l'exécution, c'est le jugement.
  Chiffre optionnel, à manier avec précaution (petits volumes, Angleterre
  seulement) : 72 offres "Product Engineer" sur 6 mois au 5 janvier 2026,
  contre 20 un an avant. Directionnel, pas probant - le dire. Non revérifiable
  le 11/09/2026 (la page itjobswatch ne publie plus cette fenêtre) : le garder
  à l'oral en le donnant pour ce qu'il est, ou le sauter.
  Sources : Turing College, The Rise of the AI Product Engineer ; SFEIR, page
  concept Product Engineer ; itjobswatch.co.uk (Product Engineer, England).
---

:::grid{class="grid large-height center-align"}
:::col{class="s4"}

##### La spécialisation était une réponse à la *complexité*.

:::

:::col{class="s4"}

##### L'IA absorbe une partie de la complexité d'*exécution*.

:::

:::col{class="s4"}

##### Donc la *division du travail* perd sa raison d'être.

:::

:::col{class="s12"}

<div class="small-space"></div>

<blockquote class="left-align">
  <h5>Avant : un ingénieur livre, un PM va voir les utilisateurs, la boucle continue.</h5>
  <h5>Maintenant : <em>une seule personne</em> livre, itère, collecte le feedback, et itère de nouveau.</h5>
</blockquote>

:::

:::col{class="s12"}

#### La ressource rare n'est plus l'exécution. C'est le *jugement*.

:::
:::

---
layout: default
title: Le paradoxe de productivité
sources:
  - faros.ai/blog/lab-vs-reality-ai-productivity-study-findings
  - linearb.io/resources/engineering-benchmarks
  - metr.org/blog/2026-02-24-uplift-update
  - dora.dev/ai/roi/report
speakernotes: |
  S4 - Le personnage est posé, on montre son problème. Trois colonnes, trois
  sources indépendantes qui disent la même chose, puis la conclusion.
  Faros AI, télémétrie sur plus de 10 000 développeurs et 1 255 équipes : à
  l'échelle individuelle +21 % de tâches complétées et +98 % de PR mergées,
  mais aussi +91 % de temps de revue, +154 % de taille de PR, +9 % de bugs -
  et au niveau de l'organisation, des métriques DORA (fréquence de déploiement,
  lead time, change failure rate) sans amélioration mesurable.
  LinearB, 8,1 millions de PR sur 4 800 équipes : le code généré par IA attend
  4,6 fois plus longtemps sa première revue ; 67,3 % de ces PR échouent à la
  première revue contre 15,6 % pour le code humain. Nuance honnête, à dire si
  on me la sort : une fois prise en main, la revue est 2 fois plus rapide. Le
  problème est la file d'attente, pas la lecture.
  METR, 16 développeurs expérimentés sur 246 issues réelles de leurs propres
  dépôts : 19 % de RALENTISSEMENT mesuré, alors qu'ils s'attendaient à 24 % de
  gain et croyaient encore après coup avoir gagné 20 %. C'est ce gap
  perception/mesure qui est le vrai résultat, et c'est lui qui parle à la
  salle.
  ATTENTION - correction du brief : le suivi METR de 2026 n'a PAS inversé le
  résultat en accélération de 18 %. Le point estimé reste à environ 18 % de
  ralentissement, et METR a annoncé revoir le design de l'étude à cause
  d'effets de sélection. Ne pas raconter de "courbe en J" : c'est faux, et
  quelqu'un dans la salle l'aura lu.
  Conclusion : l'IA ne répare pas une équipe, elle amplifie ce qui est déjà là.
  Sources : Faros AI (télémétrie 10 000+ devs) ; LinearB, benchmarks 2026 ;
  METR (metr.org, étude 2025 + mise à jour 2026-02-24) ; DORA 2026, The ROI of
  AI-assisted Software Development (Google) ; arXiv 2609.00252 pour le cadrage
  académique du paradoxe.
---

:::grid{class="grid large-height center-align"}
:::col{class="s4"}

<h3><em>+21 %</em></h3>

<p class="xlarge-text">de tâches complétées<br/>et <em>+98 %</em> de PR mergées</p>

<p class="xlarge-text">&hellip; mais <strong>+91 %</strong> de temps de revue,
et des métriques de delivery <strong>plates</strong> au niveau de
l'organisation.</p>

<h6>Faros AI : 10 000+ devs, 1 255 équipes</h6>

:::

:::col{class="s4"}

<h3><em>4,6&times;</em></h3>

<p class="xlarge-text">plus d'attente<br/>avant la première revue</p>

<p class="xlarge-text"><strong>67 %</strong> des PR générées par IA échouent à
la première revue, contre <strong>16 %</strong> du code humain.</p>

<h6>LinearB : 8,1 millions de PR</h6>

:::

:::col{class="s4"}

<h3><em>&minus;19 %</em></h3>

<p class="xlarge-text">de vitesse <strong>mesurée</strong><br/>sur des devs
expérimentés</p>

<p class="xlarge-text">Les mêmes devs se croyaient <strong>20 % plus
rapides</strong>. L'écart entre le ressenti et la mesure est le vrai
résultat.</p>

<h6>METR : 16 devs, 246 issues réelles</h6>

:::

:::col{class="s12"}

#### L'IA ne répare pas une équipe. Elle *amplifie* ce qui est déjà là.

:::
:::

---
layout: default
title: Le vrai goulot n'est pas le code
class: main responsive large-height
sources:
  - dora.dev/ai/roi/report
  - hitechnology.io
speakernotes: |
  S5 - LA charnière du talk, avec S30. Trois temps, dans cet ordre.
  1. Le diagnostic : le temps de développement est rarement le vrai frein. Les
  goulots sont en amont, dans des exigences floues et des décisions produit
  faibles. Une IA qui rend le code moins cher ne répare rien de ça, elle
  l'expose : la partie du process qui absorbait le mou avance désormais plus
  vite que tout ce qui l'entoure.
  2. La phrase charnière. La lire à voix haute, lentement, et marquer un temps.
  C'est le seul moment du talk où le lien product engineer / SDD se formule en
  une phrase. Elle revient à l'identique en S30.
  3. Le plan en 4 temps, à l'oral uniquement, rien à l'écran : les prérequis,
  le SDD et son intérêt, le panorama des frameworks, découpler et accélérer.
  20 secondes, pas plus, puis enchaîner.
  Corollaire à dire, pas à écrire : l'ingénieur le plus utile n'est pas celui
  qui produit le plus de code, c'est celui qui opère bien sur tout le chemin
  d'un problème métier à un changement livré et fonctionnel.
  Source du diagnostic : hitechnology.io, The Rise of the Product Engineer
  (étude 100+ leaders) ; DORA 2026 (Google).
---

<div class="large-space"></div>

### Le temps de développement est rarement le vrai frein.

##### Les goulots sont en amont : des exigences floues, des décisions produit faibles.

<div class="small-space"></div>

#### Une IA qui rend le code moins cher ne répare rien de ça. Elle l'*expose*.

<div class="large-space"></div>

<blockquote class="charniere">
  <h4>Sans SDD, un product engineer n'est qu'un dev surchargé à qui on a retiré son PM.</h4>
  <h5>Il a besoin d'un artefact pour porter l'intention jusqu'à la machine. Cet artefact, c'est la <em>spec</em>.</h5>
</blockquote>

---
layout: bare
class: transition
speakernotes: |
  Transition partie 1 - Les prérequis. 8 minutes, S6 à S10.
  Elle arrive juste après la charnière S5 : laisser retomber la phrase avant
  d'afficher celle-ci.
  Ce qu'on promet : avant de parler SDD, quatre choses à avoir en tête -
  où en est chacun dans la délégation, comment marche vraiment un agent,
  pourquoi le vibe coding casse, et pourquoi le DDD revient.
---

<h6 class="numero">Partie 1</h6>

<h2>Avant le SDD</h2>

<h5>Ce qu'il faut avoir en tête pour que la suite <em>tienne</em></h5>

---
layout: default
title: L'échelle de délégation
sources:
  - arxiv.org/abs/2609.00252
speakernotes: |
  S6 - Autocomplétion, chat, agent, agent autonome.
  Le secteur passe de pratiques assistées comme le vibe coding, où l'assistant
  accélère un développeur isolé, vers l'Agentic Software Engineering où des
  agents autonomes reçoivent des tâches au niveau de l'objectif.
  Faire lever la main : qui est où aujourd'hui ? Ça calibre la salle pour toute
  la suite. Si la salle est majoritairement en bas de l'échelle, prévoir de
  transférer 3 minutes de la partie 3 vers les parties 1 et 2.
  Le papier propose une typologie de cinq patterns d'interaction humain-agent
  par lesquels le rôle humain est redéfini : l'échelle du slide est une
  simplification pédagogique de ce cadre, pas une citation.
  Précaution à porter si on me cite le papier : il se présente lui-même comme
  un premier pas vers un consensus, pas comme une théorie validée.
  Source : arXiv 2609.00252, Spec-Driven Development for Agentic Software
  Engineering: Harnessing Human-Agent Teamwork.
---

:::grid{class="grid large-height center-align"}

:::col{class="s3"}

##### 1. Autocomplétion

<p class="xlarge-text">Tu écris, l'outil finit la ligne.</p>

<h6>Tu délègues la frappe.</h6>

:::

:::col{class="s3"}

##### 2. Chat

<p class="xlarge-text">Tu décris une fonction, tu relis, tu colles.</p>

<h6>Tu délègues un bout de code.</h6>

:::

:::col{class="s3"}

##### 3. Agent

<p class="xlarge-text">Tu donnes une tâche. Il lit le dépôt, édite, lance les tests.</p>

<h6>Tu délègues une tâche.</h6>

:::

:::col{class="s3"}

##### 4. Agent autonome

<p class="xlarge-text">Tu donnes un objectif. Il découpe le travail lui-même.</p>

<h6>Tu délègues l'intention.</h6>

:::

:::col{class="s12"}

<div class="small-space"></div>

#### Plus on monte, moins l'intention tient dans ta tête. Il faut l'*écrire*.

<h5>Vous êtes où, aujourd'hui ?</h5>

:::
:::

---
layout: default
title: Comment marche vraiment un agent
sources:
  - arxiv.org/abs/2604.08224
  - arxiv.org/abs/2605.01160
speakernotes: |
  S7 - Le slide le plus utile du talk. Si celui-là passe, le reste coule.
  Pas de mémoire persistante, seulement une fenêtre de contexte.
  Tout le SDD découle de là : puisque les LLM n'ont pas de mémoire persistante,
  il faut créer des artefacts durables et versionnés (constitution.md, spec.md,
  plan.md) pour externaliser l'état du projet.
  Introduire le terme context engineering ici : traiter la fenêtre comme une
  ressource rare, et ne dépenser le budget d'attention que sur des tokens à
  fort signal.
  La nuance à ne pas sauter, c'est elle qui évite la caricature : être présent
  dans la fenêtre ne garantit pas d'être retrouvé. Un prompt long n'est donc
  pas une mémoire durable.
  Lien explicite à faire avec S4, trois slides plus tôt : la contrainte de
  fenêtre de contexte est l'un des deux mécanismes qui amplifient le paradoxe
  de productivité, l'autre étant le goulot de la revue de code.
  Prendre le temps. C'est le slide à ne pas accélérer même en retard.
  Sources : l'image RAM / disque du context engineering ; arXiv 2604.08224
  (revue sur l'externalisation dans les agents LLM : mémoire, skills,
  protocoles, harness) pour la non-fiabilité de la fenêtre comme support
  d'état ; arXiv 2605.01160 pour la contrainte de fenêtre de contexte comme
  amplificateur du paradoxe de productivité.
---

:::grid{class="grid large-height center-align"}

:::col{class="s6"}

##### La fenêtre de contexte, c'est de la *RAM*

<p class="xlarge-text">Petite, rapide, et vidée à la fin de la session. Tout ce que l'agent sait de ton projet tient là-dedans, et seulement là.</p>

:::

:::col{class="s6"}

##### Ton dépôt, c'est le *disque*

<p class="xlarge-text">Grand, lent, persistant. L'agent n'y lit que ce qu'on lui dit d'aller chercher, et il repart de zéro à chaque fois.</p>

:::

:::col{class="s12"}

<blockquote class="left-align">
  <h5>Et dans la fenêtre, rien n'est garanti : être présent ne veut pas dire être retrouvé.</h5>
  <h5>Un prompt très long n'est pas une mémoire.</h5>
</blockquote>

:::

:::col{class="s12"}

#### Pas de mémoire persistante ? Alors on l'écrit.

<p class="xlarge-text"><strong>constitution.md</strong>, <strong>spec.md</strong>,
<strong>plan.md</strong> : des fichiers durables et versionnés, relus à chaque
session. C'est tout le context engineering, et c'est de là que sort le SDD.</p>

:::
:::

---
layout: default
title: Les modes de défaillance du vibe coding
sources:
  - arxiv.org/abs/2609.00252
speakernotes: |
  S8 - Concret, ils l'ont tous vécu. Raconter, ne pas lister.
  L'agent produit mille lignes qui ont l'air correctes et qu'on n'a pas
  demandées. On voulait un petit fix, on récupère un refactor, une logique
  "améliorée", trois nouveaux fichiers, et des tests qui passent parce qu'ils
  ne testent rien.
  Trois causes racines : l'intention ne vit que dans l'historique de chat ; la
  dérive ; l'output non vérifiable faute de critères d'acceptation.
  Ces trois causes annoncent les trois réponses du SDD - le dire.
  La formulation académique, à garder pour la question qui viendra : le SDD
  reconstitue, sous une forme centrée sur la spécification, les trois contrats
  que le vibe coding dissout, la responsabilité (accountability), la
  vérifiabilité et la transférabilité. Raconter la scène d'abord, les mots
  savants ensuite, jamais l'inverse.
  Source : arXiv 2609.00252, Spec-Driven Development for Agentic Software
  Engineering (conclusion).
---

:::grid{class="grid large-height center-align"}

:::col{class="s12"}

<blockquote class="left-align">
  <h5>On demandait un petit fix. On récupère un refactor, trois fichiers en plus, une logique « améliorée », et des tests qui passent parce qu'ils ne testent rien.</h5>
</blockquote>

:::

:::col{class="s4"}

##### L'intention vit dans le chat

<p class="xlarge-text">On ferme l'onglet, il ne reste que du code. Plus personne ne peut dire ce qui avait été demandé.</p>

<h6>Contrat perdu : la transférabilité</h6>

:::

:::col{class="s4"}

##### L'agent dérive

<p class="xlarge-text">Du code que personne n'a demandé, des décisions que personne n'a prises.</p>

<h6>Contrat perdu : la responsabilité</h6>

:::

:::col{class="s4"}

##### Rien n'est vérifiable

<p class="xlarge-text">Sans critères d'acceptation écrits avant, « ça a l'air bon » tient lieu de recette.</p>

<h6>Contrat perdu : la vérifiabilité</h6>

:::

:::col{class="s12"}

#### Ces trois contrats perdus, ce sont les trois *réponses* du SDD.

:::
:::

---
layout: default
title: Rien de tout ça n'est nouveau
speakernotes: |
  S9 - Slide anti-objection. Frise : user story + critères d'acceptation,
  Gherkin/BDD, TDD, contract-first OpenAPI, ADR, DDD.
  Message : on a toujours écrit des specs. Ce qui change, c'est qu'elles
  doivent être lisibles par une machine et versionnées.
  Glisser le glossaire en encart sur ce slide plutôt qu'en slide séparé :
  spec, plan, task, constitution, gate, subagent, drift. Gain d'une minute.
  L'encart n'en porte que six : subagent est sorti de l'écran pour que la carte
  reste lisible, il est défini au moment où il sert vraiment, en S21
  (Superpowers). Le dire à l'oral ici en une phrase, ou ne pas le dire du tout.
---

:::grid{class="grid large-height"}

:::col{class="s7 left-align"}

<ul class="xlarge-text">
  <li><strong>User story + critères d'acceptation</strong> : dire ce qu'on attend, avant de le construire.</li>
  <li><strong>Gherkin, BDD</strong> : la règle métier écrite dans une forme que la machine relit.</li>
  <li><strong>TDD</strong> : le critère d'abord, le code ensuite.</li>
  <li><strong>Contract-first, OpenAPI</strong> : le contrat fait foi, l'implémentation suit.</li>
  <li><strong>ADR</strong> : la décision et sa raison, versionnées avec le code.</li>
  <li><strong>DDD</strong> : les mots du métier, partagés par toute l'équipe.</li>
</ul>

:::

:::col{class="s5 left-align"}

<article class="round border">
  <h6>Le vocabulaire de la suite</h6>
  <ul class="xlarge-text">
    <li><strong>spec</strong> : le quoi et le pourquoi, avec les critères</li>
    <li><strong>plan</strong> : le comment, les choix techniques</li>
    <li><strong>task</strong> : une unité de travail vérifiable</li>
    <li><strong>constitution</strong> : les règles valables pour tout le projet</li>
    <li><strong>gate</strong> : le point où un humain valide avant la suite</li>
    <li><strong>drift</strong> : l'écart qui se creuse entre la spec et le code</li>
  </ul>
</article>

:::

:::col{class="s12 center-align"}

#### On a toujours écrit des specs. Ce qui change : elles doivent être *relisibles par une machine*, et versionnées.

:::
:::

---
layout: default
title: "DDD : pourquoi il compte plus qu'avant"
sources:
  - threedots.tech/post/ddd-and-ai-coding
  - arxiv.org/abs/2605.01160
speakernotes: |
  S10 - Langage ubiquitaire, bounded contexts, context map, agrégats,
  invariants.
  L'argument : un nommage précis et cohérent affûte les prompts et empêche les
  agents de confondre des concepts distincts. Les agents ont accès à tout le
  dépôt et cherchent les noms qu'on leur donne ; ils peuvent naïvement essayer
  d'unifier des entités similaires, donc il faut expliciter qu'elles sont
  séparées pour une raison.
  Trois usages directs : langage ubiquitaire = désambiguïsation de l'agent ;
  bounded contexts = périmètre de contexte et de parallélisation ; invariants =
  critères vérifiables.
  Mise en garde qui compte : la valeur du modèle de domaine réside dans la
  compréhension partagée qu'il construit dans l'équipe, pas dans l'artefact.
  Une spec générée par IA que personne ne lit ne résout rien.
  D'où la punchline du slide : ceux qui comprennent profondément leur domaine
  deviennent ceux qui peuvent juger si la sortie de l'IA est correcte, ce qui
  rend l'expertise du domaine plus précieuse que l'expertise d'un framework.
  Corollaire à dire : on conçoit et on discute la solution en équipe avant de
  générer, un meilleur contexte donne un meilleur résultat d'agent et une revue
  plus fluide.
  Sources : threedots.tech, Domain-Driven Design matters more when AI writes
  your code ; arXiv 2605.01160 (section DDD/DDT par tiers de délégation) ;
  talk gitnation From Prompt Spaghetti to Bounded Contexts.
---

:::grid{class="grid large-height center-align"}

:::col{class="s4"}

##### Langage ubiquitaire

<p class="xlarge-text">Un mot, un sens, partout. L'agent cherche les noms qu'on lui donne : deux concepts qui portent le même mot finissent fusionnés.</p>

:::

:::col{class="s4"}

##### Bounded contexts

<p class="xlarge-text">Des frontières écrites. Elles disent à l'agent où s'arrête son périmètre, et elles découpent le travail en tranches parallélisables.</p>

:::

:::col{class="s4"}

##### Invariants

<p class="xlarge-text">Les règles que le domaine ne tolère pas de casser. Ce sont déjà des critères d'acceptation.</p>

:::

:::col{class="s12"}

<blockquote class="left-align">
  <h5>La valeur du modèle de domaine, c'est la compréhension partagée qu'il construit dans l'équipe, pas le fichier. Une spec générée que personne ne lit ne résout rien.</h5>
</blockquote>

:::

:::col{class="s12"}

#### Connaître son domaine vaut désormais plus que connaître un *framework*.

:::
:::

---
layout: bare
class: transition
speakernotes: |
  Transition partie 2 - Le SDD et son intérêt. 10 minutes, S11 à S17. C'est le
  cœur du talk.
  Ce qu'on promet : la définition, les trois niveaux d'ambition, la boucle,
  ce qui fait une bonne spec, pourquoi ça marche mécaniquement - et le coût,
  honnêtement.
---

<h6 class="numero">Partie 2</h6>

<h2>Le spec-driven development</h2>

<h5>Ce que c'est, pourquoi ça marche, et ce que ça <em>coûte</em></h5>

---
layout: default
title: "Spec-Driven Development : la définition"
sources:
  - martinfowler.com/articles/exploring-gen-ai/sdd-3-tools.html
speakernotes: |
  S11 - Le SDD traite les spécifications comme des contrats exécutables
  desquels les agents dérivent le code, en empêchant la dérive architecturale
  par une application automatisée plutôt que par de la documentation passive.
  La formule à afficher : the spec is the prompt.
  Source : Birgitta Böckeler, Understanding Spec-Driven Development
  (martinfowler.com) ; podcast Thoughtworks What is spec-driven development.
---

:::grid{class="grid large-height center-align"}

:::col{class="s12"}

### La spec devient un *contrat exécutable*, dont les agents dérivent le code.

:::

:::col{class="s6"}

##### La documentation

<p class="xlarge-text">Elle décrit ce qu'on a fait. Elle est passive : rien ne se
casse quand le code s'en éloigne, et personne ne s'en aperçoit.</p>

:::

:::col{class="s6"}

##### La spec

<p class="xlarge-text">Elle dit ce qu'on attend. Elle est appliquée : la dérive
se détecte par l'outillage, pas par la bonne volonté de chacun.</p>

:::

:::col{class="s12"}

#### *The spec is the prompt.*

<p class="xlarge-text">Ce qu'on écrit une fois sert trois fois : à demander,
à revoir, à vérifier.</p>

:::
:::

---
layout: content
title: Les 3 niveaux d'ambition
sources:
  - martinfowler.com/articles/exploring-gen-ai/sdd-3-tools.html
speakernotes: |
  S12 - Le cadre analytique le plus cité du domaine (Böckeler), à ne pas
  sauter.
  Spec-first : la spec guide l'implémentation puis est abandonnée, le code
  reste l'artefact maintenu. Point d'entrée pragmatique, là où sont la plupart
  des équipes qui démarrent.
  Spec-anchored : la spec persiste comme contrat vivant, versionnée et mise à
  jour. La cible réaliste 2026.
  Spec-as-source : les humains ne maintiennent que la spec, tout le code est
  généré, marqué comme généré, et jamais édité à la main. Vision long terme,
  encore largement expérimentale.
  L'analogie à raconter sur le niveau 3, c'est elle qui fait comprendre : c'est
  ce que les compilateurs ont fait à l'assembleur, devenu un artefact que plus
  personne n'édite à la main.
  Dire explicitement à quel niveau on propose de jouer : ça évite 80 % des
  malentendus dans la salle. C'est aussi la moitié de la réponse à la question
  "c'est le cycle en V déguisé ?".
  Source : Böckeler, martinfowler.com/articles/exploring-gen-ai/sdd-3-tools.html
---

<table class="border xlarge-text">
  <thead>
    <tr>
      <th>Niveau</th>
      <th>Ce qui arrive à la spec</th>
      <th>Ce qu'on maintient</th>
      <th>Où ça en est</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><strong>1. Spec-first</strong></td>
      <td>Écrite avant le code, elle pilote le premier passage de l'agent. Puis on la laisse tomber.</td>
      <td>Le code</td>
      <td>Le point d'entrée. Là où sont la plupart des équipes qui démarrent.</td>
    </tr>
    <tr>
      <td><strong>2. Spec-anchored</strong></td>
      <td>Elle survit à la livraison : c'est le document vivant de la fonctionnalité, itération après itération.</td>
      <td>Le code <em>et</em> la spec</td>
      <td>La cible réaliste pour 2026. C'est là qu'on joue aujourd'hui.</td>
    </tr>
    <tr>
      <td><strong>3. Spec-as-source</strong></td>
      <td>Tout le code est généré, marqué comme généré, et jamais édité à la main.</td>
      <td>La spec, et rien d'autre</td>
      <td>Vision long terme, encore largement expérimentale.</td>
    </tr>
  </tbody>
</table>

<p class="xlarge-text center-align">Le niveau 3, c'est ce que les compilateurs
ont fait à l'assembleur : un artefact que <em>plus personne n'édite à la
main</em>.</p>

---
layout: default
title: La boucle canonique
speakernotes: |
  S13 - Le diagramme à produire soi-même, et à réutiliser 3 fois dans le deck
  comme repère de progression.
  intention, clarification, spec (quoi/pourquoi + critères), plan (comment),
  tasks atomiques, implémentation, vérification, réconciliation spec/code.
  Les gates humains en rouge.
---

:::split{height=xlarge}

<pre class="mermaid">
flowchart LR
    classDef gate fill:#d32f2f,stroke:#d32f2f,color:#ffffff,font-weight:bold
    classDef step fill:#ffffff,stroke:#0f15fd,color:#0f15fd,stroke-width:2px

    I([Intention]) --> C{{"Clarification<br/>humain"}}
    C --> S["Spec"]
    S --> V{{"Validation<br/>humain"}}
    V --> P["Plan"]
    P --> T["Tasks"]
    T --> IMP["Implémentation"]
    IMP --> VER["Vérification"]
    VER --> R{{"Réconciliation<br/>humain"}}
    R -. "la spec repasse devant le code" .-> S

    class C,V,R gate
    class S,P,T,IMP,VER step
</pre>

:::

---
layout: default
title: Anatomie d'une bonne spec
sources:
  - github.com/github/spec-kit
speakernotes: |
  S14 - Le slide dont on se souviendra, à condition de le nourrir avec une
  spec réelle de mon propre code, en deux versions : floue vs complète, avec le
  résultat produit par l'agent dans les deux cas. C'est du travail à produire.
  Une bonne spec fixe les résultats attendus, les limites de périmètre, les
  contraintes, les décisions antérieures, la découpe en tâches et les critères
  de vérification.
  Punchline : tout ce que tu n'écris pas, l'agent l'invente.
  Source : github/spec-kit (lire les templates, pas le README).
---

:::grid{class="grid large-height"}

:::col{class="s7 left-align"}

<ul class="xlarge-text">
  <li><strong>Le résultat attendu</strong> : ce qui doit être vrai quand c'est fini.</li>
  <li><strong>Les limites du périmètre</strong> : ce qu'on ne touche pas, et qu'on ne veut pas voir bouger.</li>
  <li><strong>Les contraintes</strong> : performance, sécurité, compatibilité, budget.</li>
  <li><strong>Les décisions déjà prises</strong> : et la raison pour laquelle on ne les rouvre pas.</li>
  <li><strong>La découpe</strong> : des tâches assez petites pour être vérifiées une par une.</li>
  <li><strong>Les critères de vérification</strong> : écrits <em>avant</em>, jamais après.</li>
</ul>

:::

:::col{class="s5 left-align"}

<article class="round border">
  <h6>Le test qui ne trompe pas</h6>
  <p class="xlarge-text">Donne ta spec à un collègue qui ne connaît pas le
  sujet. S'il doit te poser une question avant de commencer, note-la : c'est
  exactement le trou que l'agent, lui, ne signalera pas. Il choisira tout
  seul, et il choisira vite.</p>
</article>

:::

:::col{class="s12 center-align"}

#### Tout ce que tu n'écris pas, l'agent l'*invente*.

:::
:::

---
layout: default
title: "Pourquoi ça marche, mécaniquement"
speakernotes: |
  S15 - Cinq mécanismes, pas cinq bénéfices marketing.
  1. Externalise l'état que le LLM n'a pas.
  2. Déplace la revue en amont : 200 lignes de markdown au lieu de 2000 lignes
  de diff.
  3. Rend l'output vérifiable.
  4. Rend le travail reprenable entre sessions et transférable entre agents.
  5. Permet le parallélisme.
  Appui : le SDD attrape les violations architecturales et la dérive de contrat
  d'API que les tests unitaires ne peuvent structurellement pas détecter, et il
  scale sur des agents parallèles en séparant le rôle qui implémente de celui
  qui vérifie.
---

:::grid{class="grid large-height center-align"}

:::col{class="s6"}

##### En amont : l'*intention*

<ul class="xlarge-text left-align">
  <li>La spec porte l'état que le modèle n'a pas. Le travail redevient reprenable d'une session à l'autre, et transférable d'un agent à l'autre.</li>
  <li>Des contrats écrits et des périmètres séparés : plusieurs agents avancent en même temps sans se marcher dessus.</li>
</ul>

:::

:::col{class="s6"}

##### En aval : la *vérification*

<ul class="xlarge-text left-align">
  <li>Les critères sont écrits avant, donc la sortie se juge au lieu de se deviner. Et celui qui vérifie n'est plus celui qui implémente.</li>
  <li>Violation d'architecture, contrat d'API qui glisse : un test unitaire ne peut structurellement pas les voir. Une spec, si.</li>
</ul>

:::

:::col{class="s12"}

<blockquote class="left-align">
  <h5>La revue remonte d'un cran : 200 lignes de markdown qu'on a écrites, au lieu de 2 000 lignes de diff qu'on n'a pas écrites.</h5>
</blockquote>

:::

:::col{class="s12"}

#### Aucun de ces gains ne vient du modèle. Ils viennent du fait d'avoir écrit.

:::
:::

---
layout: content
title: "SDD, TDD, BDD, cycle en V"
sources:
  - martinfowler.com/articles/exploring-gen-ai/sdd-3-tools.html
  - marmelab.com
speakernotes: |
  S16 - Tableau compact, puis traiter frontalement la critique.
  Fowler soutient qu'une spécification utile dépend de l'apprentissage acquis
  pendant le développement, et que la clé de l'usage plein de l'IA est
  d'accélérer les boucles de feedback.
  La nuance à porter : Beck s'oppose à l'écriture de la spécification complète
  avant l'implémentation. Le niveau 2 traite la spec comme un document vivant
  tout au long de l'implémentation. La critique vise surtout le niveau 3 et
  tout workflow qui gèle les hypothèses.
  Avec S12, ce slide forme l'autre moitié de la réponse à "c'est le cycle en V
  déguisé ?". La question tombera : ne pas attendre les questions pour y
  répondre, la traiter ici, frontalement.
  Sources : Kent Beck et Martin Fowler ; François Zaninotto, Waterfall Strikes
  Back (la critique à connaître pour tenir les questions).
---

<table class="border xlarge-text">
  <thead>
    <tr>
      <th></th>
      <th>Ce qu'on écrit avant</th>
      <th>Pour qui</th>
      <th>Ce qui se passe quand on apprend en chemin</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><strong>TDD</strong></td>
      <td>Un test qui échoue</td>
      <td>La machine</td>
      <td>On change le test. Toutes les deux minutes.</td>
    </tr>
    <tr>
      <td><strong>BDD</strong></td>
      <td>Un scénario en mots du métier</td>
      <td>L'équipe, puis la machine</td>
      <td>On réécrit le scénario avec le métier.</td>
    </tr>
    <tr>
      <td><strong>Cycle en V</strong></td>
      <td>La spécification complète</td>
      <td>Le contrat</td>
      <td>On ouvre un avenant. L'apprentissage coûte cher, donc on l'évite.</td>
    </tr>
    <tr>
      <td><strong>SDD</strong></td>
      <td>Le quoi, le pourquoi, les critères</td>
      <td>L'agent, puis le reviewer</td>
      <td>On met la spec à jour, et on relance. <em>C'est la manœuvre normale</em>, pas l'exception.</td>
    </tr>
  </tbody>
</table>

<blockquote class="left-align">
  <h5>La critique, telle qu'elle se formule : une spec utile dépend de ce qu'on apprend en développant, et l'intérêt de l'IA est d'accélérer les boucles de feedback, pas de les geler.</h5>
  <h5>Elle est juste. Elle vise le niveau 3, et tout workflow qui fige ses hypothèses. Au niveau 2, la spec <em>change pendant</em> l'implémentation.</h5>
</blockquote>

---
layout: default
title: "Le coût, honnêtement"
sources:
  - martinfowler.com/articles/exploring-gen-ai/sdd-3-tools.html
  - dora.dev/ai/roi/report
speakernotes: |
  S17 - Slide de crédibilité. Trois coûts : cérémonie mal dimensionnée,
  tokens, dérive spec/code.
  L'anecdote parfaite : Böckeler a lancé Kiro sur une petite correction de bug
  et a obtenu quatre user stories et seize critères d'acceptation. Le SDD passe
  mal à l'échelle vers le bas : pour un null check d'une ligne, on n'a pas à
  toucher à la constitution. La formule qui circule : le marteau-pilon pour un
  petit clou.
  Le chiffre pour la question "ça coûte combien ?", qui tombera : un run de
  workflow BMAD consomme environ 31 700 tokens en moyenne, et sur de gros
  projets cela représente 800 à 2 000 dollars par mois et par développeur en
  coûts d'API.
  Le chiffre de dette, à donner comme ce que le SDD cherche à éviter, pas comme
  ce qu'il cause : DORA 2026 mesure 30 à 41 % de dette technique en plus avec
  l'adoption de l'IA, et des PR générées par IA qui portent 1,7 fois plus de
  problèmes que celles écrites par des humains.
  Sur la dérive : aucun de ces outils ne réconcilie automatiquement. Il faut le
  déclencher et relire la sortie, et c'est cette charge que la plupart des
  équipes esquivent jusqu'à ce que leurs specs aient six mois de retard.
  Ce slide porte la réponse à deux des cinq questions attendues : le coût en
  tokens, et qui maintient les specs quand elles divergent. Assumer que c'est
  le point faible non résolu du domaine.
  Sources : Böckeler, martinfowler.com ; DORA 2026, The ROI of AI-assisted
  Software Development (Google) ; mesures de consommation de tokens sur des
  runs BMAD.
---

:::grid{class="grid large-height center-align"}

:::col{class="s4"}

<h3><em>4 stories, 16 critères</em></h3>

<p class="xlarge-text">pour <strong>un seul petit bug</strong></p>

<p class="xlarge-text">La cérémonie ne se dimensionne pas toute seule. Pour un
null check d'une ligne, on ne touche pas à la constitution.</p>

<h6>Birgitta Böckeler, en lançant Kiro</h6>

:::

:::col{class="s4"}

<h3><em>31 700</em></h3>

<p class="xlarge-text">tokens en moyenne<br/>par run de workflow</p>

<p class="xlarge-text">Sur de gros projets, ça fait <strong>800 à 2 000 $ par
mois et par développeur</strong> de coûts d'API. La cérémonie se paie en
tokens.</p>

<h6>Runs BMAD</h6>

:::

:::col{class="s4"}

<h3><em>+30 à 41 %</em></h3>

<p class="xlarge-text">de dette technique<br/>avec l'adoption de l'IA</p>

<p class="xlarge-text">Et des PR générées par IA qui portent <strong>1,7 fois
plus de problèmes</strong>. C'est ce que le SDD cherche à éviter, pas ce qu'il
garantit d'éviter.</p>

<h6>DORA 2026</h6>

:::

:::col{class="s12"}

<blockquote class="left-align">
  <h5>Le marteau-pilon pour un petit clou : le SDD passe mal à l'échelle vers le bas.</h5>
  <h5>Et la dérive spec/code, personne ne l'a résolue. Aucun outil ne réconcilie tout seul : il faut le déclencher, et relire la sortie.</h5>
</blockquote>

:::
:::

---
layout: bare
class: transition
speakernotes: |
  Transition partie 3 - Le panorama des frameworks. 10 minutes, S18 à S24.
  Le piège n°1 du talk est ici : le catalogue d'outils. Le dire à voix haute en
  affichant cette slide - "ce n'est pas un comparatif d'outils, c'est une
  démonstration que tous font la même chose".
  Si la salle est plus junior que prévu, c'est ici qu'on coupe : fusionner
  S19 à S23 en un seul tableau et récupérer 3 minutes pour les parties 1 et 2.
---

<h6 class="numero">Partie 3</h6>

<h2>Le panorama des frameworks</h2>

<h5>Mode plan, Spec Kit, OpenSpec, Superpowers, BMAD, et ce qu'ils ont <em>en commun</em></h5>

---
layout: default
title: "La carte par couches, et le niveau 0 : le mode plan"
speakernotes: |
  S18 - Deux choses sur un slide.
  D'abord pourquoi c'est confus : les outils opèrent à des couches différentes
  (définition des artefacts d'exigence, conversion en graphe de tâches,
  exécution du code, intégration IDE). Une cartographie communautaire
  recensait plus de 30 outils début 2026.
  Ensuite la baseline : le mode plan (Claude Code, Cursor). Plan éphémère,
  aucun artefact persisté, aucun gate, zéro installation. Excellent pour une
  tâche de 30 minutes, insuffisant dès qu'il y a plusieurs sessions, plusieurs
  agents ou une revue par un tiers.
  C'est le témoin de comparaison sur les 4 slides suivants - le dire
  explicitement, sinon le panorama devient un catalogue.
---

:::grid{class="grid large-height"}

:::col{class="s5 left-align"}

<h6>Pourquoi c'est illisible : ils ne jouent pas au même étage</h6>

<ul class="xlarge-text">
  <li><strong>Définir</strong> les artefacts d'exigence</li>
  <li><strong>Convertir</strong> en graphe de tâches</li>
  <li><strong>Exécuter</strong> le code</li>
  <li><strong>S'intégrer</strong> à l'IDE</li>
</ul>

<p class="xlarge-text">Plus de <strong>30 outils</strong> recensés début 2026.
Deux d'entre eux ne font souvent pas le même métier.</p>

:::

:::col{class="s7 left-align"}

<article class="round border">
  <h6>Niveau 0 : le mode plan (Claude Code, Cursor)</h6>
  <p class="xlarge-text"><strong>Ce qu'il fait</strong> : l'agent écrit son
  plan, tu le lis, tu corriges, il exécute. Zéro installation.</p>
  <p class="xlarge-text"><strong>Où il s'arrête</strong> : le plan est
  éphémère. Rien n'est persisté, rien n'est versionné, aucun gate n'est
  imposé. Parfait pour une tâche de 30 minutes. Insuffisant dès qu'il y a
  plusieurs sessions, plusieurs agents, ou une revue par un tiers.</p>
</article>

:::

:::col{class="s12 center-align"}

#### Le mode plan est notre *témoin*. Les quatre outils qui suivent se lisent par rapport à lui.

:::
:::

---
layout: default
title: GitHub Spec Kit
sources:
  - github.com/github/spec-kit
  - github.github.com/spec-kit
  - ranthebuilder.cloud/blog/i-tested-three-spec-driven-ai-tools-here-s-my-honest-take
  - reenbit.com/bmad-vs-spec-kit-vs-openspec-choosing-your-spec-driven-ai-framework
speakernotes: |
  S19 - CLI Python, et des dizaines d'intégrations d'agents.
  ATTENTION - correction du brief : ne PAS annoncer de chiffre d'étoiles. Les
  sources vont de 80 000 à 120 000 selon la date, le "129 000" qui circulait
  n'est corroboré nulle part. Même chose pour le nombre d'intégrations
  d'agents : 24+ ou 38 selon la source, donc dire "des dizaines" et passer.
  C'est de toute façon la plus forte distribution de la catégorie, et c'est ça
  le point, pas le chiffre.
  Trois commandes portent le workflow : /speckit.specify capture le contexte
  métier et les critères de succès, /speckit.plan traduit en décisions
  d'architecture, /speckit.tasks décompose en unités testables.
  Différenciateurs : une constitution définie une fois pour le projet dont
  chaque spec hérite, et des templates qui marquent les inconnues en
  NEEDS CLARIFICATION plutôt que de deviner.
  Faiblesses : pas d'étape de revue de code intégrée, et changer de direction
  implique de relancer les commandes concernées, chacune régénérant tout son
  document.
  Profil : greenfield, gates forts, verbeux.
  Gabarit commun aux quatre slides d'outils : à gauche ce qu'il ajoute au mode
  plan, à droite ce qu'il demande en retour, en bas une phrase de profil. La
  salle compare d'un coup d'oeil au lieu de relire.
  Sources : github/spec-kit (lire les templates, pas le README) ;
  ranthebuilder.cloud, I Tested Three Spec-Driven AI Tools ;
  reenbit.com/bmad-vs-spec-kit-vs-openspec. Vérifié le 11/09/2026 : les
  commandes, la constitution et le marqueur NEEDS CLARIFICATION sont solides,
  les chiffres de popularité ne le sont pas.
---

:::grid{class="grid large-height"}

:::col{class="s7 left-align"}

<h6>Ce qu'il ajoute au mode plan</h6>

<ul class="xlarge-text">
  <li><strong>Trois commandes, trois fichiers qui restent</strong> :
  <code>specify</code> pour le métier et les critères de succès,
  <code>plan</code> pour les décisions d'architecture, <code>tasks</code> pour
  la découpe en unités testables.</li>
  <li><strong>Une constitution</strong> écrite une fois pour le projet, dont
  chaque spec hérite ensuite.</li>
  <li><strong>Les inconnues sont marquées</strong>, pas devinées : les
  templates posent un <code>NEEDS CLARIFICATION</code> là où l'agent aurait
  choisi tout seul.</li>
</ul>

:::

:::col{class="s5 left-align"}

<article class="round border">
  <h6>Ce qu'il demande en retour</h6>
  <ul class="xlarge-text">
    <li>Un CLI Python à installer, et beaucoup de texte à lire.</li>
    <li>Aucune étape de revue de code dans le workflow.</li>
    <li>Changer de direction, c'est relancer la commande : elle régénère tout
    son document.</li>
  </ul>
</article>

:::

:::col{class="s12 center-align"}

#### Projet neuf, gates visibles, et on accepte la *verbosité*. C'est l'outil le plus distribué de la catégorie.

:::
:::

---
layout: default
title: "OpenSpec : l'anti-cérémonie"
sources:
  - github.com/Fission-AI/OpenSpec
  - reenbit.com/bmad-vs-spec-kit-vs-openspec-choosing-your-spec-driven-ai-framework
  - ranthebuilder.cloud/blog/i-tested-three-spec-driven-ai-tools-here-s-my-honest-take
speakernotes: |
  S20 - L'empreinte la plus légère : on écrit des delta specs, uniquement ce
  qui change.
  Installation en 5 minutes contre 30, pas de Python, sortie d'environ 250
  lignes contre 800, conçu pour les bases de code existantes, 3 commandes IA
  contre 8, mise à jour de n'importe quel artefact à tout moment sans phase
  gates rigides.
  Workflow OPSX : explore, propose, apply, sync, archive.
  Contrepartie, à dire honnêtement : pas de gates de revue entre phases, et un
  support communautaire limité.
  C'est aussi la réponse à la question "et le legacy ?", qui tombera.
  Le mouvement à souligner, sinon la slide devient une fiche produit : c'est
  exactement le même travail que Spec Kit, avec une dose de cérémonie plus
  faible. Pas un outil de plus, la même chose moins chère.
  Sources : github.com/Fission-AI/OpenSpec, version 1.2.0 (docs/opsx.md et
  docs/commands.md) ; reenbit.com/bmad-vs-spec-kit-vs-openspec et
  ranthebuilder.cloud pour les chiffres de comparaison. Vérifié le 11/09/2026.
---

:::grid{class="grid large-height"}

:::col{class="s7 left-align"}

<h6>Ce qu'il ajoute au mode plan</h6>

<ul class="xlarge-text">
  <li><strong>Des delta specs</strong> : on décrit ce qui change, pas le
  système entier. Chaque document reste court et relisable.</li>
  <li><strong>Cinq étapes, trois commandes</strong> : explore, propose, apply,
  sync, archive.</li>
  <li><strong>N'importe quel artefact se corrige à tout moment</strong>, sans
  phase gate à repasser.</li>
</ul>

<p class="xlarge-text">Installation en <strong>5 minutes</strong> au lieu de
30, pas de Python, environ <strong>250 lignes</strong> produites au lieu de
800.</p>

:::

:::col{class="s5 left-align"}

<article class="round border">
  <h6>Ce qu'il demande en retour</h6>
  <ul class="xlarge-text">
    <li>Aucun gate de revue entre les phases : la discipline reste chez
    toi.</li>
    <li>Un support communautaire limité.</li>
  </ul>
</article>

:::

:::col{class="s12 center-align"}

#### Le même travail que Spec Kit, avec *moins de cérémonie*. Et c'est lui qui est pensé pour du code qui existe déjà.

:::
:::

---
layout: default
title: "Superpowers : la méthode comme artefact"
sources:
  - claude.com/plugins/superpowers
  - github.com/obra/superpowers-marketplace
  - simonwillison.net/2025/Oct/10/superpowers
speakernotes: |
  S21 - Le cas le plus intéressant pédagogiquement. Ce n'est pas un dépôt de
  specs, c'est une méthodologie livrée comme skills.
  Plugin qui impose un workflow structuré avant qu'une seule ligne de code ne
  soit écrite : brainstormer d'abord, isoler sa branche, écrire un plan
  détaillé, exécuter. Chaque étape conditionne la suivante.
  Le plan découpe le travail en tâches de 2 à 5 minutes avec chemins de
  fichiers exacts et commandes exactes. Des subagents implémentent chaque
  tâche, avec une revue intégrée.
  Pratiques imposées : cycles TDD red-green-refactor où les tests doivent
  échouer avant l'implémentation, méthodologie de debug en quatre phases
  exigeant l'investigation de la cause racine avant tout correctif,
  brainstorming socratique qui force l'humain à formuler ce qu'il veut
  vraiment, pas ce qu'il croit vouloir.
  Les deux formulations vérifiées qui portent le slide, à dire telles quelles :
  le plugin traite les agents comme un manager expérimenté traite des juniors ;
  et ses skills sont écrites pour contrer la tendance de l'agent à rationaliser
  le fait de sauter les process structurés.
  L'angle à souligner devant des devs : c'est une réponse directe au problème
  de discipline. Le SDD est un problème de discipline avant d'être un problème
  d'outil, et ce slide est le seul endroit du talk où on peut le dire aussi
  directement. Quelqu'un écrit une bonne skill specify-plan-implement,
  l'utilise une semaine, puis retourne discrètement au prompting non structuré
  dès qu'une deadline approche.
  Anecdote : Jesse Vincent, créateur de RT, contributeur de Perl 5, auteur du
  client mail K-9, a quasiment cessé de coder lui-même.
  Chiffres vérifiés le 11/09/2026, utilisables à l'oral : sorti en octobre
  2025, plus de 40 000 étoiles GitHub, et plus d'installations sur le
  marketplace Claude Code que Playwright.
  Sources : claude.com/plugins/superpowers ; obra/superpowers-marketplace ;
  blog de Jesse Vincent sur la genèse du plugin.
---

:::grid{class="grid large-height"}

:::col{class="s7 left-align"}

<h6>Ce qu'il ajoute au mode plan</h6>

<ul class="xlarge-text">
  <li><strong>Un ordre imposé</strong>, avant la moindre ligne de code :
  brainstormer, isoler sa branche, écrire le plan, exécuter. Chaque étape
  conditionne la suivante.</li>
  <li><strong>Un plan en tâches de 2 à 5 minutes</strong>, avec les chemins de
  fichiers et les commandes exactes. Des subagents les implémentent, revue
  comprise.</li>
  <li><strong>Des pratiques non négociables</strong> : le test doit échouer
  avant l'implémentation, le debug passe par la cause racine avant tout
  correctif.</li>
</ul>

:::

:::col{class="s5 left-align"}

<article class="round border">
  <h6>Ce qu'il demande en retour</h6>
  <ul class="xlarge-text">
    <li>Le plugin traite les agents comme un manager expérimenté traite des
    juniors. À toi de tenir le rôle du manager.</li>
    <li>Ses skills sont écrites pour contrer la tendance de l'agent à
    rationaliser le fait de sauter les étapes.</li>
  </ul>
</article>

:::

:::col{class="s12 center-align"}

#### Ce n'est pas un dépôt de specs, c'est une méthode livrée comme des skills. Le SDD est d'abord un problème de *discipline*.

:::
:::

---
layout: default
title: "BMAD-METHOD : le poids lourd"
sources:
  - github.com/bmad-code-org/BMAD-METHOD
speakernotes: |
  S22 - Simule une équipe agile complète via des personas nommés : Analyst,
  PM, Architect, UX Designer, Scrum Master, Developer, QA, Tech Writer. Chaque
  agent reçoit une fenêtre de contexte étroitement délimitée et produit un
  artefact versionné.
  ATTENTION - correction du brief : 12 agents, pas 19. Le chiffre de 19 qui
  figurait ici est faux, ne pas l'annoncer.
  L'agent Scrum Master crée des fichiers de story détaillés portant le contexte
  architectural, les guidelines d'implémentation et les critères de test pour
  l'agent Dev.
  Trois pistes de planification : Quick Flow (tech spec seule, pour un bugfix
  ou une petite feature), BMad Method (PRD, architecture et UX, pour un produit
  ou une plateforme), Enterprise (planification étendue avec sécurité, DevOps
  et tests, pour la conformité).
  Le mouvement du slide, à ne pas rater : c'est encore le même squelette que
  Spec Kit et OpenSpec, avec la dose de cérémonie poussée au maximum. Le coût
  chiffré arrive au slide suivant, ne pas le donner ici.
  Source : github.com/bmad-code-org/BMAD-METHOD, version 6.0.3 (février 2026).
  Vérifié le 11/09/2026.
---

:::grid{class="grid large-height"}

:::col{class="s7 left-align"}

<h6>Ce qu'il ajoute au mode plan</h6>

<ul class="xlarge-text">
  <li><strong>12 agents aux rôles nommés</strong> : Analyst, PM, Architect, UX
  Designer, Scrum Master, Developer, QA, Tech Writer. Chacun reçoit un contexte
  étroit et rend un artefact versionné.</li>
  <li><strong>Le Scrum Master écrit les stories que le Dev implémente</strong>
  : contexte d'architecture, guidelines, critères de test, dans le fichier.</li>
  <li><strong>Trois pistes selon la taille</strong> : Quick Flow pour un
  bugfix, BMad Method avec PRD, architecture et UX pour un produit, Enterprise
  quand la conformité s'en mêle.</li>
</ul>

:::

:::col{class="s5 left-align"}

<article class="round border">
  <h6>Ce qu'il demande en retour</h6>
  <ul class="xlarge-text">
    <li>Le rituel complet, même quand la tâche ne le justifie pas.</li>
    <li>Du temps, et des tokens. Le chiffre est sur le slide suivant.</li>
  </ul>
</article>

:::

:::col{class="s12 center-align"}

#### Toujours le même squelette que les deux précédents, avec la cérémonie poussée au *maximum*.

:::
:::

---
layout: content
title: Le tableau de décision
sources:
  - reenbit.com/bmad-vs-spec-kit-vs-openspec-choosing-your-spec-driven-ai-framework
  - ranthebuilder.cloud/blog/i-tested-three-spec-driven-ai-tools-here-s-my-honest-take
speakernotes: |
  S23 - Le chiffre qui fait rire la salle : sur un même build de dashboard CRM,
  la même tâche a pris 12 minutes avec OpenSpec, 90 minutes avec Spec Kit et
  5 h 30 avec BMAD.
  Puis la grille : pour la plupart des équipes travaillant sur du code
  existant, OpenSpec offre le meilleur équilibre vitesse/flexibilité ; pour les
  nouveaux projets avec des rôles clairs, Spec Kit apporte structure et
  documentation ; pour la complexité d'échelle entreprise, BMAD gère
  l'orchestration multi-agents.
  L'honnêteté qui fait la différence : BMAD est excellent, mais aussi coûteux
  et disproportionné pour la plupart du travail hebdomadaire d'ingénierie.
  Mentionner en une ligne, sans slide : Kiro, Tessl, cc-sdd, Antigravity, et la
  convergence AGENTS.md / Agent Skills comme lingua franca émergente.
  Le coût de BMAD, vérifié, bien plus parlant que "c'est cher" : environ
  31 700 tokens par run de workflow, et 800 à 2 000 dollars par mois et par
  développeur en coûts d'API sur de gros projets. Les 31 700 tokens sont déjà
  passés à l'écran en S17, donc seul le chiffre en dollars est repris ici. Si
  la question du coût revient, c'est le moment de rappeler les deux.
  La ligne "mode plan" en tête de tableau n'est pas décorative : elle referme
  la comparaison ouverte en S18 et empêche la partie de se lire comme un
  catalogue.
  Sources des durées et de la grille : reenbit.com/bmad-vs-spec-kit-vs-openspec
  ; ranthebuilder.cloud, I Tested Three Spec-Driven AI Tools ; mesures de
  consommation de tokens sur des runs BMAD. Vérifié le 11/09/2026.
---

<table class="border xlarge-text">
  <thead>
    <tr>
      <th></th>
      <th>Même tâche, même build de dashboard CRM</th>
      <th>Quand c'est lui qu'il faut prendre</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><strong>Mode plan</strong></td>
      <td>Rien à installer, rien à ranger</td>
      <td>Une tâche, une session, un seul lecteur : toi.</td>
    </tr>
    <tr>
      <td><strong>OpenSpec</strong></td>
      <td>12 minutes</td>
      <td>Du code qui existe déjà. Le meilleur équilibre vitesse / flexibilité pour la plupart des équipes.</td>
    </tr>
    <tr>
      <td><strong>Spec Kit</strong></td>
      <td>90 minutes</td>
      <td>Un projet neuf avec des rôles clairs, quand on veut de la structure et de la documentation.</td>
    </tr>
    <tr>
      <td><strong>BMAD</strong></td>
      <td>5 h 30</td>
      <td>La complexité d'échelle entreprise, quand l'orchestration multi-agents se justifie vraiment.</td>
    </tr>
  </tbody>
</table>

<blockquote class="left-align">
  <h5>BMAD est excellent. Il est aussi coûteux, et disproportionné pour la plus grande partie du travail d'une semaine d'ingénierie : <strong>800 à 2 000 $ par mois et par développeur</strong> en API sur de gros projets.</h5>
  <h5>Le bon réflexe n'est pas de choisir le plus complet. C'est de choisir la <em>plus petite dose de cérémonie</em> qui tient le problème.</h5>
</blockquote>

---
layout: default
title: Ce sur quoi ils sont tous d'accord
speakernotes: |
  S24 - Slide de sortie de partie. Quatre primitives universelles :
  règles/constitution, spec, plan, tasks, plus des gates humains.
  Malgré des approches différentes, tous ces frameworks s'accordent sur un
  point : l'humain reste dans la boucle, mais pas pour tout.
  Conclusion libératrice, c'est le message à laisser : tu peux commencer demain
  avec trois fichiers markdown et un gate. Le framework est une optimisation,
  pas un prérequis.
---

<div class="large-space"></div>

### règles &rarr; spec &rarr; plan &rarr; tasks

##### Plus des gates humains. Les quatre outils écrivent les mêmes artefacts : ils ne diffèrent que par la dose de cérémonie.

<div class="large-space"></div>

#### L'humain reste dans la boucle. Mais pas sur tout.

<div class="small-space"></div>

### Tu peux commencer demain avec *trois fichiers markdown* et un gate.

##### Le framework est une optimisation. Pas un prérequis.

---
layout: bare
class: transition
speakernotes: |
  Transition partie 4 - Découpler et accélérer. 7 minutes, S25 à S29.
  C'est la partie qui répond à la cinquième question du brief : en quoi le SDD
  accélère réellement.
  Ce qu'on promet : cinq découplages, celui qui compte vraiment, les chiffres
  avec leurs limites, là où ça n'accélère pas, et quoi mesurer.
---

<h6 class="numero">Partie 4</h6>

<h2>Découpler, et accélérer</h2>

<h5>Ce que le SDD sépare, et ce que ça fait gagner, <em>vraiment</em></h5>

---
layout: content
title: Les 5 découplages
speakernotes: |
  S25 - 1. Le quoi / le comment : la spec devient l'interface entre jugement
  humain et exécution machine.
  2. La conception / l'exécution dans le temps : tu spécifies maintenant, les
  agents exécutent pendant que tu fais autre chose. Ton temps de clavier n'est
  plus le facteur limitant.
  3. La revue / le code : revoir l'intention (200 lignes lisibles) au lieu du
  diff (2000 lignes qu'on n'a pas écrites).
  4. Les développeurs entre eux : specs comme contrats + bounded contexts =
  agents parallèles sur des tranches indépendantes, sans collision (worktrees,
  subagents).
  5. L'humain / la session : spec et plan comme mémoire durable. Reprise après
  crash de session, transfert entre agents, onboarding.
  Annoncer que le n°3 est celui sur lequel on s'arrête.
  Ne pas commenter les cinq lignes une par une : lire la colonne de gauche,
  s'arrêter sur la ligne 3, et enchaîner sur S26.
---

<table class="border xlarge-text">
  <thead>
    <tr>
      <th>Ce que le SDD sépare</th>
      <th>Ce que ça débloque</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><strong>1. Le quoi et le comment</strong></td>
      <td>La spec devient l'interface entre le jugement humain et l'exécution machine.</td>
    </tr>
    <tr>
      <td><strong>2. La conception et l'exécution</strong></td>
      <td>Tu spécifies maintenant, les agents exécutent pendant que tu fais autre chose. Ton temps de clavier n'est plus le facteur limitant.</td>
    </tr>
    <tr>
      <td><strong>3. La revue et le code</strong></td>
      <td>On revoit <em>l'intention</em>, 200 lignes lisibles, au lieu d'un diff de 2 000 lignes qu'on n'a pas écrites.</td>
    </tr>
    <tr>
      <td><strong>4. Les développeurs entre eux</strong></td>
      <td>Des specs comme contrats, des bounded contexts, et des agents en parallèle sur des tranches qui ne se percutent pas.</td>
    </tr>
    <tr>
      <td><strong>5. L'humain et la session</strong></td>
      <td>Spec et plan tiennent lieu de mémoire durable : reprise après un crash, passage d'un agent à l'autre, onboarding.</td>
    </tr>
  </tbody>
</table>

<blockquote class="left-align">
  <h5>Quatre de ces découplages font gagner du confort. Le troisième est le seul qui déplace vraiment le goulot, et c'est celui qu'on regarde maintenant.</h5>
</blockquote>

---
layout: default
title: "Le découplage qui compte : la revue"
sources:
  - dora.dev/ai/roi/report
speakernotes: |
  S26 - Pourquoi le découplage n°3 est LE gain.
  Le rapport DORA 2026 montre que l'IA n'élimine pas les goulots, elle déplace
  souvent le problème à l'étape suivante : si l'équipe augmente sa production
  de code mais continue à revoir les changements de la même façon, avec un
  contexte limité et peu de signaux de risque clairs, une partie de la vitesse
  gagnée en développement se perd en validation.
  Le SDD attaque ce point précis en donnant au reviewer l'intention explicite
  contre laquelle juger.
  Faire le lien avec le 4,6x de S4.
  Le rapport nomme le phénomène, et c'est le mot à poser à voix haute : la
  taxe de vérification (verification tax), c'est-à-dire la charge cognitive de
  relire et de valider du code généré par IA. L'effort économisé sur la frappe
  ne disparaît pas, il se relocalise en aval.
  Les trois formulations vérifiées, à reprendre telles quelles : la fréquence
  de déploiement n'augmente pas parce que le goulot s'est déplacé de l'écriture
  vers la revue ; le lead time ne baisse pas parce que le code arrive plus vite
  à la PR mais y attend plus longtemps ; le temps de revue augmente souvent
  alors même que le lead time total diminue, parce que les relecteurs font face
  à un volume croissant de code qu'ils n'ont pas écrit.
  C'est le sommet de la partie : ce slide nomme, avec la source la plus citée
  du domaine, exactement la thèse posée en S5, et il referme le 4,6x ET les
  +91 % de temps de revue affichés en S4. Faire le lien explicitement, en
  renvoyant à l'écran du début.
  Source : DORA 2026, The ROI of AI-assisted Software Development (Google).
---

:::grid{class="grid large-height center-align"}

:::col{class="s12"}

<h3>DORA 2026 lui a donné un nom : la <em>taxe de vérification</em>.</h3>

<p class="xlarge-text">L'effort économisé sur la frappe ne disparaît pas. Il se
relocalise en aval, sur la revue.</p>

:::

:::col{class="s4 left-align"}

<h6>La fréquence de déploiement ne monte pas</h6>

<p class="xlarge-text">Le goulot s'est déplacé de l'écriture vers la revue. On
produit plus, on ne livre pas plus souvent.</p>

:::

:::col{class="s4 left-align"}

<h6>Le lead time ne baisse pas</h6>

<p class="xlarge-text">Le code arrive plus vite à la PR. Et il y attend plus
longtemps.</p>

:::

:::col{class="s4 left-align"}

<h6>Le temps de revue, lui, augmente</h6>

<p class="xlarge-text">Souvent même quand le lead time total diminue : les
relecteurs encaissent un volume croissant de code qu'ils n'ont pas écrit.</p>

:::

:::col{class="s12"}

<blockquote class="left-align">
  <h5>C'est le 4,6&times; d'attente et les +91 % de temps de revue du début du talk, expliqués en un mot.</h5>
  <h5>Le SDD ne supprime pas cette taxe. Il donne au relecteur l'intention écrite contre laquelle juger.</h5>
</blockquote>

:::
:::

---
layout: default
title: "Les chiffres d'accélération, et leurs limites"
sources:
  - github.com/ianhxu/agentic-engineering-field-study
  - engineering.mercari.com
  - github.com/github/spec-kit
speakernotes: |
  S27 - Le meilleur point de données à l'échelle d'une équipe : Mercari, place
  de marché japonaise d'environ 22 millions d'utilisateurs mensuels, a publié
  sa méthodologie interne "Agent Spec-Driven Development" en décembre 2025.
  Après six mois, elle rapportait un gain de vitesse de 150 % sur sa baseline
  traditionnelle et de 80 % sur le prompting IA en format libre.
  Compléments : GitHub rapporte que les équipes utilisant Spec Kit livrent avec
  environ un ordre de grandeur moins de cycles de régénération from scratch ;
  AWS documente des cas où des fonctionnalités de 40 heures ont été livrées en
  moins de 8 heures de temps humain quand elles étaient d'abord rédigées comme
  specs.
  Dire les limites à voix haute : ce sont des chiffres auto-rapportés, sur un
  seul contexte d'équipe chacun, publiés par des acteurs qui ont intérêt au
  résultat. Aucun n'est un essai contrôlé. C'est cette honnêteté-là qui rend
  crédible tout ce qui précède, donc elle est aussi écrite à l'écran, pas
  seulement dite.
  Le chiffre AWS vient de cas clients documentés autour de Kiro : des
  fonctionnalités estimées à 40 heures livrées en moins de 8 heures de temps
  humain quand elles étaient d'abord rédigées comme specs.
  Sources : github.com/ianhxu/agentic-engineering-field-study
  (04-spec-driven-development.md, cas Mercari) ; GitHub, retours d'usage de
  Spec Kit ; AWS, cas clients Kiro. Vérifié le 11/09/2026.
---

:::grid{class="grid large-height center-align"}

:::col{class="s4"}

<h3><em>+150 %</em></h3>

<p class="xlarge-text">de vitesse<br/>sur la baseline traditionnelle</p>

<p class="xlarge-text">Et <strong>+80 %</strong> sur du prompting IA en format
libre. Méthodologie interne publiée, six mois de recul.</p>

<h6>Mercari, place de marché japonaise, 22 millions d'utilisateurs par mois</h6>

:::

:::col{class="s4"}

<h3><em>1 ordre de grandeur</em></h3>

<p class="xlarge-text">de régénérations<br/><strong>from scratch</strong> en
moins</p>

<p class="xlarge-text">Les équipes qui travaillent avec Spec Kit recommencent
bien moins souvent à zéro qu'avec du prompting ad hoc.</p>

<h6>GitHub, retours d'usage de Spec Kit</h6>

:::

:::col{class="s4"}

<h3><em>40 h &rarr; 8 h</em></h3>

<p class="xlarge-text">de temps humain<br/>sur la même fonctionnalité</p>

<p class="xlarge-text">Des fonctionnalités estimées à 40 heures, livrées en
moins de 8, quand elles sont d'abord rédigées comme specs.</p>

<h6>AWS, cas clients Kiro</h6>

:::

:::col{class="s12"}

<blockquote class="left-align">
  <h5>Trois chiffres auto-rapportés, sur un seul contexte d'équipe chacun, publiés par des gens qui ont intérêt au résultat. Aucun n'est un essai contrôlé.</h5>
  <h5>Je les donne quand même, parce qu'ils sont ce qu'on a de mieux aujourd'hui. Pas parce qu'ils prouvent quoi que ce soit.</h5>
</blockquote>

:::
:::

---
layout: default
title: Où le SDD n'accélère pas
sources:
  - dora.dev/ai/roi/report
speakernotes: |
  S28 - Indispensable pour la crédibilité, et c'est le slide qui fait gagner la
  salle.
  DORA 2026 mesure environ 35 à 40 % de gains sur des tâches simples en
  greenfield, mais 10 % ou moins pour des développeurs expérimentés sur du
  brownfield complexe. C'est le même écart que le talk décrit depuis le début :
  plus le contexte existant est lourd, moins la génération de code seule aide.
  Donc : petits changements, spikes exploratoires, domaine encore flou, legacy
  sans tests.
  Message : le SDD est un investissement dont le retour dépend de la taille du
  changement, et il passe mal à l'échelle vers le bas. Renvoyer explicitement à
  l'anecdote Kiro de S17 : quatre user stories et seize critères d'acceptation
  pour une correction de bug, c'est exactement ce cas-là.
  Source : DORA 2026, The ROI of AI-assisted Software Development (Google).
---

:::grid{class="grid large-height center-align"}

:::col{class="s6"}

<h3><em>+35 à 40 %</em></h3>

<p class="xlarge-text">sur des tâches simples,<br/>en <strong>greenfield</strong></p>

:::

:::col{class="s6"}

<h3><em>10 % ou moins</em></h3>

<p class="xlarge-text">pour des développeurs expérimentés,<br/>sur du
<strong>brownfield complexe</strong></p>

:::

:::col{class="s12"}

<h6>DORA 2026</h6>

<p class="xlarge-text">Petit changement, spike exploratoire, domaine encore
flou, legacy sans tests : la cérémonie coûte alors plus cher qu'elle ne
rapporte.</p>

<blockquote class="left-align">
  <h5>Le SDD est un investissement, et son retour dépend de la taille du changement.</h5>
  <h5>Il passe mal à l'échelle vers le bas : seize critères d'acceptation pour un bug d'une ligne, on l'a déjà vu tout à l'heure.</h5>
</blockquote>

:::
:::

---
layout: default
title: Quoi mesurer
sources:
  - dora.dev/ai/roi/report
speakernotes: |
  S29 - Pas les lignes de code.
  Lead time jusqu'à la prod, temps de revue, taux de rework, change failure
  rate, et surtout la part du rework due à une intention mal comprise : c'est
  la métrique que le SDD prétend améliorer.
  Chiffre à citer : seules 7,3 % des équipes ont un taux de rework sous 2 %, ce
  qui révèle une taxe cachée sur la productivité que la génération de code par
  IA peut facilement aggraver.
  Les deux repères qui éclairent ce 7,3 % et qu'il faut donner dans la foulée,
  sinon le chiffre ne dit rien : sous 10 % de rework, c'est la performance
  élite au sens DORA ; une organisation d'ingénierie typique se situe entre
  20 et 30 %.
  Insister sur la dernière ligne du tableau : la part du rework due à une
  intention mal comprise est la seule métrique que le SDD prétend directement
  améliorer. Les autres bougent par effet de bord.
  Source : DORA 2026, The ROI of AI-assisted Software Development (Google).
---

:::grid{class="grid large-height"}

:::col{class="s7 left-align"}

<h6>Ce qu'on regarde, à la place des lignes de code</h6>

<ul class="xlarge-text">
  <li><strong>Le lead time jusqu'à la prod</strong>, pas jusqu'à la PR.</li>
  <li><strong>Le temps de revue</strong>, compté à part : c'est là que la taxe
  se paie.</li>
  <li><strong>Le change failure rate</strong>, pour voir si la vitesse se prend
  sur la qualité.</li>
  <li><strong>Le taux de rework</strong>, et surtout la part due à une
  <em>intention mal comprise</em>. C'est la seule métrique que le SDD prétend
  améliorer directement.</li>
</ul>

:::

:::col{class="s5 left-align"}

<article class="round border">
  <h6>Le repère qui manque à presque tout le monde</h6>
  <ul class="xlarge-text">
    <li><strong>7,3 %</strong> des équipes seulement ont un taux de rework sous
    2 %.</li>
    <li>Sous <strong>10 %</strong>, c'est la performance élite.</li>
    <li>Une organisation d'ingénierie typique se situe entre <strong>20 et
    30 %</strong>.</li>
  </ul>
  <h6>DORA 2026</h6>
</article>

:::

:::col{class="s12 center-align"}

#### Le rework est une taxe cachée sur la productivité. Générer plus de code peut très vite l'*aggraver*.

:::
:::

---
layout: bare
class: transition
speakernotes: |
  Transition partie 5 - Retour au product engineer. 3 minutes, S30 à S33.
  Volontairement la plus courte du deck : deux lignes, et on enchaîne. Elle
  marque la fin de la partie technique, elle ne doit pas voler l'effet de S30.
  Ne rien ajouter à l'oral ici. Afficher, respirer, passer.
---

<h6 class="numero">Partie 5</h6>

<h2>Retour au product engineer</h2>

---
layout: default
title: "Boucle fermée : l'artefact du product engineer"
class: main responsive large-height
speakernotes: |
  S30 - La résolution de S5. Même layout, même phrase, même mise en forme :
  la salle doit voir la boucle se fermer sans qu'on l'explique.
  1. Le product engineer avait besoin d'un artefact pour porter l'intention. On
  vient de passer 35 minutes à décrire cet artefact et son outillage.
  2. Relire la phrase charnière. C'est le même bloc qu'en S5, au mot près.
  Marquer le même temps d'arrêt qu'au début.
  3. La nouvelle pile de compétences : cadrage de problème, modélisation de
  domaine, écriture de critères vérifiables, arbitrage, design de la
  vérification. Et la vitesse de frappe compte moins.
  Ne pas résumer le talk ici : les slides 31 à 33 portent la sortie.
---

<div class="large-space"></div>

### Le product engineer avait besoin d'un artefact pour porter l'intention.

#### On vient de passer 35 minutes à décrire cet artefact, et son outillage.

<div class="large-space"></div>

<blockquote class="charniere">
  <h4>Sans SDD, un product engineer n'est qu'un dev surchargé à qui on a retiré son PM.</h4>
  <h5>Il a besoin d'un artefact pour porter l'intention jusqu'à la machine. Cet artefact, c'est la <em>spec</em>.</h5>
</blockquote>

<div class="large-space"></div>

<div class="grid center-align">
  <div class="s2"><h6>Cadrer un problème</h6></div>
  <div class="s3"><h6>Modéliser un domaine</h6></div>
  <div class="s3"><h6>Écrire des critères vérifiables</h6></div>
  <div class="s2"><h6>Arbitrer</h6></div>
  <div class="s2"><h6>Concevoir la vérification</h6></div>
  <div class="s12"><h5><em>La vitesse de frappe compte moins.</em></h5></div>
</div>

---
layout: default
title: "Les risques, honnêtement"
sources:
  - cio.com/article/4190086
  - letalentclub.substack.com
  - productengineer.org
speakernotes: |
  S31 - Tous les ingénieurs ne peuvent pas devenir de bons product engineers :
  le rôle exige profondeur technique, intuition produit, communication,
  conscience business et forte autogestion.
  L'erreur la plus dangereuse : déléguer l'autorité produit trop tôt, sans
  supervision suffisante ni maturité organisationnelle.
  Ajouter : atrophie des compétences chez les juniors, capacité de revue,
  responsabilité en cas d'incident.
  Rappeler l'avertissement de la product engineer française (Flowie, arrivée
  par la voie PM) : beaucoup de gens qui veulent faire du produit veulent faire
  de la stratégie et prendre des décisions, alors qu'au quotidien c'est surtout
  de la delivery. Ne devenez pas product engineer pour faire de la stratégie.
  Ce n'est pas la mort du PM : le manifeste s'adresse directement aux designers
  dans une lettre ouverte pour désamorcer l'idée d'un empiètement, et le rôle
  reste compatible avec un PM qui garde du recul sur la priorisation globale.
  Le risque le mieux formulé de tous, et c'est pour ça qu'il est en haut du
  slide : un bon product engineer exige un cadre solide autour de lui, des
  processus de release disciplinés, des frontières de responsabilité claires,
  une infrastructure de test fiable et un leadership technique expérimenté.
  Sans ce cadre, le rôle ne tient pas, quelle que soit la personne.
  Le risque générationnel, à développer à l'oral : les plus exposés sont les
  profils mi-carrière, exactement la population qui porte la connaissance
  institutionnelle et les meilleurs réflexes produit. Et le ralentissement des
  embauches juniors casse le transfert de compétences, restructure le pipeline
  de talents et enferme le recrutement sur des profils seniors, plus chers et
  plus disputés.
  La question "donc on supprime les PM ?" tombera, elle n'est pas à l'écran :
  répondre que non, le Product Engineer Manifesto s'adresse même directement
  aux designers dans une lettre ouverte pour désamorcer l'idée d'un
  empiètement, et que le rôle reste compatible avec un PM qui garde du recul
  sur la priorisation globale.
  Sources : CIO, The rise of the product engineer ; Le Talent Club ; Product
  Engineer Manifesto.
---

:::grid{class="grid large-height center-align"}

:::col{class="s12"}

<h3>Un bon product engineer a besoin d'un <em>cadre solide autour de lui</em>.</h3>

<p class="xlarge-text">Des releases disciplinées, des frontières de
responsabilité claires, une infra de test fiable, un leadership technique
expérimenté. Sans ça, le rôle ne tient pas.</p>

:::

:::col{class="s4 left-align"}

<h6>Tout le monde ne peut pas l'être</h6>

<p class="xlarge-text">Profondeur technique, intuition produit, communication,
conscience business, forte autogestion. Ça fait beaucoup pour une seule
personne, et ce n'est pas une question de bonne volonté.</p>

:::

:::col{class="s4 left-align"}

<h6>Le risque générationnel</h6>

<p class="xlarge-text">Les plus exposés sont les profils mi-carrière, ceux qui
portent la connaissance institutionnelle. Et moins d'embauches juniors, c'est
le transfert de compétences qui casse.</p>

:::

:::col{class="s4 left-align"}

<h6>L'erreur la plus coûteuse</h6>

<p class="xlarge-text">Déléguer l'autorité produit trop tôt, sans supervision
suffisante ni maturité organisationnelle. Le titre arrive avant le cadre, et
c'est l'équipe qui paie.</p>

:::

:::col{class="s12"}

<blockquote class="left-align">
  <h5>« Beaucoup de gens qui veulent faire du produit veulent faire de la stratégie et prendre des décisions. Au quotidien, c'est surtout de la <em>delivery</em>. »</h5>
  <h5>Ne devenez pas product engineer pour faire de la stratégie.</h5>
</blockquote>

:::
:::

---
layout: content
title: Lundi matin
speakernotes: |
  S32 - Cinq actions par coût croissant.
  1. Écrire un AGENTS.md / constitution.md pour un repo.
  2. Utiliser le mode plan systématiquement pendant une semaine.
  3. Faire une vraie feature avec OpenSpec.
  4. Versionner les specs dans git à côté du code.
  5. Instaurer un seul gate humain - validation de spec avant implémentation -
  et mesurer le temps de revue avant/après.
  À produire : un AGENTS.md / constitution.md d'exemple, court, lisible en 20
  secondes à l'écran. Il ne tient pas sur ce slide : celui-ci ne porte que la
  liste des cinq actions, parce que c'est le slide que la salle photographie et
  qu'il doit rester lisible en une seconde. Si l'exemple de fichier est fait,
  il va sur un slide à lui, juste après.
  Dire la phrase de sortie à voix haute : aucune de ces cinq actions ne demande
  d'installer un framework. C'est le même message qu'en S24.
---

<table class="border xlarge-text">
  <thead>
    <tr>
      <th>Ce que tu peux faire lundi</th>
      <th>Ce que ça coûte</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><strong>1.</strong> Écrire un <code>AGENTS.md</code> (ou une <code>constitution.md</code>) pour un seul repo.</td>
      <td>Une demi-heure, une fois.</td>
    </tr>
    <tr>
      <td><strong>2.</strong> Passer par le mode plan systématiquement, pendant une semaine.</td>
      <td>Rien à installer. Une habitude à prendre.</td>
    </tr>
    <tr>
      <td><strong>3.</strong> Faire une vraie fonctionnalité avec OpenSpec.</td>
      <td>Une demi-journée, sur quelque chose qui compte.</td>
    </tr>
    <tr>
      <td><strong>4.</strong> Versionner les specs dans git, à côté du code.</td>
      <td>Une décision d'équipe, et une revue de plus.</td>
    </tr>
    <tr>
      <td><strong>5.</strong> Instaurer un seul gate humain : la spec est validée avant l'implémentation. Puis mesurer le temps de revue avant et après.</td>
      <td>Un accord d'équipe, et un chiffre à regarder dans un mois.</td>
    </tr>
  </tbody>
</table>

<p class="xlarge-text center-align">Aucune de ces cinq actions ne demande
d'installer un framework. Elles sont rangées par coût croissant : <em>commence
par la première</em>.</p>

---
layout: default
title: La suite de la demi-journée
sources:
  - medium.com/@damien.gouron
speakernotes: |
  S33 - Ce qui sera fait en atelier cet après-midi, et les 3 questions laissées
  ouvertes.
  Rappeler que la démo n'est pas dans ce talk : elle est dans l'atelier.
  Ressource à partager après le talk (FR) : medium.com/@damien.gouron,
  Spec-Driven Development, le guide pratique.
  Puis les coordonnées et les questions.
  À COMPLÉTER par le speaker : le contenu exact de l'atelier n'est pas décrit
  ici, la colonne de gauche reste volontairement générique. Y mettre le repo de
  départ, le format et ce que la salle repart avec.
  Les trois questions ouvertes affichées sont reprises du talk lui-même, pas
  ajoutées : la dérive spec/code laissée sans solution en S17, le niveau 3 de
  Böckeler encore expérimental en S12, et le transfert de compétences abordé en
  S31. Si ce ne sont pas les trois retenues, les remplacer, mais garder des
  questions déjà posées dans le talk : la salle doit les reconnaître.
  Terminer là-dessus, sans revenir sur le plan du talk.
---

:::grid{class="grid large-height"}

:::col{class="s6 left-align"}

<h6>Cet après-midi, l'atelier</h6>

<p class="xlarge-text">La démo n'est pas dans ce talk, elle est là-bas. On
passe du support à du vrai code, et on écrit des specs pour de bon.</p>

<h6>À lire ensuite, en français</h6>

<p class="xlarge-text">Damien Gouron, <em>Spec-Driven Development : le guide
pratique</em>, sur Medium.</p>

:::

:::col{class="s6 left-align"}

<h6>Trois questions que je laisse ouvertes</h6>

<ul class="xlarge-text">
  <li>La dérive spec/code : aucun outil ne réconcilie tout seul. Qui tient les
  specs à jour dans six mois ?</li>
  <li>Le niveau 3 : est-ce qu'on ira vraiment vers la spec comme seule source,
  ou est-ce que ça restera une vision ?</li>
  <li>Si les juniors écrivent moins de code, où apprennent-ils à juger celui
  des agents ?</li>
</ul>

:::

:::col{class="s12 center-align"}

### Merci. Et maintenant, vos questions.

##### @ldussart &middot; ldussart.bsky.social &middot; Ludovic Dussart sur LinkedIn &middot; blog.zatsit.fr

:::
:::
