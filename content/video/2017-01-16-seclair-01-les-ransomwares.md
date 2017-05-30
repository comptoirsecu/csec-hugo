---
id: 1338
title: '[SEClair #01] Les Ransomwares'
subheadline: "Vous allez regretter d'avoir oublié de backup!"
author: morgan
date: 2017-01-16

aliases: /2017/01/seclair-01-les-ransomwares/
podcast:
  feed: http://media.blubrry.com/comptoirsecu/p/podcast.comptoirsecu.fr/CSEC.HS30.2016-01-16.SEClair_Les_Ransomwares.mp3
  descripton: "SEClair est une variante un peu plus courte que SECompris (et avec beaucoup, beaucoup plus de cuts).

    Je vous cache pas que c'est un essai, n'hésitez pas à me dire dans les commentaires ce que vous en pensez. Je me dis que ça va peut-être un peu trop vite niveau information à assimiler.

    A bientôt pour d'autres émissions/podcasts!"
image:  /images/2017/01/SEClair-01-Les-ransomwares.png
categories:
  - SEClair
tags:
  - podcast
  - rançongiciel
  - ransomware
  - SEClair
  - virus
  - Youtube
video: https://www.youtube.com/embed/3tA1x3LhXy0
---
Bonjour à tous!

SEClair est une variante un peu plus courte que SECompris (et avec  beaucoup plus de cuts).

Je vous cache pas que c'est un essai, n'hésitez pas à me dire dans les commentaires ce que vous en pensez. Je me dis que ça va peut-être un peu trop vite niveau information à assimiler.

<!--more-->

L'objectif pour le moment est d'avoir des SEClair en complément, et non en remplacement, des SECompris. En effet certains sujets ne peuvent clairement pas se boucler en moins de 10 minutes. Désolé pour les personnes ayant du mal avec les formats long mais je ne pourrais pas toujours y couper .

Une fois n'est pas coutûme, j'avais défini un script et je m'y suis majoritairement tenu, donc cadeau pour ceux que ça intéresse, je vous le met en bas de ce billet!

Comme pour SECompris, je mets la bande sonore en podcast pour les allergiques de YouTube. Ce n'est pas une copie parfaite, j'ai extrait le son avant d'injecter les extraits vidéos, ça devrait être moins perturbant.

{{< podigee >}}

A bientôt pour d'autres émissions/podcasts!

## **Script:**

&nbsp;

### Introduction

<div>
  Bonjour et bienvenu sur SEClair, la variante de SECompris pour les personnes ayant un trouble de l'attention!
</div>

<div>
</div>

<div>
  Non je déconne&#8230;quoi que&#8230; cette émission aura pour objectif de traiter des sujets de façon plus courtes, du genre en 5 minutes au lieu de 20.
</div>

<div>
</div>

<div>
  Forcément pour faire si court soit le sujet sera plus petit soit il sera traité de façon plus superficielle.
</div>

<div>
</div>

<div>
  Bon trêve de bavardage, j'avais dis rapide, on va inaugurer l&#8217;emission avec un sujet à la mode: Les ransomwares.
</div>

<div>
</div>

<div>
</div>

### Ransomware: Kezako?

