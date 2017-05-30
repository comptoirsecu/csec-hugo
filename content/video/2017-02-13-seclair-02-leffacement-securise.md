---
id: 1380
title: "[SEClair #02] L'effacement sécurisé"
subheadline: "Parce qu'il y en a marre de voir conseiller des effacements à 7 passes"
author: morgan
date: 2017-02-13

aliases: /2017/02/seclair-02-leffacement-securise/
podcast:
  feed: http://media.blubrry.com/comptoirsecu/p/podcast.comptoirsecu.fr/CSEC.HS35.2017-02-13.SEClair_Effacement_Securise.mp3
image:  /images/2017/02/SEClair-02-Effacement-securise-vignette-lowres-v2.jpg
categories:
  - SEClair
tags:
  - disque dur
  - effacement sécurisé
  - gutmann
  - photorec
  - podcast
  - recuva
  - SEClair
  - secure erase
  - ssd
video: https://www.youtube.com/embed/8CCYQpvmu3s
---
Bonjour à tous,

SEClair est de retour! Cette fois-ci sur l'effacement sécurisé.

Alors à votre avis combien de passe de formatage sur votre disque pour qu'il ne soit plus possible de récupérer des données dessus? 1? 3? 5? 7? 21? Vous le saurez en écoutant cet épisode ;).
<!--more-->

Pour les allergiques au format vidéo, la bande son est disponible en podcast:

{{< podigee >}}

A bientôt pour d'autres émissions/podcasts!

## **Script:**

&nbsp;

### Introduction

Aujourd'hui on va parler de l'effacement sécurisé des données.

En d'autre terme, comment s'assurer qu'on ne puisse pas récupérer mes données après effacement si jamais je vends, donne ou jette mon disque dur !

Oui parce que j'en aie marre des articles qui conseillent de faire 3, 7, 21 voire 35 passes de formatage ! Donc messieurs les paranos, rangez vos casques en papier alu et écoutez attentivement la suite :).

&nbsp;

Revendre des disques dur tout seul ça arrive pas souvent mais un PC ou un téléphone d'occasion c'est déjà plus fréquent. Les particuliers ne veulent pas que l'on retrouve leurs mots de passe ou leurs photos coquines. Les entreprises ne veulent pas qu'on puisse récupérer des données confidentielles métiers et autres secrets industriels en fouillant leurs poubelles.

### Comment marche un disque dur

Déjà il faut rappeler une chose importante : La plupart du temps vous n'effacez pas vraiment vos données quand vous pensez l'avoir fait.

En effet, Supprimer un fichier via la fonction &laquo;&nbsp;supprimer&nbsp;&raquo; de Windows/Mac ou autre OS ne supprime pas réellement le fichier. Effacer tout un disque via le &laquo;&nbsp;formatage rapide&nbsp;&raquo; n'efface pas réellement le disque non plus.

Si vous ne me croyez pas essayez la chose suivante :

  * Vous prenez une clé USB, si possible neuve
  * Vous mettez quelques photos dessus
  * Vous en effacez 2-3 avec suppr
  * Voir même vous faites un formatage rapide sur la clé
  * Vous utilisez un outil comme Photorec et vous constatez.

Sorcellerie ? Mais non, pour faire court, un disque peut être vue comme un bouquin ou chaque chapitre serait un fichier.

Si je vous demande de trouver un chapitre précis vous pouvez vous amusez à parcourir les pages une à une jusqu’ à tomber dessus mais ça va être long. Si vous êtes un peu malin vous allez utiliser la table des matières qui va référencer la page ou commence le chapitre. Bah le système de fichier sur votre disque dur c'est pareil, il y a une espèce de grosse table des matières qui référence les fichiers et leur localisation.

Quand vous effacer un fichier, vous effacer juste la référence dans la table des matières. Le chapitre est toujours dans le bouquin, il n'est juste plus référencé. Quand vous faites un formatage &laquo;&nbsp;rapide&nbsp;&raquo;, vous effacer juste toute la table des matières. Les pages qui contenaient le chapitre/fichiers sont maintenant déclarées comme &laquo;&nbsp;vide&nbsp;&raquo;

Et un jour, quand vous allez créer ou télécharger un fichier, et qu’il faudra de la place pour le stocker. Votre disque dur choisira cet emplacement qui a la bonne taille et écrira par-dessus. Avec un bouquin ça serait illisible, faudrait d'abord effacer, mais sur un disque ça marche, ça remplace ;).

