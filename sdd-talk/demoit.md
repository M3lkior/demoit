---
layout: cover
speakernotes: |
  S1 - Titre + toi. Qui je suis, pourquoi je parle de ça, la promesse de la
  demi-journée.
  Annoncer le format : 45 minutes d'introduction, puis atelier.
  Thèse unique à poser dès maintenant : l'IA a rendu le code abondant, le
  goulot s'est déplacé en amont (l'intention) et en aval (la vérification).
---

<!-- Corps volontairement vide : le titre, le sous-titre et le logo viennent de
     talk.yml, et le layout cover les rend lui-même. -->

---
layout: default
title: Je suis **Ludovic Dussart**
class: slide-main slide-wide slide-prose flex flex-col justify-center text-center
---
<div class="grid grid-cols-12 items-center gap-4">
  <div class="col-span-6">
    <div>
      <img class="rounded-card" style="block-size: 15rem; " src="/images/me.jpg">
    </div>
  </div>
  <div class="col-span-6">
    <div class="text-center">
      <h4>Solutions architect <em>@zatsit</em></h4>
      <h4><em>AsyncAPI</em> maintainer</h4>
      <h4>Open Source <em>fanatic</em></h4>
    </div>
  </div>
  <div class="col-span-12"><h2 class="text-left"><em>Où me trouver ?</em></h2></div>
  <!-- contact-row porte la colonne (span 4), la ligne flex et l'avatar rond :
       c'est la paire s1 + s3 left-align d'avant, en une seule classe. -->
  <div class="contact-row"><img src="/images/bsky.png"/><h5>ldussart.bsky.social</h5></div>
  <div class="contact-row"><img src="/images/linkedin.png"/><h5>Ludovic Dussart</h5></div>
  <div class="col-span-12"><h2 class="text-left"><em>Où nous trouver ?</em></h2></div>
  <div class="contact-row"><img src="/images/web.png"/><h5>https://zatsit.fr</h5></div>
  <div class="contact-row"><img src="/images/bsky.png"/><h5>zatsit.bsky.social</h5></div>
  <div class="contact-row"><img src="/images/blog.png"/><h5>https://blog.zatsit.fr</h5></div>
</div>
<div class="h-10"></div>

<h6 class="quote">Engager notre <em>expertise numérique</em> au service de <em>l'impact des entreprises</em>, en créant un écosystème <em>durable</em>, <em>partenarial</em> et <em>positif</em></h6>



---
layout: quote
title: Disclamer
---


<h4>L'écosystème IA évolue bien trop vite pour que l'humain puisse suivre et tout maitriser.</h4>

---
layout: bare
class: slide-main slide-hero flex flex-col items-center justify-center text-center
speakernotes: |
  Transition partie 0 - Un nouveau métier, et son problème. 7 minutes, S2 à S5.
  Enchaîner vite, 5 secondes à l'écran : la transition sert de repère à la
  salle, pas de temps de parole.
  Ce qu'on promet ici et qu'il faudra tenir : on part du personnage, on montre
  son problème, et le problème débouche sur le SDD en S5.
---

<h1 class="mb-0 text-6xl leading-tight font-bold tracking-tight">Le <em>Product Engineer</em></h1>

<p class="text-fg-muted mt-6 text-3xl font-medium">Nouveau métier, ou évolution d'un métier existant ?</p>

---
layout: default
title: "Product Engineer : le métier que le marché a inventé deux fois"
class: slide-main slide-wide slide-prose flex flex-col justify-center
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

### *2018 &ndash; 2024*

<p class="text-[1.6rem]">Né dans des scale-ups produit, pour une raison
d'organisation : le <em>coût des relais</em> dépasse leur bénéfice quand
l'équipe est petite et que la boucle de feedback doit rester courte.</p>

<ul class="text-[1.6rem]">
  <li><strong>PostHog</strong>: un développeur directement responsable de l'amélioration du produit, ayant un avis sur la roadmap, capable de prendre seul des décisions de design, faisant du support utilisateur directement, moins attaché aux fonctionnalités qu'il a ajoutées par le passé, conscient de la place de son travail dans la stratégie, et motivé avant tout par l'impact et les résultats. 187 personnes pour 47 équipes (4,2 personnes en moyenne)</li>
  <li><strong>incident.io</strong> : des devs attachés au résultat et à l'impact, pas à l'implémentation</li>
  <li><strong>JM. Lemieux</strong>, ex-VP Eng Shopify : des ingénieurs qui ont soif d'utiliser la technologie pour court-circuiter les problèmes humains et utilisateurs.</li>
  <li><strong>Product Engineer Manifesto</strong>, 2024 : c'est la responsabilité des builders de chercher d'abord à <em>comprendre le problème</em> avant de plonger dans les solutions, <em>et de s'occuper des domaines design, technique et business en prenant une part active dans chacun.</em></li>
</ul>

---
layout: default
title: "Product Engineer : le métier que le marché a inventé deux fois"
class: slide-main slide-wide slide-prose flex flex-col justify-center
sources:
  - sfeir.com/concepts/product-engineer
  - turingcollege.com/blog/rise-of-the-ai-product-engineer
