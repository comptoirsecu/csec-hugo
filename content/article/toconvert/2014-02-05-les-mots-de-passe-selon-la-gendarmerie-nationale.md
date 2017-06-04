---

title: "Les mots de passe selon la gendarmerie nationale"
date: 2014-02-05T09:25:21+00:00
publisher: morgan


aliases: /2014/02/les-mots-de-passe-selon-la-gendarmerie-nationale/
views: 2220
image: /images/covers/2014-02-password-gendarmerie.jpg
categories:
  - Article
tags:
  - authentification
  - complexité
  - gendarmerie
  - mot de passe
---
Je suis tombé sur ce tweet de la gendarmerie nationale récemment :

![tweet-gendarmerie](/images/2014/02/tweet-gendarmerie.png)

Alors oui, ce message est plein de bonnes volontés, et effectivement, tout ce qui y est dit est parfaitement juste.

De nos jours, un mot de passe doit au minimum faire 12 caractères pour être sur qu'il ne figure dans aucune rainbow table.

> Mais les rainbow table ça ne sert à rien si les développeurs ont fait correctement leur travail et ajouté un sel à notre mot de passe non ?

Oui, cependant arrêtez de penser que le développeur moyen fait les choses correctement. Les mots de passe stockés en clair se font de plus en plus rares, mais la majorité se contente de le hacher, sans sel, avec un algorithme qui n'a jamais été conçu pour être utilisé dans de la protection de mots de passe.

Oui, il faut éviter des mots de passe utilisant des données personnelles, telle que la date de naissance des enfants ou le nom du chien, pour éviter les attaques ciblées.

Oui, il faut changer de mots de passe régulièrement, à minima après qu'une fuite de donnée a été signalée sur le site en question.

Oui, il ne faut pas réutiliser son mot de passe pour qu'éviter qu'une attaque réussite sur le site X permette de piller les données de votre compte Y.

Oui, il ne faut pas noter ses mots de passe sur un post-it collé à l'écran, sous le clavier ou sous le téléphone.

![password-gendarmerie](/images/2014/02/password-gendarmerie.jpg)

Mais tout ça mis bout à bout, cela vous semble gérable au quotidien pour un utilisateur moyen ? Je peux comprendre la réticence pour la gendarmerie à faire de la publicité pour un coffre-fort de mots de passe commercial. Mais au moins, rappeler que ce genre de solutions existe ? Qu'elles se basent sur des modèles cryptographiques fiables ? Pourquoi ne pas conseiller un produit gratuit et audité par [Keepass](http://www.ssi.gouv.fr/fr/produits-et-prestataires/produits-certifies-cspn/certificat_cspn_2010_07.html" >l'ANSSI</a> comme <a href="http://keepass.info/) ?

J'ai fait le compte rapidement, j'ai plus de 350 comptes à mon actif sur Internet, 350.

D'accord, je suis quelqu'un de très connecté, je ne me sers probablement pas au quotidien des 9/10ème de ces 350 comptes, admettons. Mais même 1/10ème de ce patrimoine, cela fait 35 comptes.

35 mots de passe différents, de plus de 12 caractères, avec des caractères spéciaux, des chiffres et des changements de casse. 35 mots de passe qui n'ont aucun lien qui permet d'être retrouvé à partir de mon identité ou du nom du site. Cela vous semble réalisable de retenir cette bouillie de chiffres, de lettres et de ponctuations de tête ? Moi non.

[Dashlane ](https://www.dashlane.com/)fait son travail de communication correctement et a répondu au tweet, la première réaction rencontrée ? La voici :

![tweet-gendarmerie-reponse-dashlane](/images/2014/02/tweet-gendarmerie-reponse-dashlane.png)

Pourquoi faire peur aux gens ? Est-ce que la personne comprend au moins comment fonctionne un coffre-fort numérique tel que Dashlane ?

Avec un mot de passe décent et la demande d'un second facteur d'authentification lorsque le compte est utilisé depuis une IP inconnue, on a quelque chose de robuste, très robuste.

Une des attaques théoriques envisageables nécessite la prise de contrôle du poste de travail de l'utilisateur comme pivot pour se connecter à son compte sans demander de deuxième facteur. Il peut aussi ajouter un périphérique type smartphone sans token s’il a accès à votre boite mail pour valider l'ajout du périphérique.

Effectivement, dans ce cas-là, l'attaquant possède tous les mots de passe d'un coup. Mais entre nous, s’il a la main sur votre ordinateur ou votre boite mail, vous avez perdu, quelle que soit votre méthode de gestion des mots de passe...

Bref, je trouve cela dommage. Pour une fois que la sécurité apporte par la même occasion du confort, il ne fauter pas l'écarter des sensibilisations. Les coffres fort numériques sont encore trop méconnus des utilisateurs.
