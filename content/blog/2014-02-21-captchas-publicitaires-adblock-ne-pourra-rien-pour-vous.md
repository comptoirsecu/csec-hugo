---

title: "CAPTCHAs publicitaires : Adblock ne pourra rien pour vous !"
date: 2014-02-21T09:25:36+00:00
authors:
 - morgan

aliases: /2014/02/captchas-publicitaires-adblock-ne-pourra-rien-pour-vous/
views: 2092
image: /images/covers/2014-02-playcaptcha.jpg
categories:
  - Article
tags:
  - adblock
  - CAPTCHA
  - Mechanical Turk
  - publicité
  - reCaptcha
---
Vous pensiez être sauvé de la publicité ? Les marques ne baissent pas les bras si facilement, voici une méthode qui devrait vous forcer à écouter leur accroche, et même à la répéter avec vos petits doigts ! J’ai nommé les CAPTCHAs publicitaires.

{{<toc>}}

# Un Captchquoi ?

Pour rappel, CAPTCHA signifie : "Completely Automated Public Turing test to tell Computers and Humans Apart". En bref, un mécanisme permettant de s’assurer que le visiteur de la page est bien un humain et non pas une machine.

Les CAPTCHA ont pour objectif d’endiguer le spam des formulaires en ligne, en particulier les commentaires. On le voit bien au comptoir, on a un sacré paquet de bots qui tentent leur chance pour caser leur lien sponsorisé discrètement dans le corps du commentaire ou dans leur pseudonyme.

Et Dieu sait que ces bots nous aiment ! Jugez plutôt :

![commentaires-indesirables](/images/misc/2014-02-commentaires-indesirables.jpg)