speakernotes: |
    Punchline : ce rôle n'a pas été créé par l'IA, il a été révélé par elle.

    - **Le "one-person band"** : le Chief AI Officer de Pendo décrit le glissement : traditionnellement un ingénieur livre, un product manager parle aux utilisateurs pour voir si c'était viable, et la boucle continue ; en se rapprochant du product engineer, on obtient une seule personne qui livre, itère, collecte le feedback et itère de nouveau, au lieu de répartir ça entre engineering, produit et design. Et parce qu'ils passent moins de temps à écrire du code grâce à Codex, Claude Code ou l'agent de leur choix, ce rôle se concentre sur le travail à plus forte valeur : la résolution créative de problèmes et les échanges avec les utilisateurs.
    - **L'argument économique (le plus solide)** : la ressource rare n'est plus l'exécution, c'est le jugement : savoir quoi construire, et savoir si ce qu'on a construit est bon. La rémunération monte pour ceux qui construisent concrètement et baisse pour ceux qui ne font que coordonner. Le rôle se situe à l'intersection de quatre disciplines : ingénierie logicielle/IA, product management, UX/UI et jugement sur la donnée.
---

:::grid{class="grid grid-cols-12 gap-4 h-stage-large"}

:::col{class="col-span-12 text-center"} 
:::
:::col{class="col-span-12 text-center"}
#### Ce rôle n'a pas été créé par l'IA. Il a été *révélé* par elle.
:::

:::col{class="col-span-12"}

### *2025 &ndash; 2026*

<p class="text-[1.6rem]">Le même rôle passe de
<em>niche culturelle</em> à <em>norme</em>, parce que l'IA crée des goulots en amont des phases de production de codes, et plus seulement dans les petites équipes.</p>

<ul class="text-[1.6rem]">
  <li>SFEIR positionne le product engineer comme le rôle central <em>de l'ère 10x</em>, qui ne se définit plus par une spécialité technique unique mais par sa capacité à coordonner toute la chaîne de production logicielle augmentée par l'IA.</li>
 
</ul>
:::


:::

---
layout: default
title: "Product Engineer : la définition de Sfeir"
class: slide-main slide-wide slide-prose flex flex-col justify-center
sources:
  - sfeir.com/concepts/product-engineer
speakernotes: |
---

<div class="text-center">
  <a href="https://sfeir.com/concepts/product-engineer" target="_blank">
    <img src="/images/sfeir_product_engineer.png" alt="Product Engineer" />
  </a>
</div>

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

:::grid{class="grid grid-cols-12 gap-4 h-stage-large text-center"}
:::col{class="col-span-4 panel" reveal}

##### La cause

<p class="text-[1.6rem]">La spécialisation était une réponse à la <em>complexité</em>.</p>

:::

:::col{class="col-span-4 panel" reveal}

##### Le changement

<p class="text-[1.6rem]">L'IA absorbe une partie de la complexité d'<em>exécution</em>.</p>

:::

:::col{class="col-span-4 panel" reveal}

##### La conséquence

<p class="text-[1.6rem]">Donc la <em>division du travail</em> perd sa raison d'être.</p>

:::

:::col{class="col-span-12"}

<div class="h-8"></div>

<blockquote class="text-left">
  <h5 class="reveal"><em>Aujourd'hui (Pizza Team)</em> : les <em>PO/PM</em> écrivent les besoins, <em>un ingénieur</em> produit le code et livre, un <em>PO/PM</em> va voir les utilisateurs, la boucle continue.</h5>
  <h5 class="reveal"><em>Demain (Sandwich Team)</em>: <em>une seule personne</em> livre, itère, collecte le feedback, et itère de nouveau.</h5>
</blockquote>

:::

:::


---
layout: default
title: Le vrai goulot n'est pas le code
class: slide-main slide-narrow slide-prose flex flex-col justify-center
sources:

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

<div class="h-12"></div>

### Le temps de développement est rarement le vrai frein.

##### Les goulots sont en amont : des exigences floues, des décisions produit faibles, des expressions vagues.

<div class="h-8"></div>

#### Une IA qui rend le code moins cher ne répare rien de ça. Elle l'*expose*.

<div class="h-12"></div>

<blockquote class="quote charniere reveal">
  <h4>Sans SDD, un product engineer n'est qu'un dev augmenté à qui on a retiré son PM.</h4>
  <h5>Il a besoin d'un artefact pour porter l'intention. Cet artefact, c'est la <em>spec</em>.</h5>
</blockquote>

---
layout: bare
class: slide-main slide-hero flex flex-col items-center justify-center text-center
speakernotes: |
  Transition partie 1 - Les prérequis. 8 minutes, S6 à S10.
  Elle arrive juste après la charnière S5 : laisser retomber la phrase avant
  d'afficher celle-ci.
  Ce qu'on promet : avant de parler SDD, quatre choses à avoir en tête -
  où en est chacun dans la délégation, comment marche vraiment un agent,
  pourquoi le vibe coding casse, et pourquoi le DDD revient.
---

<h1 class="mb-0 text-6xl leading-tight font-bold tracking-tight">Avant le <em>SDD</em></h1>

<p class="text-fg-muted mt-6 text-3xl font-medium">Faisons le point</p>

---
layout: default
title: L'échelle de délégation
sources:
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

:::grid{class="grid grid-cols-12 gap-4 h-stage-large text-center"}

:::col{class="col-span-3 panel text-center" reveal}

##### 1. Autocomplétion

<p class="text-[1.6rem]">On écrit, l'outil finit la ligne.</p>

<h6>On délègue la frappe.</h6>

:::

:::col{class="col-span-3 panel text-center" reveal}

##### 2. Chat

<p class="text-[1.6rem]">On décrit une fonction, on la relis, on la colle.</p>

<h6>On délègue un bout de code.</h6>

:::

:::col{class="col-span-3 panel text-center" reveal}

##### 3. Agent

<p class="text-[1.6rem]">On donne une tâche. L'agent lit le dépôt, édite, lance les tests.</p>

<h6>On délègue une tâche.</h6>

:::

:::col{class="col-span-3 panel text-center" reveal}

##### 4. Agent autonome

<p class="text-[1.6rem]">On donne un objectif. L'agent découpe le travail lui-même et travaille en quasi-autonomie.</p>