Pourquoi on fait ça ? Parce que c'est vachement plus rapide ! Essayez de faire un formatage complet sur un disque de 1To vous allez vite comprendre pourquoi par défaut on se contente d'effacer la table des matières. Et en plus d'être plus rapide, ça évite d'user le disque avec des écritures inutiles.

Bon, maintenant comment a fait PhotoRec pour récupérer vos images ? Bah il a parcouru le bouquin comme un bourrin page à page et ne s'est pas fié à la table des matières. Dès qu'il trouve une suite de lettre…enfin octets, qui ressemble à une image, il essaye de la restaurer.

Après ce n’est pas magique, si vous avez réécrit sur votre disque et si une partie de l’image voir toute l’image a été écrasé par un autre contenu, vous la retrouverez pas l’image ! Si vous voulez essayez, reprenez votre clé USB et cette fois-ci remplissez l’entièrement avec d'autres documents. Ressayez d'utiliser PhotoREC, vous pouvez dire adieu à votre photo.

Vous avez donc compris, pour être sûr que la donnée n'est pas récupérable, il faut l'écraser avec une autre donnée. C'est ce que font la plupart des outils d'effacement sécurisé. On remplit le disque avec des 0 et c'est marre. On peut aussi remplir le disque avec des 1, des motifs de 0 et 1, ou un flux aléatoire de 0 et 1, ce qui semble la méthode préférable.

Mais, si on a tout réécrit une fois, on peut plus rien récupérer, pourquoi il faut faire plein de passes alors ?

### La magie du magnétisme

&#8230;Parce qu'il y a un monsieur qui s'appelle Peter Gutmann et qui, il y a 20 ans, a fait des recherches sur le sujet. A savoir : Est-il possible de savoir quelle donnée il y avait avant la réécriture ?

Le monsieur explique dans son papier qu'il est possible de récupérer le bit qui était stocké avant réécriture. Et cela même après plusieurs réécritures.

