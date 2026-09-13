# SDD : Spec-Driven Development

Talk de 45 minutes, en français, devant un club de développeurs, en ouverture
d'une demi-journée dont la suite est un atelier. Le deck est `demoit.md`, la
syntaxe Markdown du moteur est décrite dans le `CLAUDE.md` à la racine du
dépôt, section « Writing slides » pour les conventions de rédaction.

Le brief qui a servi à construire le plan est `sdd-product-engineer-brief.md`,
à côté. Il reste utile comme source, mais **le deck fait foi** : plusieurs de
ses chiffres se sont révélés faux à la vérification et ont été corrigés dans
les speaker notes des slides concernées.

## Lancer le talk

Build du moteur depuis la racine du dépôt, puis :

```shell
demoit sdd-talk
```

Avec le live reload pendant l'écriture des slides (la page se recharge à chaque
sauvegarde de `demoit.md`) :

```shell
demoit --dev sdd-talk
```

Les slides sont servies sur http://localhost:8888. La fenêtre de notes est sur
http://localhost:8888/speakernotes, la vue d'ensemble sur `/grid`.

## Prérequis de démo

Aucun. Le brief a explicitement écarté la démo live : avec ce volume de
concepts, une démo qui plante coûte huit minutes et la crédibilité. La
démonstration est renvoyée à l'atelier de l'après-midi.

Une seule dépendance réseau à l'écran : la slide S13 rend son diagramme avec
mermaid, importé depuis un CDN par `.demoit/js/demoit.js`. Sans réseau, ce
diagramme reste invisible. À vérifier avant de monter sur scène.

## Structure

42 slides, six parties, chacune ouverte par une slide de transition
(`layout: bare`, `class: transition`) :

| Partie | Durée | Slides |
|---|---|---|
| 0. Un nouveau métier, et son problème | 7 min | S1 à S5 |
| 1. Avant le SDD : les prérequis | 8 min | S6 à S10 |
| 2. Le SDD et son intérêt | 10 min | S11 à S17 |
| 3. Le panorama des frameworks | 10 min | S18 à S24 |
| 4. Découpler et accélérer | 7 min | S25 à S29 |
| 5. Retour au product engineer | 3 min | S30 à S33 |

S5 et S30 sont les deux charnières : même mise en forme, même phrase, l'une
pose le lien product engineer / SDD, l'autre le résout. Elles se modifient
ensemble ou pas du tout.

## État

Contenu écrit de bout en bout, toutes les slides rendent sans erreur, les
sources sont en speaker notes. Restent à la main du speaker :

- **S14** : la comparaison spec floue / spec complète sur son propre code,
  avec le résultat produit par l'agent dans les deux cas. Le brief la désigne
  comme le slide dont la salle se souviendra.
- **S32** : un `AGENTS.md` ou `constitution.md` d'exemple, lisible en vingt
  secondes à l'écran, s'il veut en montrer un.
- **S33** : le contenu réel de l'atelier, et les trois questions laissées
  ouvertes (celles qui y figurent sont déduites du deck, à confirmer).
- Une relecture au vidéoprojecteur : S9 et S13 sont les deux slides les plus
  denses, et le rendu mermaid de S13 n'a pas pu être vérifié autrement qu'à
  l'œil.