<h6>On délègue l'intention.</h6>

:::

:::col{class="col-span-12"}

<div class="h-8"></div>

<h4 class="reveal"> Plus on monte, moins l'intention tient dans la tête. Il faut l'<em>écrire</em>.</h4>



<h5 class="mb-0 text-6xl leading-tight font-bold tracking-tight reveal">Vous êtes où, aujourd'hui ?</h1>

:::
:::

---
layout: default
title: Comment fonctionne un agent
sources:
  
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

:::grid{class="grid grid-cols-12 gap-4 h-stage-large text-center"}

:::col{class="col-span-6" reveal}

##### La fenêtre de contexte, c'est de la *RAM*

<p class="text-[1.6rem]">Petite, rapide, et vidée à la fin de la session. Tout ce que l'agent sait du projet tient là-dedans, et seulement là.</p>

:::

:::col{class="col-span-6" reveal}

##### Le dépôt, c'est le *disque*

<p class="text-[1.6rem]">Grand, lent, persistant. L'agent n'y lit que ce qu'on lui dit d'aller chercher, et il repart de zéro à chaque fois.</p>

:::

:::col{class="col-span-12"}

<div class="reveal text-center">
  <h5>Et dans la fenêtre, <em>rien n'est garanti</em> : être présent ne veut pas dire être retrouvé.</h5>
  <h5>Un prompt très long n'est pas une mémoire.</h5>
</div>

:::

:::col{class="col-span-12" reveal}

#### Pas de mémoire persistante ? <em>Alors on l'écrit</em>.

<p class="text-[1.6rem]"><strong>AGENTS.md</strong>, <strong>constitution.md</strong>, <strong>spec.md</strong>,
<strong>plan.md</strong>, <strong>PRD.md</strong> : des fichiers durables et versionnés, relus à chaque
session. </p>

::: 
:::col{class="col-span-12" reveal}

<h3>On parle de <em>Context Engineering</em>, et c'est sur cela que s'appuie le <em>SDD</em>.</p>

:::
:::

---
layout: default
title: Les modes de défaillance du vibe coding
sources:
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

:::grid{class="grid grid-cols-12 gap-4 h-stage-large text-center"}

:::col{class="col-span-12"}

<div class="text-center">
  <h5>On demandait un petit fix. </h5>
  <h6>On récupère un refactor, trois fichiers en plus, une logique « améliorée », et des tests qui passent parce qu'ils ne testent rien.</h5>
</div>

:::

:::col{class="col-span-4" reveal}

##### <em>L'intention vit dans le chat</em>

<p class="text-[1.6rem]">On ferme l'onglet, il ne reste que du code. Plus personne ne peut dire ce qui avait été demandé.</p>

:::

:::col{class="col-span-4" reveal}

##### <em>L'agent dérive</em>

<p class="text-[1.6rem]">Du code que personne n'a demandé, des décisions que personne n'a prises.</p>

:::

:::col{class="col-span-4" reveal}

##### <em>Rien n'est vérifiable</em>

<p class="text-[1.6rem]">Sans critères d'acceptation écrits avant, « ça a l'air bon » tient lieu de recette.</p>

:::

:::col{class="col-span-12" reveal}

#### Ces trois problématiques, le SDD y réponds.

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

:::grid{class="grid grid-cols-12 gap-4 h-stage-large"}


:::col{class="col-span-12 text-center"}

#### On a toujours écrit des specs. 
#### Ce qui change : elles doivent être *relisibles par une machine*, et versionnées.

:::

:::col{class="col-span-12 text-center" reveal}

<ul class="text-[1.6rem] reveal text-left">
  <li><strong>User story + critères d'acceptation</strong> : dire ce qu'on attend, avant de le construire.</li>
  <li><strong>Gherkin, BDD</strong> : la règle métier écrite dans une forme que la machine relit.</li>
  <li><strong>TDD</strong> : le critère d'abord, le code ensuite.</li>
  <li><strong>Contract-first, OpenAPI, AsyncAPI</strong> : le contrat fait foi, l'implémentation suit.</li>
  <li><strong>ADR</strong> : la décision et sa raison, versionnées avec le code.</li>
  <li><strong>DDD</strong> : les mots du métier, partagés par toute l'équipe.</li>
  <li><strong>C4Model</strong> : la représentation de l'architecture.</li>
</ul>

:::

:::

---
layout: default
title: "DDD : pourquoi il contribue au SSD ?"
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

:::grid{class="grid grid-cols-12 gap-4 h-stage-large text-center"}

:::col{class="col-span-4" reveal}

##### <em>Langage ubiquitaire</em>

<p class="text-[1.6rem]">Un mot, un sens, partout. L'agent cherche les noms qu'on lui donne : deux concepts qui portent le même mot finissent fusionnés.</p>

:::

:::col{class="col-span-4" reveal}

##### <em>Bounded contexts</em>

<p class="text-[1.6rem]">Des frontières écrites. Elles disent à l'agent où s'arrête son périmètre, et elles l'aident à découper son travail en tranches parallélisables.</p>

:::

:::col{class="col-span-4" reveal}

##### <em>Invariants</em>

<p class="text-[1.6rem]">Les règles que le domaine ne tolère pas de casser. Ce sont déjà des critères d'acceptation et de validations.</p>

:::

:::col{class="col-span-12"}

<h5 class="reveal"">Une spec générée que personne ne lit et ne comprends, ne résoudra rien.</h5>

:::

:::col{class="col-span-12" reveal}

#### Connaître son domaine vaut désormais plus que connaître un *framework*.

:::
:::