Heureusement qu’[Akismet](http://akismet.com/) nous facilite le travail et sait séparer le grain de l’ivraie.

## Comment fonctionne un CAPTCHA, exemple de reCAPTCHA

Quand on parle de CAPTCHA, vous avez sûrement en tête ces suites de deux mots totalement illisibles :

![recaptcha](/images/misc/2014-02-recaptcha.jpg)

Ces gribouillis sont fournis gratuitement par le service [reCAPTCHA](https://www.google.com/recaptcha), détenu par…Google.

Le projet est louable, ils se servent de ce mécanisme pour améliorer la qualité de leur OCR. Un OCR (Optical Character Recognition), c’est une famille d’algorithmes faits pour numériser des documents imprimés en reconnaissant ce qui y est écrit (et si possible, la mise en page).

C’est un travail très difficile, la page n’est jamais scannée de façon parfaitement droite, l’image peut être floue, le papier abîmé, déchiré, la police exotique… Les meilleures méthodes se basent sur des algorithmes réalisant de l’autoapprentissage. On leur envoie des millions et des millions d’exemples. Lorsqu’ils se trompent, ou qu’ils ne sont pas sûrs, un être humain les corrige, ils apprennent de leur erreur et affinent la méthode.

Sauf que demander à un humain de faire ce genre de travaille, ça coûte cher, autant profiter de ces millions d’internautes qui cherchent à prouver leur humanité pour troller sur le dernier article de Korben ou pour valider leur inscription à la bêta test du dernier jeu à la mode !

![recaptcha-how](/images/misc/2014-02-recaptcha-how.jpg)

Et là, les plus filous se diront :

> « Mais, s’il n’arrive pas à déchiffrer le mot, comment peut-il savoir que l’on a mis la bonne réponse ? »

C’est justement pour cela qu’il n’y a pas un, mais deux mots à recopier. Un des deux est connu de l’OCR et servira à vérifier votre humanité. Le second sera utilisé pour améliorer l’algorithme.

> « Mais alors, si je réponds comme il faut au mot de contrôle et complètement à côté de la plaque sur le mot qu’il ne connait pas, j’induis l’algorithme en erreur ? »

Bien évidemment non, le mot est présenté non pas à une, mais à plusieurs personnes, quand les réponses convergent vers un mot du dictionnaire on peut affirmer avec un bon taux de confiance que la réponse est la bonne.

Malheureusement, il est vraiment important que le spam ne passe pas. Il faut donc absolument que la reconnaissance du mot soit insurmontable, du coup reCaptcha rajoute de la difficulté en « défigurant » le mot avec quelques distorsions, qui ont le bon goût de nous rendre dingue lorsque l’on a à les résoudre.

Pas plus tard qu’hier j’ai dû m’y reprendre à 5 fois avant que LinkedIn accepte mon CAPTCHA, quand on tape tout ça sur un smartphone ça a de quoi rendre chèvre :

![recaptcha-secu](/images/misc/2014-02-recaptcha-secu.jpg)

## À la recherche d’alternatives aux hiéroglyphe

Comment faire autrement ? L’idée est toute bête, plutôt que de nous faire reconnaître du texte, chose pour lequel les ordinateurs sont de plus en plus doués, demandons plutôt de réaliser une action demandant de la compréhension sémantique et de la reconnaissance d’image. Là-dessus, les ordinateurs sont beaucoup, beaucoup moins bons !

Cela peut se faire avec une simple phrase, de type :

> « Inscrivez dans le champ ci-contre l’animal de la liste ci-après qui ne vit pas dans l’eau :

> Sardine, baleine, cochon, phoque. »

Ou une animation flash nous demandant de cliquer sur la fourchette parmi une liste d’éléments présents.

![areyouahuman-example](/images/misc/2014-02-areyouahuman-example.jpg)

Cependant, quelques entrepreneurs y ont vu un autre filon : pourquoi ne pas forcer l’utilisateur à regarder attentivement une publicité et lui poser une question sur celle-ci ? Mieux encore ! Pourquoi ne pas lui faire répéter gentiment notre slogan ?

C’est là qu’arrivent ces nouveaux CAPTCHA publicitaires, avec des entreprises comme [AreYouHuman](http://areyouahuman.com/demo-playthru/), [Solvemedia](http://www.solvemedia.com/advertisers/captcha-type-in) ou [FutureadLabs](http://www.futureadlabs.com/wp-content/uploads/2014/02/Publisher-PlayCaptcha-Infographic.png)

![solvemedia-example](/images/misc/2014-02-solvemedia-example.jpg)

# CAPTCHA et sécurité

Parlons maintenant un peu de l’aspect sécurité des CAPTCHA. J'y vois quatre problèmes potentiels :

## La mise en forme de ce type de publicités

Si toutes les publicités de solvemedia sont du type « Un bandeau publicitaire avec en dessous ‘Please enter : XXX’ », je me dis que cela peut facilement se reconnaître. Ils expliquent disposer de moyens de détection des bots et d’adapter [le CAPTCHA en conséquence](http://www.solvemedia.com/security/index.html). Après tout si le système fonctionne, pourquoi pas ?

> « CAPTCHA is more than just squiggly letters and pictures. We use mathematical algorithms to calculate the probability of a user being a human or spammer before we generate a puzzle. We then adjust the difficulty of the puzzle we serve accordingly. The algorithms we use are well established, thoroughly researched, and peer reviewed. »

![solvemedia-secu](/images/misc/2014-02-solvemedia-secu.jpg)

## L'accessibilité pour tous...ou pas

Les mécanismes d’accessibilité : Quand une alternative est accessible et plus facile, les attaquants n’y réfléchissent pas à deux fois. Le cas typique est le mode "accessibilité" avec reconnaissance vocale pour les aveugles. Il faut être sûr que ces contrôles « accessibles » ne le soient pas pour les bots.

![captcha-audio](/images/misc/2014-02-captcha-audio.jpg)

CAPTCHA et l'art de la délocalisation

Délocaliser la complétion des CAPTCHA : ne riez pas, c’est déjà en place. Il est possible à l’heure actuelle de payer [des heures de résolution de CAPTCHA](http://www.troyhunt.com/2012/01/breaking-captcha-with-automated-humans.html) sur internet. [Amazon s'y met](https://www.mturk.com/mturk/) à sa façon !

![mechanical turk](/images/misc/2014-02-mechanical-turk.jpg)

Ici, à part faire du contrôle sur le volume de CAPTCHA résolu par IP/Zone géographique et mettre en place du ban ou des temps d’attente obligatoires, je ne vois pas de solutions miracles.

## Problématique d'entropie

Si le captcha n’est pas aléatoire ou qu’il n’y a pas un assez grand pool de combinaison, il est possible d’enregistrer toutes les combinaisons possibles et de faire de la reconnaissance sur une caractéristique différenciant.

C’est un peu comme ça que le système de sécurité de [SnapChat s’est fait retourner en seulement quelques heures](http://gadgets.ndtv.com/apps/news/snapchats-brand-new-find-the-ghost-security-feature-gets-hacked-474690) par un chercheur.

Sur ce dernier point, les plus protégés sont sûrement les CAPTCHA publicitaires demandant un peu d’interactivités, surtout à la souris :

![playcaptcha-example](/images/misc/2014-02-playcaptcha-example.jpg)

# Conclusion : Que penser de la nouvelle approche publicitaire ?

Dans l’ensemble, je suis assez séduit par l’idée, aussi machiavélique soit-elle, et je suis étonné que cela ne soit pas plus répandu. Je suis d’autant plus étonné étant donné que reCAPTCHA est possédé par Google !

La pub, ils connaissent, et ils savent aussi à quel point il est difficile de faire en sorte que l’internaute les regarde pour que les acheteurs d’audience soient satisfaits de l’investissement.

Il faut reconnaître que cela rend la résolution de CAPTCHA plus agréable, voir ludique pour areyouahuman par exemple. Tout ce qui rend la sécurité plus facile à supporter pour l’utilisateur est une bonne chose. Et une chose est sûre, les gens détestent les CAPTCHA classiques.

Et vous, que pensez-vous de ces CAPTCHA publicitaires ?
