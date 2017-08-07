---
title: "Bienvenue sur le nouveau site du Comptoir Sécu!"
subheadline: "Au menu: JAMStack avec Hugo, Netlify, Foundation, etc."
publisher: morgan
date: 2017-07-31
tags:
  - JAMStack
  - Hugo
  - Netlify
  - Foundation
  - ISSO
  - Podigee
  - Gulp
---
Nous vous l'avons "teasé" depuis des semaines, pour ne pas dire des mois, enfin, le voilà : la nouvelle version de notre site internet!

Cela faisait quelque temps que nous souhaitions nous séparer de cette passoire qu'est WordPress (et son cortège de plugins plus ou moins maintenus), les gros problèmes de performances dus à des plugins trop gourmands sur un serveur trop fébrile nous ont enfin poussés à nous lancer.

Ce nouveau site est entièrement statique...bon..ok...sauf les commentaires.

Pour les curieux sur les détails techniques, voici un aperçu :

Pour le moment nous n'avons pas trouvé de meilleure alternative au système de commentaires wordpress que le moteur open source [ISSO](https://posativ.org/isso/). Il tourne sur une instance ec2 "Micro" ce qui devrait nous permettre de rester sur le plan gratuit d'AWS. Nous ne voulions pas de Disqus, qui dans son plan gratuit impose l'affichage de publicité, en plus d'être bourré de trackers, ce qui alourdit considérablement les pages web qui l'utilisent. J'ai bien vu 2-3 autres moteurs intrigants, comme celui se reposant sur [AWS Lambda](https://github.com/jimpick/lambda-comments), mais cela me semble pour le moment un peu trop artisanal. Les systèmes "statiques" utilisant les [GitHub comments](http://donw.io/post/github-comments/) ou les pull requests semblent bien trop contraignants en l'état.

Nous avons également changé de lecteur audio pour passer à celui de [Podigee](https://www.podigee.com/en/podcast-player/), qui a le bon goût de proposer automatiquement un lien de téléchargement direct, ainsi que de supporter le chapitrage.

Au niveau de la génération du site, nous utilisons [Hugo](https://gohugo.io/), qui ne fait que monter en popularité et finira, je pense, par décrocher le monopole de [Jekyll](https://jekyllrb.com/) à un moment ou un autre. Cela ne serait-ce que pour la rapidité de génération des pages et ses capacités de génération avec les ["custom formats"](https://gohugo.io/templates/output-formats/).

Au niveau de l'apparence visuelle, c'est grosso modo du fait maison en s'appuyant sur le framework CSS [Foundation](http://foundation.zurb.com/sites/docs/).

Pour l'optimisation du CSS, javascript et HTML, la génération et compression des vignettes, et j'en passe, nous nous servons de [Gulp](https://gulpjs.com/).

Pour ce qui est de l'hébergement, c'est une combinaison entre [GitHub](https://github.com/comptoirsecu/csec-hugo) qui héberge notre code source et [Netlify](https://www.netlify.com/) qui, en réaction à chaque commit, le récupère, le compile et le déploie.

Seule ombre au tableau, nous utilisons désormais [Google Analytics](https://www.google.com/analytics/) pour avoir un aperçu sur les statistiques du site. Nous passerons si possible sur un système self-hosted à l'avenir, du genre de [Piwik](https://piwik.org/). Pour le moment, nous nous contenterons de donner notre âme à Google.

 Avec tout ça, on est passé d’un temps de chargement moyen des pages de plusieurs secondes à environ 800ms. On espère que vous apprécierez ce gain en réactivité.

Nous avons profité de ce changement d'architecture pour revoir l'organisation et l'allure du site.

Vous retrouverez désormais sur la page d'accueil une mise en avant de nos contenus récents au format podcast, live ou vidéo.

![Contenus récents](/images/misc/newsite_1.jpg)

La dernière annonce importante, comme ce billet, est ensuite affichée.

![Annonces](/images/misc/newsite_2.jpg)

Vous retrouvez enfin nos derniers contenus par catégorie:

* Les derniers épisodes du comptoir, incluant les hors-séries comme nos couvertures des conférences sécu;
* Les SECHebdo, revue d'actualité hebdomadaire (ou presque);
* Les vidéos, que ce soit les SECompris, SEClairs, PoC, ou peut-être des Points Sécu et autres nouveaux formats.

Vous pouvez accéder aux archives de chacun de ces contenus assez facilement via le menu en haut du site, qui inclue également un accès à nos (vieux) articles, mais aussi à une toute nouvelle rubrique: le registre!

Eh oui, vous en rêviez (ou pas), on l’a fait! Désormais, chaque nouveau billet annonçant un SECHebdo, un épisode du Comptoir ou une vidéo inclura dans ses métadonnées

* les [membres du Comptoir Sécu](/authors) qui y ont participé;

![authors](/images/misc/new_site3.jpg)

* les [invités](/guests);

![invités](/images/misc/new_site4.jpg)


* les [boissons](/drinks) consommées et

![drinks](/images/misc/new_site5.jpg)


* les [musiques](/songs) diffusées en interludes!

![musiques](/images/misc/new_site6.jpg)


Cela nous a pris un temps considérable, mais nous avons pris le soin de mettre à jour tous nos anciens contenus pour qu'ils affichent ces métadonnées.

Nous avons également profité du changement de notre lecteur audio pour implémenter le chapitrage de nos épisodes lorsque celui-ci était disponible dans l'article.

![chapitres](/images/misc/new_site7.jpg)

Nous avons d'autres surprises sous le coude, certaines déjà mises en place comme l'affichage du site lors d'un live. D'autres encore en préparation comme le support du format [AMP](https://www.ampproject.org/learn/overview/) pour charger les pages encore plus vite sur mobile.

Un autre avantage indéniable de cette nouvelle architecture, c'est la capacité que vous avez à contribuer! En effet, le code source étant maintenant disponible publiquement sur [GitHub](https://github.com/comptoirsecu/csec-hugo), n'importe qui peut nous proposer des [pull requests](https://help.github.com/articles/about-pull-requests/).

Vous vous sentez de réaliser le chapitrage des anciens épisodes? Nous avons fait une coquille dans un article ou les métadonnées? Vous avez trouvé un moyen facile d'implémenter la pagination dans les registres? Faites nous une [pull request](http://blog.zenika.com/2017/01/24/pull-request-demystifie/) et participer à l'amélioration du comptoir!