---
layout: bare
class: slide-main slide-hero flex flex-col items-center justify-center text-center
speakernotes: |
  Transition partie 2 - Le SDD et son intérêt. 10 minutes, S11 à S17. C'est le
  cœur du talk.
  Ce qu'on promet : la définition, les trois niveaux d'ambition, la boucle,
  ce qui fait une bonne spec, pourquoi ça marche mécaniquement - et le coût,
  honnêtement.
---

<h1 class="mb-0 text-6xl leading-tight font-bold tracking-tight">Le <em>spec-driven development</em></h1>

<p class="text-fg-muted mt-6 text-3xl font-medium">Deep dive</p>

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

:::grid{class="grid grid-cols-12 gap-4 h-stage-large text-center"}

:::col{class="col-span-12"}

### La spec devient un *contrat exécutable*, dont les agents dérivent le code.

:::

:::col{class="col-span-6" reveal}

##### <em>La documentation</em>

<p class="text-[1.6rem]">Elle décrit ce qu'on a fait. Elle est passive : rien ne se
casse quand le code s'en éloigne, et personne ne s'en aperçoit.</p>

:::

:::col{class="col-span-6" reveal}

##### <em>La spec</em>

<p class="text-[1.6rem]">Elle dit ce qu'on attend. Elle est appliquée : la dérive
se détecte par l'outillage.</p>

:::

:::col{class="col-span-12" reveal}

#### *La spec devient le prompt.*

<p class="text-[1.6rem]">Ce qu'on écrit une fois sert trois fois : à demander,
à revoir, à vérifier.</p>

:::
:::

---
layout: default
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

<!-- Le gabarit des cinq tableaux du deck. beercss donnait inline-size:100%,
     border-spacing:0 et 0.5rem de padding de cellule : le chassis les porte
     maintenant. Restent le corps de 1.6rem (xlarge-text) et le filet sous
     chaque ligne sauf la dernière (table.border), redessiné par les variantes
     [&_...] ci-dessous. -->
<table class="w-full border-collapse text-[1.6rem] [&_thead_th]:border-b [&_thead_th]:border-outline [&_tbody_tr:not(:last-child)_td]:border-b [&_tbody_tr:not(:last-child)_td]:border-outline">
  <thead>
    <tr>
      <th>Niveau</th>
      <th>Ce qui arrive à la spec</th>
      <th>Ce qu'on maintient</th>
      <th>Maturité</th>
    </tr>
  </thead>
  <tbody>
    <tr class="reveal">
      <td><strong>1. Spec-first</strong></td>
      <td>Écrite avant le code, elle pilote le premier passage de l'agent. Puis on la laisse tomber.</td>
      <td>Le code</td>
      <td>Là où sont la plupart des équipes qui démarrent.</td>
    </tr>
    <tr class="reveal">
      <td><strong>2. Spec-anchored</strong></td>
      <td>Elle survit à la livraison : c'est le document vivant de la fonctionnalité, itération après itération.</td>
      <td>Le code <em>et</em> la spec</td>
      <td>La cible réaliste pour 2026. C'est ce qui doit être fait aujourd'hui.</td>
    </tr>
    <tr class="reveal">
      <td><strong>3. Spec-as-source</strong></td>
      <td>Tout le code est généré, marqué comme généré, et jamais édité à la main.</td>
      <td>La spec, et rien d'autre</td>
      <td>Vision long terme, encore largement expérimentale.</td>
    </tr>
  </tbody>
</table>


<p class="text-[1.6rem] text-center reveal">Le niveau 3, c'est ce que les compilateurs
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

:::grid{class="grid grid-cols-12 gap-4 h-stage-large"}

:::col{class="col-span-12 text-left"}

<ul class="text-[2rem]">
  <li><strong>Le résultat attendu</strong> : ce qui doit être vrai quand c'est fini.</li>
  <li><strong>Les limites du périmètre</strong> : ce qu'on ne touche pas, et qu'on ne veut pas voir bouger.</li>
  <li><strong>Les contraintes</strong> : performance, sécurité, compatibilité, budget.</li>
  <li><strong>Les décisions déjà prises</strong> : et la raison pour laquelle on ne les rouvre pas.</li>
  <li><strong>La découpe</strong> : des tâches assez petites pour être vérifiées une par une.</li>
  <li><strong>Les critères de vérification</strong> : écrits <em>avant</em>, jamais après.</li>
</ul>

:::



:::col{class="col-span-12 text-center" reveal}

#### Tout ce qui n'est pas écrit, l'agent l'*invente*.

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

:::grid{class="grid grid-cols-12 gap-4 h-stage-large text-center"}

:::col{class="col-span-6" reveal}

##### En amont : l'*intention*

<ul class="text-[1.6rem] text-left">
  <li>La spec porte l'état que le modèle n'a pas. Le travail redevient reprenable d'une session à l'autre, et transférable d'un agent à l'autre.</li>
  <li>Des contrats écrits et des périmètres séparés : plusieurs agents avancent en même temps sans se marcher dessus.</li>
</ul>

:::

:::col{class="col-span-6" reveal}

##### En aval : la *vérification*

<ul class="text-[1.6rem] text-left">
  <li>Les critères sont écrits avant, donc la sortie se juge au lieu de se deviner. Et celui qui vérifie n'est plus celui qui implémente.</li>
  <li>Violation d'architecture, contrat d'API qui glisse : un test unitaire ne peut structurellement pas les voir. Une spec, si.</li>
</ul>

:::

:::col{class="col-span-12"}

<blockquote class="text-left reveal">
  <h5><em>La revue remonte d'un cran</em> : 200 lignes de markdown qu'on a écrites, au lieu de 2 000 lignes de diff qu'on n'a pas écrites.</h5>
</blockquote>

:::