<div>
  <p>
    Un ransomware, ou rançongiciel en français, est un type de virus qui va vous extorquer vos précieux deniers en tenant en otage vos données.
  </p>

  <p>
    Pour ça, il va généralement chiffrer vos documents importants et vous demander de payer pour obtenir la clé de déchiffrement.
  </p>

  <p>
    Et ils ne sont pas frileux sur les sommes, c'est souvent entre 200 et 500 euro pour obtenir le précieux sésame. Alors la méthode a rien de nouvelle hein, les rançongiciels ça existe depuis plus de 20 ans Et pourtant c'est revenu très fortement à la mode depuis seulement 2-3 ans.
  </p>

  <p>
    Pourquoi ? Parce qu'il est maintenant beaucoup plus facile de vous faire payer sans se faire pincer ou se faire bloquer les fonds en cours de route ! Pour ça ils n'essayent plus de se faire payer par carte mais vous demande de régler en cryptomonnaie, généralement Bitcoin, Il vous fournit la clé via un réseau anonyme de type darknet, genre TOR.
  </p>

  <p>
    Si je vous parle en chinois, je vous laisse regardez les épisodes du Point Sécu fait par Justin sur les Bitcoin et les Darknets, parce que là&#8230;pas le temps :).
  </p>

  <h3>
    Est ce que ça marche vraiment ?
  </h3>

  <p>
    &nbsp;
  </p>

  <p>
    Oh que oui ! En 2015 on a estimé que les auteurs de Cryptowall ont récolté près de 300 millions de dollars avec cette méthode.
  </p>

  <p>
    Et ça ce n'est qu'UN ransomware parmi des dizaines en circulation, je suis sûr que vous avez entendu parler de Cryptolocker, TeslaCrypt, Petya, CERBER ou encore Locky.
  </p>

  <p>
    On trouve même des kit Ransomware clé en main sur les darknets, l'auteur vous demandera juste une commission sur les rançons !
  </p>

  <h3>
    Comment ça marche ?
  </h3>

  <p>
    Comme les virus, la plupart des ransomwares ne cherchent pas à exploiter une vulnérabilité, pourquoi s’embêter ? Ils s'arrangent juste pour vous faire ouvrir le virus tout seul, comme un grand !
  </p>

  <p>
    La méthode la plus courante est via phishing&#8230;pour plus de détails, il y a un épisode SECompris dédié au sujet, la maintenant&#8230;pas le temps. Certains se servent d'exploit kit qui leur permettent d'exploiter une liste de vulnérabilités les plus répandues
  </p>

  <p>
    Quand c'est le cas vous avez rien à faire pour qu'ils vous infecte, ou presque. Genre afficher une page web, ça peut suffire.
  </p>

  <p>
    Pour les plus vicieux, ils peuvent aller jusqu'à corrompre le site officiel d'un logiciel très utilisé, et infecter l'installeur. C'est rare mais c'est déjà arrivé, genre le ransomware KeyRanger sur Mac qui infectait le client bittorrent Transmission.
  </p>

  <p>
    Une fois lancé, l'attaque se passe en 3 phases :
  </p>

  <p>
    <strong>Phase 1 : Reconnaissance.</strong> Le virus scan les répertoires les plus intéressant à chiffrer, genre mes documents, le bureau, le répertoire backup etc.
  </p>

  <p>
    Les plus bourrin scan l'ensemble des fichiers accessible, y compris les partage réseau, ça fait mal en entreprise.
  </p>

  <p>
    Le scan permet de dresser une liste des fichiers intéressant à chiffrer. L'idée c'est d'aller vite, il faut chiffrer suffisamment avant de se faire gauler par l'utilisateur.
  </p>

  <p>
    On va donc se concentrer sur des fichiers pas très gros mais irremplaçables, genre vos photos de vacance, l'enregistrement des premiers pas du petit dernier, ou vos documents de travail.
  </p>

  <p>
    <strong>Phase 2 : On chiffre</strong> tout ça, et vite ! Avant c'était souvent des algos de chiffrement maison. Par maison, faut comprendre pourris.
  </p>

  <p>
    Faut croire que les attaquants en ont marre de se faire flinguer leur travail par des chercheurs en sécu qui craquent leur algo de chiffrement. Du coup ils se mettent de plus en plus à utiliser des solutions de chiffrement professionnelles, genre AES. Je ne vais pas détailler, pas le temps. Dites-vous juste que si c'est de l'AES, vous pouvez vous gratter pour pouvoir récupérer vos fichiers sans la clé.
  </p>

  <p>
    <strong>Phase 3 : On vous file la note de rançon</strong>
  </p>

  <p>
    Les plus bourrins changeront le fond d'écran pour une variante de &laquo;&nbsp;Haha on t'as bien fumé ! Maintenant, payes !&nbsp;&raquo; Les autres poseront sur votre bureau, ou dans chaque dossier chiffré, une note explicative sur qui payer, combien, et comment.
  </p>

  <p>
    Oui, comment, parce que payer en bitcoin sur le réseau TOR, ce n’est pas vraiment à la portée du premier clampin. Certains attaquants vont même jusqu'à mettre en place des services de hotline pour faire support technique !
  </p>

  <p>
    &nbsp;
  </p>

  <h3>
    Comment se protéger ?
  </h3>

  <p>
    &nbsp;
  </p>

  <p>
    Bon, je vais passer très vite sur comment se protéger des virus de manière générale, ça mériterait un épisode de SECompris complet et&#8230;Je vous ai déjà dit qu’on n’avait pas le temps ?
  </p>

  <p>
    En très bref, c'est le classique ABC :
  </p>

  <p>
    <strong>A : On Install un Antivirus</strong>, avec le scan en temps réel
  </p>

  <p>
    <strong>B : On Baintien à jour</strong>, son système, son antivirus, ses programmes et plugins courants genre Chrome, Firefox, Flash, PDF, Office.
  </p>

  <p>
    <strong>C : Le Comportement</strong> : On évite d'aller sur les sites de porn louche, de cracking de jeu ou logiciel, de téléchargement illégaux, etc.… ou si on y va on se protège et on la joue méfiant.
  </p>

  <p>
    Oui, ok, Maintien à jour, ça commence avec un M et pas un B, mais après ça pète mon moyen mnémotechnique avec ABC alors&#8230;Alors considérez que je vois enrubé, boila !
  </p>

  <p>
    Passons donc aux méthode spécial anti ransomware, j'en vois 3 :
  </p>

  <p>
    <strong>1 : L'heuristique</strong>, ou analyse comportementale, si un programme se comporte comme un ransomware, on bloque tout et on demande à l'utilisateur si c'est normal.
  </p>

  <p>
    Bah oui parce que des programmes qui accèdent en écriture à tous vos documents perso, les modifie, écrase les originaux, et tout ça super vite, ça n’arrive pas tous les 4 matins !
  </p>

  <p>
    La plupart des antivirus récent ont un module qui fait ça, que ce soit Bitdefender, McAfee, Kaspersky, vous avez l&#8217;embarras du choix
  </p>

  <p>
    <strong>2 : Le pot de miel</strong>, on place un peu partout sur le système de faux fichiers qui ont l'air alléchant pour le virus, et on attend qu'un programme essaye d'y toucher.
  </p>

  <p>
    L'utilisateur n'a aucune raison d'ouvrir ces faux fichiers, donc quand ça arrive, ça pue. Encore une fois, on bloque tout et on demande si c'est normal. En programme gratuit qui utilise cette méthode, on a par exemple Ransomfree, en payant il y a CryptoPrevent en version premium.
  </p>

  <p>
    <strong>3 : Les sauvegardes</strong>, parce que le meilleur moyen de pas avoir à payer la rançon c'est encore de pouvoir restaurer ses fichiers tout seul.
  </p>

  <p>
    Bien sûr il faut faire des sauvegardes fréquemment, mais de façon asynchrones et déconnectées. Si le backup se fait en temps réel ou qu'il est stocké sur un autre disque dur accessible depuis le PC infecté, le ransomware se fera un plaisir de le chiffrer aussi !
  </p>

  <p>
    Certaines solutions de backup synchrones gardent toutefois les anciennes versions des fichiers écrasés, c'est le cas de Dropbox. Vous devriez pouvoir restaurer vos fichiers même si le ransomware passe dessus.
  </p>

  <h3>
    J'ai été ransomwarisé, je fais quoi ?
  </h3>

  <p>
    &nbsp;
  </p>

  <p>
    Aïe, je vous le cache pas, ça pue. Si vous aviez fait des backups comme expliqué juste avant, alors c'est simple, vous restaurer les données et voilà.
  </p>

  <p>
    Pensez à désinfecter la machine avant, sinon vous risquez d'être bon pour recommencer. Pour les paranos, le meilleur moyen d'être sûr qu'il ne reste plus rien, c'est de formater et tout réinstaller.
  </p>

  <p>
    &nbsp;
  </p>

  <p>
    Si vous n’avez pas de sauvegarde, vous avez plus qu'a prié pour que l'attaquant ai fait son boulot comme un sagouin. En effet si le chiffrement est pourri, quelqu'un créera peut-être un programme pour retrouver la clé sans payer. Pas mal de chercheurs en virus se donnent du mal pour ça quand c'est possible. Autrement si l'attaquant se fait pincer par la police, les autorités essayent parfois de mettre à disposition du publique l'ensemble des clés de chiffrements qu'il avait en stock.
  </p>

  <p>
    Certaines boites d'antivirus en profitent pour faire des &laquo;&nbsp;decryptor&nbsp;&raquo;, petit outil pour savoir quelle clé est la vôtre et déchiffrer vos documents. Ne rêvez pas trop cependant, ça n'arrive pas souvent. Dans le doute, les plus patients peuvent garder leurs fichiers de côté et prier pour qu'une tel solution sorte. Le site <a href="https://www.nomoreransom.org" target="_blank">No More Ransom</a> est une bonne source pour ce genre d'outils.
  </p>

  <p>
    Enfin, la dernière méthode est bien sûr de payer&#8230;Alors la réponse officielle c'est ne payez jamais, il ne faut pas collaborer avec les terroristes.En effet ça finance leur activité, les encourage à recommencer, et vous n'avez qu'une garantie qu'ils respecteront leur part du marché et vous enverrons la clé.
  </p>

  <p>
    Bon, ne soyons pas faux-cul, c'est facile à dire quand c'est pas nos photos des 10 dernières années qui sont bonnes à jeter à la poubelle. La plupart du temps payer, ça fait mal au portemonnaie&#8230; mais ça marche.
  </p>

  <h3>
    Épilogue
  </h3>

  <p>
    &nbsp;
  </p>

  <p>
    Voilà, c'est fini pour aujourd'hui, j'espère que ce nouveau format vous a plu ! Celui-ci ne sera que sur YouTube, le rythme effréné ne colle pas trop avec le principe du podcast.
  </p>

  <p>
    Comme d'habitude n'hésitez pas à me laisser des commentaires, positifs comme négatifs, tant que c'est constructif, je suis content&#8230;tif :).*
  </p>

  <p>
    Sur ce je vous dit à bientôt pour un prochain SECompris, SEClair ou un prochain podcast du Comptoir Sécu !
  </p>

  <p>
    Il est temps de fermer le comptoir, à plus tard !
  </p>

  <p>
    *Cette blague a été retiré de l'enregistrement, j'avais honte 😀
  </p>
</div>

<div>
</div>