Pour rappeler très vite fait comment marche un disque magnétique. Vous avez des sortes de galettes fait dans un matériau très dur (d'où le nom disque dur) recouvert d'une surface magnétisable et découpées en plein de toutes petites cases. Chaque case permet de stocker un bit d'information, donc un 0 ou un 1.

Une case magnétisée représente un 1, une case non magnétisée un 0. Le disque tourne très vite et une toute petite tète de lecture/écriture (ça ressemble à un bras de lecture vinyle en gros) qui parcours le disque. Cette tête peut soit lire l’état magnétique d'une case pour lire si c'est un 0 ou un 1. Soit magnétiser ou démagnétiser une case pour y écrire un 1 ou un 0 respectivement. Sauf que je dis 0 et 1 depuis tout à l’heure, mais en vrai c'est pas binaire on est dans le monde de l'analogique. Donc une case est plus ou moins magnétisée.

Imaginons que l’état soit un nombre entre positif ou négatif. -10 étant idéalement démagnétisée et +10 idéalement magnétisée. Bah la tête lira peut-être -9, +11, -8, +10. La tête de lecture elle va numériser tout ça, par exemple si c'est négatif c'est 0 et si c'est positif c'est 1. Sur ce postulat, -9 +11 -8, +10 va donner respectivement 0, 1, 0, 1 en binaire.

Selon Gutmann, les services de renseignements étatiques disposeraient de Microscopes à Force Magnétique qui permettrait de lire précisément l'état magnétique de chaque petite case du disque dur, et d'en déduire l'état avant réécriture.

Par exemple si l'état est -9 au lieu du -10 attendu. Il en déduit que l'état avant était dans le positif et que donc avant ce 0 il y avait un 1. La méthode peut théoriquement s'appliquer plusieurs fois d'affilé pour remonter plusieurs réécritures.

<pre>Signal numérique :       1       0       1       0       1       0

Signal analogique :       +11.1   -8.9    +9.1    -11.1   +10.9   -9.1

Signal idéal :            +10.0   -10.0   +10.0   -10.0   +10.0   -10.0

Différence :               +1.1    +1.1    -0.9    -1.1    +0.9    +0.9

Signal précédent déduis :  +11     +11     -9      -11     +9      +9

Signal numérique précédent : 1     1       0       0       1       1</pre>

&nbsp;

2ème passe

<pre>Signal numérique :         1       1    0     0     1     1

Signal analogique :          +11    +11   -9   -11    +9    +9

Signal idéal :               +10.0 +10.0 -10.0 -10.0 +10.0 +10.0

Différence :                 +1    +1    +1    -1    -1    -1

Signal précédent déduis :    +10   +10   +10   -10   -10   -10

Signal numérique précédent : 1     1     1     0     0     0</pre>

Alors c'est super, sauf que tout c’est purement théorique et n'a même pas été vérifié en laboratoire.

De plus le papier se base sur une technologique de disque dur qu'on utilise plus depuis 15 ans ! D'ailleurs, un papier sorti en 2008 démonte les conclusions de Peter Gutmann en démontrant que la probabilité de récupérer un bit est seulement d'une chance sur deux.

[Overwriting Hard Drive Data: The Great Wiping Controversy](https://www.vidarholen.net/~vidar/overwriting_hard_drive_data.pdf) (Craig Wright, Dave Kleiman, Shyaam Sundhar R.S.)

&nbsp;

Attention ça veut dire une chance sur deux de récupérer un BIT, par un document. En récupérer plusieurs d'affilé est statistiquement encore plus dure. Genre récupérer un octet entier ça serait plus proche de 3%, alors un document de plusieurs kilooctets voire mégaoctets, soit des milliers ou des millions d'octet&#8230;

Autant dire mission impossible.

En plus maintenant les disques durs récent sont beaucoup, beaucoup, beaucoup plus denses et plus complexe. Je me souviens fièrement de mon premier disque de 3 Go&#8230;maintenant 1To c'est le standard et du 3To n'est pas bien plus cher. Aucune chance de récupérer quelque chose d'exploitable.

&nbsp;

Bref oublier les méthodes Gutmann complètement tarées à 35 passes. Oubliez même les méthodes à 7 passes ou 3 passes. En vrai, une passe suffi, de 0 ou d'aléatoire entre 0 et 1.On ne pourra rien récupérer en désossant votre disque dur pour l'analyser au microscope. Encore moins avec des logiciels comme Recuva ou Photorec.

&nbsp;

Voilà c'est fini merci d'avoir suivi &#8230;

NON

&#8230;quoi non ?

Et les SSD dans tout ça hein ? \*trollface\*

### Cas particulier: Les disques SSD

&nbsp;

…Oui c'est vrai, les disque SSD sont différent. Ils ne sont pas basés sur non pas sur un support magnétique mais sur de la mémoire Flash. C'est un peu des grosses clé USB.

Ce système de stockage se base sur la capture d'un état électrique dans ce qu'on appelle des semi-conducteur, et qui conserve l'état même si on éteint le PC. Bref, ça a plein d'avantage, notamment de rapidité d'accès à la donnée etc., mais la durée de vie des composants qui stock le bit 0 ou 1 est plus faible que sur un disque dur.

&nbsp;

Du coup il y a tout plein de techniques utilisées pour optimiser la durée de vie du disque. L'OS choisi plus ou il met la donnée physiquement, le SSD le fait pour lui en s'arrangeant pour que tout le disque soit utilisé de manière homogène. Il y a du stockage en rab contrer la perte de cellule sans affecter la capacité du disque.

Tout ça pour dire qu'on peut plus tout bêtement réécrire des 0 sur un document.

Parce qu’on n’est pas sûr que le disque va effectivement les écrire sur le document qu'on cherchait à effacer. Par chance, les constructeurs y ont pensé et ont implémenté des fonctionnalité d'effacement sécurisées. Pour les utiliser il faut généralement se servir d'outils fournis par les constructeurs des SSD en question.

Bien entendu c'est franchement déconseillé de faire un effacement sécurisé de tout le SSD tous les 4 matin si vous voulez qu'il tienne un moment&#8230;

Comment faire alors ? Et bien ça dépend du besoin.

  * Vous voulez effacer un disque avant de le revendre ?
      * Une passe de 0 ou d'aléatoire sur un disque dur, ou la fonction d'effacement sécurisé du constructeur pour les SSD. Ça suffira amplement.
  * Vous avez un disque qui change souvent de mains, par exemples des clients différents, et vous ne voulez pas passer 8 ans à chaque fois à faire un formatage complet ?
      * Chiffrez tout le disque. Plus besoin d'écraser les données, si le nouvel utilisateur n'a pas la clé, il ne pourra rien récupérer !
  * Vous comptez jeter des disques et être sûr que quelqu'un ne fouille pas les poubelles pour vous espionner ? Pas même la NSA ?
      * Allez-y plus bourrin et détruisez physiquement le disque. Un petit coup de perceuse et c'est fini !