:::col{class="col-span-12" reveal}

#### Aucun de ces gains ne vient du modèle. Ils viennent du fait d'avoir écrit.

:::
:::

---
layout: default
title: "Le (contre) coût du SSD ?"
sources:
  - martinfowler.com/articles/exploring-gen-ai/sdd-3-tools.html
  - dora.dev/ai/roi/report
  - coderabbit.ai/blog/state-of-ai-vs-human-code-generation-report
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

:::grid{class="grid grid-cols-12 gap-4 h-stage-large text-center"}

:::col{class="col-span-6" reveal}

<h3><em>4 stories, 16 critères</em></h3>

<p class="text-[1.6rem]">pour <strong>un seul petit bug</strong></p>

<p class="text-[1.6rem]">La cérémonie ne se dimensionne pas toute seule. Pour un
null check d'une ligne, on ne touche pas à la constitution.</p>

<h6><em>Birgitta Böckeler, en lançant Kiro</em></h6>

:::

:::col{class="col-span-6" reveal}

<h3><em>1,7x</em></h3>

<p class="text-[1.6rem]">plus de bugs sur les PRs générées par IA. C'est ce que le SDD cherche à éviter, pas ce qu'il
garantit d'éviter.</p>

<h6><em>CodeRabbit</em></h6>

:::

:::col{class="col-span-12" reveal}

<blockquote class="text-left">
  <h5>Le SDD est <em>peu adapté pour de petite tâche</em>.</h5>
  <h5>La <em>dérive spec/code est réélle</em>. Aucun outil ne réconcilie cela tout seul : il faut le déclencher, et relire la sortie.</h5>
</blockquote>

:::
:::

---
layout: bare
class: slide-main slide-hero flex flex-col items-center justify-center text-center
speakernotes: |
  Transition partie 3 - Le panorama des frameworks. 10 minutes, S18 à S24.
  Le piège n°1 du talk est ici : le catalogue d'outils. Le dire à voix haute en
  affichant cette slide - "ce n'est pas un comparatif d'outils, c'est une
  démonstration que tous font la même chose".
  Si la salle est plus junior que prévu, c'est ici qu'on coupe : fusionner
  S19 à S23 en un seul tableau et récupérer 3 minutes pour les parties 1 et 2.
---

<h1 class="mb-0 text-6xl leading-tight font-bold tracking-tight">Le panorama des <em>frameworks</em></h1>

<p class="text-fg-muted mt-6 text-3xl font-medium">Plus de <strong>30 outils</strong> recensés début 2026.</p>

---
layout: default
title: "La base : le mode plan"
speakernotes: |
  S18 - 
  Ensuite la baseline : le mode plan (Claude Code, Cursor). Plan éphémère,
  aucun artefact persisté, aucun gate, zéro installation. Excellent pour une
  tâche de 30 minutes, insuffisant dès qu'il y a plusieurs sessions, plusieurs
  agents ou une revue par un tiers.
  C'est le témoin de comparaison sur les 4 slides suivants - le dire
  explicitement, sinon le panorama devient un catalogue.
---

:::grid{class="grid grid-cols-12 gap-4 h-stage-large"}

:::col{class="col-span-12 text-left" reveal}


  
  <p class="text-[2rem]"><strong>Ce qu'il fait</strong> : l'agent écrit son
  plan, on le lis, on le corrige, il exécute. Built-in dans les outils.</p>
  <p class="text-[2rem]"><strong>Où il s'arrête</strong> : le plan est
  éphémère. Rien n'est persisté, rien n'est versionné.</p>

:::

:::col{class="col-span-12 text-center" reveal}

#### <em>Parfait pour une tâche de 30 minutes.</em>
#### Insuffisant dès qu'il y a plusieurs sessions, plusieurs agents, ou une revue par un tiers.

:::
:::

---
layout: default
title: GitHub Spec Kit
github: github/spec-kit
sources:
  - github.com/github/spec-kit
  - ranthebuilder.cloud/blog/i-tested-three-spec-driven-ai-tools-here-s-my-honest-take
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

:::grid{class="grid grid-cols-12 gap-4 h-stage-large"}

:::col{class="col-span-7 text-left" reveal}

<h6><strong>Principe</strong></h6>

<ul class="text-[1.6rem]">
  <li><strong>Trois commandes, trois fichiers qui restent</strong> :
  <code>specify</code> pour le métier et les critères de succès,
  <code>plan</code> pour les décisions d'architecture, <code>tasks</code> pour
  la découpe en unités testables.</li>
  <li><strong>Une constitution</strong> écrite une fois pour le projet, dont
  chaque spec hérite ensuite.</li>
  <li><strong>Les inconnues sont mises en avant</strong> : les
  templates posent un <code>NEEDS CLARIFICATION</code> là où l'agent aurait
  choisi tout seul.</li>
</ul>

:::

:::col{class="col-span-5 text-left" reveal}


  <h6><strong>Tradeoff</strong></h6>
  <ul class="text-[1.6rem]">
    <li>Un CLI Python à installer, et beaucoup de texte à lire.</li>
    <li>Aucune étape de revue de code dans le workflow.</li>
    <li>Changer de direction, c'est relancer la commande : elle régénère tout
    son document.</li>
  </ul>


:::

:::col{class="col-span-12 text-center" reveal}

#### Projet neuf, gates visibles, et on accepte la *verbosité*. C'est l'outil le plus distribué de la catégorie.

:::
:::


---
layout: content
title: GitHub Spec Kit
github: github/spec-kit
sources:
  - github.com/Fission-AI/OpenSpec
speakernotes: |

---

