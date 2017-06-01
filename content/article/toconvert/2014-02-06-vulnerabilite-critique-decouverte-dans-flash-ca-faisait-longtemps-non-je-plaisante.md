---

title: "Vulnérabilité critique découverte dans Flash, ça faisait longtemps... non je plaisante !"
date: 2014-02-06T09:25:13+00:00
author: morgan


aliases: /2014/02/vulnerabilite-critique-decouverte-dans-flash-ca-faisait-longtemps-non-je-plaisante/
views: 685
image: /images/2014/02/adobe-flash-player-hacked.jpg
categories:
  - Article
tags:
  - flash
  - remote code execution
  - vulnérabilité
---
Une vulnérabilité [Open bar sur ton PC avec tes droits utilisateurs](http://nakedsecurity.sophos.com/2014/02/04/adobe-fixes-critical-flash-flaw/" >vient d'être patchée</a> par Adobe dans le lecteur Flash. Cette vulnérabilité permettait l’exécution de code arbitraire à distance, ou "Remote Code Execution". Pour ceux qui ne sont pas à l’aise avec les termes techniques, quand vous lisez « Remote code execution » ça équivaut le plus souvent à « <a href="http://en.wikipedia.org/wiki/Arbitrary_code_execution) ».

C’est certainement la pire vulnérabilité possible. Avec cette vulnérabilité, l’attaquant peut lancer n’importe quelle commande sur l’ordinateur. Généralement, cela est limité par les droits du compte avec lequel l’exécutable vulnérable a été lancé.

En entreprise, c’est généralement des droits de simple utilisateur, ce qui limite la casse. Par contre chez un particulier, l’extrême majorité des comptes créés possèdent des droits administrateurs pour plus de confort. Et là, ça peut faire très mal, très vite.

Comme d’habitude, Adobe Flash, Adobe Acrobat Reader et Java sont des cibles de premier choix pour les cybercriminels. Leur taux de présence sur les postes de travail est tellement important que toute vulnérabilité sur une version récente assure un jackpot à malware.

## Comment se protéger ?

Supprimez ces programmes de votre ordinateur. Voilà, fini.

... Bon, OK, c'est peut-être un peu extrémiste, cherchons des alternatives plus douces.

  * **Gardez active la mise à jour automatique de ces produits**. Au moins lorsqu'une vulnérabilité de la sorte arrivera vous serez protégé dès que le patch aura été produit.
  * **Ajoutez des extensions telles que [NoScript](https://chrome.google.com/webstore/detail/flashblock/gofhjkjmkpinhpoiabjplobcaignabnl" >Flashblock</a> ou <a href="https://addons.mozilla.org/fr/firefox/addon/noscript/) pour bloquer « par défaut » les objets flash dans les pages web.** Vous pouvez les réactiver en un click, whitelister certains sites de confiance, etc. Ça vous évitera le surf malencontreux sur un site véreux avec un exploit flash au chargement de la page.
  * **Désactivez Java par défaut.** Sérieusement, les sites qui en ont besoin sont rarissimes. Coupez l'addon par défaut, les 3 fois l’an où vous en avez besoin, réactivez-le.
  * **Utilisez un lecteur PDF alternatif.** Certains navigateurs ont conçu leur propre lecteur comme [Sumatra](https://mozillalabs.com/en-US/pdfjs/" >Firefox </a>et [Foxit PDF Reader](https://support.google.com/chrome/answer/1060734?hl=fr" >Chrome</a>. Vous pouvez aussi utiliser des alternatives comme <a href="http://www.foxitsoftware.com/Secure_PDF_Reader/) ou <a href="http://www.framasoft.net/article4407.html). Non, ils ne sont pas forcément plus « securisés », par contre ils sont moins ciblés. De plus, le lecteur PDF officiel gère une quantité pas croyable d’options dont on n’a rien à faire la plupart du temps, comme intégrer des vidéos ou du flash dans les PDF. 99 % des PDF que l'on ouvre sont lisibles par des lecteurs minimalistes. Moins d’options signifient moins de lignes de code et donc moins de vulnérabilités potentielles.

Vous avez d’autres conseils à donner ? Partagez-les dans les commentaires !