::browser{src=https://speckit.org/}

---
layout: default
title: "OpenSpec"
github: Fission-AI/OpenSpec/
sources:
  - github.com/Fission-AI/OpenSpec
  - reenbit.com/bmad-vs-spec-kit-vs-openspec-choosing-your-spec-driven-ai-framework
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

:::grid{class="grid grid-cols-12 gap-4 h-stage-large"}

:::col{class="col-span-7 text-left" reveal}

<h6><strong>Principe</strong></h6>

<ul class="text-[1.6rem]">
  <li><strong>Des delta specs</strong> : on décrit ce qui change, pas le
  système entier. Chaque document reste court et relisable.</li>
  <li><strong>Cinq étapes, trois commandes</strong> : explore, propose, apply,
  sync, archive.</li>
  <li><strong>N'importe quel artefact se corrige à tout moment</strong>, sans
  phase gate à repasser.</li>
</ul>

::: 

:::col{class="col-span-5 text-left" reveal}


  <h6><strong>Tradeoff</strong></h6>

  <ul class="text-[1.6rem]">
    <li>Aucun gate de revue entre les phases : c'est à l'humain de contrôler.</li>
    <li>Un support communautaire encore limité.</li>
  </ul>


:::

:::col{class="col-span-12 text-center" reveal}

#### Le même travail que Spec Kit, avec *moins de cérémonie*. Adapté pour du code qui existe déjà.

:::
:::

---
layout: content
title: "OpenSpec"
github: Fission-AI/OpenSpec/
sources:
  - github.com/Fission-AI/OpenSpec
speakernotes: |

---

::browser{src=https://openspec.dev/}

---
layout: default
title: "Superpowers : la méthode comme artefact"
github: obra/superpowers
sources:
  - claude.com/plugins/superpowers
  - github.com/obra/superpowers-marketplace
  - simonwillison.net/2025/Oct/10/superpowers
speakernotes: |
  S21 - Le cas le plus intéressant pédagogiquement. 
  Anecdote : Jesse Vincent, créateur de RT, contributeur de Perl 5, auteur du
  client mail K-9, a quasiment cessé de coder lui-même.
  Chiffres vérifiés le 11/09/2026, utilisables à l'oral : sorti en octobre
  2025, plus de 40 000 étoiles GitHub, et plus d'installations sur le
  marketplace Claude Code que Playwright.
  Sources : claude.com/plugins/superpowers ; obra/superpowers-marketplace ;
  blog de Jesse Vincent sur la genèse du plugin.
---

:::grid{class="grid grid-cols-12 gap-4 h-stage-large"}

:::col{class="col-span-7 text-left" reveal}

<h6><strong>Principe</strong></h6>

<ul class="text-[1.6rem]">
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

:::col{class="col-span-5 text-left" reveal}


  <h6><strong>Tradeoff</strong></h6>

  <ul class="text-[1.6rem]">
    <li>Le plugin traite les agents comme un manager expérimenté traite des
    juniors. L'humain tient le rôle du manager.</li>
    <li>Peu de phases à actionner soi-meme.</li>
  </ul>


:::

:::col{class="col-span-12 text-center" reveal }

#### Ce n'est pas un dépôt de specs, c'est une méthode sur base de skills.


<div class="text-center">
  <a href="https://github.com/obra/superpowers#the-basic-workflow" target="_blank">
    obra/superpowers
  </a>
</div>
:::
:::

---
layout: default
title: "BMAD-METHOD : le poids lourd"
github: bmad-code-org/bmad-method
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

:::grid{class="grid grid-cols-12 gap-4 h-stage-large"}

:::col{class="col-span-7 text-left" reveal}

<h6><strong>Principe</strong></h6>

<ul class="text-[1.6rem]">
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

:::col{class="col-span-5 text-left" reveal}


  <h6><strong>Tradeoff</strong></h6>

  <ul class="text-[1.6rem]">
    <li>Le rituel complet, même quand la tâche ne le justifie pas.</li>
    <li>Du temps, et des tokens.</li>
  </ul>


:::

:::col{class="col-span-12 text-center" reveal}

#### La cérémonie est poussée au *maximum*.

:::
:::


---
layout: content
title: "BMAD"
github: bmad-code-org/bmad-method
sources:
  
speakernotes: |

---

::browser{src=https://docs.bmad-method.org/fr/}

---
layout: default
title: Pour résumer
speakernotes: |
  S24 - Slide de sortie de partie. Quatre primitives universelles :
  règles/constitution, spec, plan, tasks, plus des gates humains.
  Malgré des approches différentes, tous ces frameworks s'accordent sur un
  point : l'humain reste dans la boucle, mais pas pour tout.
  Conclusion libératrice, c'est le message à laisser : tu peux commencer demain
  avec trois fichiers markdown et un gate. Le framework est une optimisation,
  pas un prérequis.
---

:::col{class="col-span-12 text-center" reveal}

### <em>règles &rarr; spec &rarr; plan &rarr; tasks</em>

##### Les quatre outils écrivent les mêmes types d'artefacts : ils ne diffèrent que par la lourdeur de la cérémonie.
:::

:::col{class="col-span-12 text-center" reveal}

#### L'humain reste dans la boucle. Mais pas sur tout.

:::

:::col{class="col-span-12 text-center" reveal}

### Vous pouvez commencer demain avec  *trois fichiers markdown*.

##### Le framework est une optimisation. Pas un prérequis.

:::

---
layout: bare
class: slide-main slide-hero flex flex-col items-center justify-center text-center
speakernotes: |
  Transition partie 4 - Découpler et accélérer. 7 minutes, S25 à S29.
  C'est la partie qui répond à la cinquième question du brief : en quoi le SDD
  accélère réellement.
  Ce qu'on promet : cinq découplages, celui qui compte vraiment, les chiffres
  avec leurs limites, là où ça n'accélère pas, et quoi mesurer.
---

<h1 class="mb-0 text-6xl leading-tight font-bold tracking-tight">Découpler, et <em>accélérer</em></h1>

<p class="text-fg-muted mt-6 text-3xl font-medium">Le SDD est il un game-changer ?</p>

---
layout: default
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

<table class="w-full border-collapse text-[1.6rem] [&_thead_th]:border-b [&_thead_th]:border-outline [&_tbody_tr:not(:last-child)_td]:border-b [&_tbody_tr:not(:last-child)_td]:border-outline">
  <thead>
    <tr>
      <th>Ce que le SDD sépare</th>
      <th>Ce que ça permet</th>
    </tr>
  </thead>
  <tbody>
    <tr class="reveal">
      <td><strong>1. Le quoi et le comment</strong></td>
      <td>La spec devient l'interface entre le jugement humain et l'exécution machine.</td>
    </tr>
    <tr class="reveal">
      <td><strong>2. La conception et l'exécution</strong></td>
      <td>On spécifie maintenant, les agents exécutent pendant que l'on fait autre chose. Le temps de clavier n'est plus le facteur limitant.</td>
    </tr>
    <tr class="reveal">
      <td><strong>3. La revue et le code</strong></td>
      <td>On revoit <em>l'intention</em>, 200 lignes lisibles, au lieu d'un diff de 2 000 lignes qu'on n'a pas écrites.</td>
    </tr>
    <tr class="reveal">
      <td><strong>4. Les développeurs entre eux</strong></td>
      <td>Des specs comme contrats, des bounded contexts, et des agents en parallèle sur des tranches qui ne se percutent pas.</td>
    </tr>
    <tr class="reveal">
      <td><strong>5. L'humain et la session</strong></td>
      <td>Spec et plan tiennent lieu de mémoire durable : reprise après un crash, passage d'un agent à l'autre, onboarding.</td>
    </tr>
  </tbody>
</table>

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

:::grid{class="grid grid-cols-12 gap-4 h-stage-large text-center"}

:::col{class="col-span-12"}

<h3>DORA 2026 lui a donné un nom : la <em>taxe de vérification</em>.</h3>

<p class="text-[1.6rem]">L'effort économisé sur la frappe ne disparaît pas. Il se
relocalise en aval, sur la revue.</p>

:::

:::col{class="col-span-4 text-left" reveal}

<h5><em>La fréquence de déploiement n'augmente pas</em></h6>

<p class="text-[1.6rem]">Le goulot s'est déplacé de l'écriture vers la revue. On
produit plus, on ne livre pas plus souvent.</p>

:::

:::col{class="col-span-4 text-left" reveal}

<h5><em>Le lead time ne baisse pas</em></h6>

<p class="text-[1.6rem]">Le code arrive plus vite en PR. Et il y attend plus
longtemps.</p>

:::

:::col{class="col-span-4 text-left" reveal}

<h5><em>Le temps de revue, lui, augmente</em></h6>

<p class="text-[1.6rem]">Souvent même quand le lead time total diminue : les
relecteurs encaissent un volume croissant de code qu'ils n'ont pas écrit.</p>

:::

:::

---
layout: default
title: "Quelques chiffres d'accélération"
sources:
  - github.com/ianhxu/agentic-engineering-field-study
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

:::grid{class="grid grid-cols-12 gap-4 text-center"}


:::col{class="col-span-4" reveal}

<h3><em>Mythe ou réalité ?</em></h3>

<p class="text-[1.6rem]">Les équipes qui travaillent avec Spec Kit recommencent
bien moins souvent à zéro qu'avec du prompting ad hoc.</p>

<h6><em>GitHub, retours d'usage de Spec Kit</em></h6>

:::

:::col{class="col-span-4" reveal}

<h3><em>40 h &rarr; 8 h</em></h3>

<p class="text-[1.6rem]">de temps humain<br/>sur la même fonctionnalité</p>

<p class="text-[1.6rem]">Des fonctionnalités estimées à 40 heures, livrées en
moins de 8, quand elles sont d'abord rédigées comme specs.</p>

<h6><em>AWS, cas clients Kiro</em></h6>

:::


:::col{class="col-span-4" reveal}

<h3><em>40% (1,7x)</em></h3>

<p class="text-[1.6rem]">de jours gagnés (productivité) sur un développement AI Driven.</p>

<p class="text-[1.6rem]">Nous avons sorti un projet estimé à 660 jours en 400.</p>

<h6><em>Caisse Libre Service, Boulangeer</em></h6>

:::
:::

---
layout: default
title: Quoi mesurer ?
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

:::grid{class="grid grid-cols-12 gap-4"}

:::col{class="col-span-12 text-left"}

<h3>Ce qu'on doit regarder, à la place des lignes de code</h6>

<ul class="text-[2rem]">
  <li class="reveal"><strong>Le lead time jusqu'à la prod</strong>, pas jusqu'à la PR.</li>
  <li class="reveal"><strong>Le temps de revue</strong>.</li>
  <li class="reveal"><strong>Le change failure rate</strong>, pour voir si la vitesse de delivery dégrade la qualité.</li>
  <li class="reveal"><strong>Le taux de rework</strong>, et surtout la part due à une
  <em>intention mal comprise</em>. C'est la seule métrique que le SDD prétend
  améliorer directement.</li>
</ul>

:::

:::col{class="col-span-12 text-center" reveal}

#### Le rework est une taxe cachée sur la productivité. Générer plus de code peut très vite l'*aggraver*.

:::

:::col{class="col-span-12 text-center" reveal}

#### Seules <em>7,3 %</em> des équipes (élites) ont un taux de rework sous 2 %.

:::
:::

---
layout: bare
class: slide-main slide-hero flex flex-col items-center justify-center text-center
speakernotes: |
  Transition partie 5 - Retour au product engineer. 3 minutes, S30 à S33.
  Volontairement la plus courte du deck : deux lignes, et on enchaîne. Elle
  marque la fin de la partie technique, elle ne doit pas voler l'effet de S30.
  Ne rien ajouter à l'oral ici. Afficher, respirer, passer.
---

<h1 class="mb-0 text-6xl leading-tight font-bold tracking-tight">Revenons au <em>product engineer</em></h1>

---
layout: default
title: "L'artefact du product engineer"
class: slide-main slide-narrow slide-prose flex flex-col justify-center
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

<div class="h-12"></div>

### Le product engineer avait besoin d'un <em>artefact</em> pour porter l'intention.

#### On vient de passer plusieurs minutes à décrire cet artefact, et son outillage.

<div class="h-12"></div>

<blockquote class="quote charniere">
  <h4>Sans SDD, un product engineer n'est qu'un dev surchargé à qui on a retiré son PM.</h4>
  <h5>Il a besoin d'un artefact pour porter l'intention jusqu'à la machine. Cet artefact, c'est la <em>spec</em>.</h5>
</blockquote>

---
layout: default
title: "Un bon Product Engineer ?"
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

:::grid{class="grid grid-cols-12 gap-2 h-stage-large text-center"}

:::col{class="col-span-12"}

<h3>Un bon product engineer <em>doit être polyvalent</em>.</h3>


<div class="h-12"></div>

<div class="grid grid-cols-12 gap-4 text-center">
  <div class="col-span-2 reveal"><h6>Cadrer un besoin</h6></div>
  <div class="col-span-3 reveal"><h6>Modéliser un domaine</h6></div>
  <div class="col-span-3 reveal"><h6>Écrire des critères vérifiables</h6></div>
  <div class="col-span-2 reveal"><h6>Arbitrer</h6></div>
  <div class="col-span-2 reveal"><h6>Concevoir la vérification</h6></div>
  <div class="col-span-3 reveal"><h6>Orienter la stack technique</h6></div>
  <div class="col-span-3 reveal"><h6>Dompter l'écosystème agentique</h6></div>
  <div class="col-span-3 reveal"><h6>Maitriser l'ingénierie agentique</h6></div>
  <div class="col-span-3 reveal"><h6>Déployer</h6></div>
</div>

:::

:::col{class="col-span-12 text-left" reveal}

##### Profondeur technique, intuition produit, communication, conscience business, forte autogestion. 

:::

:::col{class="col-span-12 text-left" reveal}

##### <em>Ça fait beaucoup pour une seule personne</em>, et ce n'est pas une question de bonne volonté.</p>

:::

:::

---
layout: content
title: Vous voulez vous lancer ?
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

<table class="w-full border-collapse text-[1.6rem] [&_thead_th]:border-b [&_thead_th]:border-outline [&_tbody_tr:not(:last-child)_td]:border-b [&_tbody_tr:not(:last-child)_td]:border-outline">
  <thead>
    <tr>
      <th>Ce que vous pouvez faire, dès lundi</th>
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
      <td>Une demi-journée, sur une feature adaptée.</td>
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


---
layout: closing
speakernotes: |
  Slide de fin, celle qui reste affichée pendant les questions.
  Rappeler que le brief, les sources et le deck sont publics, et que l'atelier
  de l'après-midi reprend là où le talk s'arrête.
---

# zatsit !

<div class="h-8"></div>

### Restons en contact sur

<div class="h-4"></div>

### [www.zatsit.fr](https://zatsit.fr)
### [blog.zatsit.fr](https://blog.zatsit.fr)

<div class="h-8"></div>

<div class="contact-icons">
  <a href="https://zatsit.fr" aria-label="Le site zatsit">
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="M12 21a9.004 9.004 0 008.716-6.747M12 21a9.004 9.004 0 01-8.716-6.747M12 21c2.485 0 4.5-4.03 4.5-9S14.485 3 12 3m0 18c-2.485 0-4.5-4.03-4.5-9S9.515 3 12 3m0 0a8.997 8.997 0 017.843 4.582M12 3a8.997 8.997 0 00-7.843 4.582m15.686 0A11.953 11.953 0 0112 10.5c-2.998 0-5.74-1.1-7.843-2.918m15.686 0A8.959 8.959 0 0121 12c0 .778-.099 1.533-.284 2.253m0 0A17.919 17.919 0 0112 16.5c-3.162 0-6.133-.815-8.716-2.247m0 0A9.015 9.015 0 013 12c0-1.605.42-3.113 1.157-4.418"/></svg>
  </a>
  <a href="https://blog.zatsit.fr" aria-label="Le blog technique">
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="m18 16 4-4-4-4"/><path d="m6 8-4 4 4 4"/><path d="m14.5 4-5 16"/></svg>
  </a>
  <a href="https://www.linkedin.com/company/zatsit/" aria-label="zatsit sur LinkedIn">
    <svg viewBox="0 0 24 24" fill="currentColor" aria-hidden="true"><path d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"/></svg>
  </a>
  <a href="https://github.com/zatsit-oss" aria-label="zatsit sur GitHub">
    <svg viewBox="0 0 24 24" fill="currentColor" aria-hidden="true"><path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/></svg>
  </a>
  <a href="mailto:contact@zatsit.fr" aria-label="Écrire à contact@zatsit.fr">
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="m22 7-8.991 5.727a2 2 0 0 1-2.009 0L2 7"/><rect x="2" y="4" width="20" height="16" rx="2"/></svg>
  </a>
</div>
