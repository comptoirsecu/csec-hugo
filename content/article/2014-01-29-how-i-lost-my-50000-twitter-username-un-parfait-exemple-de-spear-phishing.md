---

title: "How I lost my 50,000$ Twitter username : Un parfait exemple de social engineering"
date: 2014-01-29T12:25:13+00:00
publisher: morgan


aliases: /2014/01/how-i-lost-my-50000-twitter-username-un-parfait-exemple-de-spear-phishing/
views: 683
image: /images/covers/2014-01-twitter-hack2.jpg
categories:
  - Article
tags:
  - godaddy
  - paypal
  - social engineering
  - twitter
---
Le Social engineering, ou ingénierie sociale, est une attaque ne se basant pas sur une vulnérabilité informatique mais sur une faiblesse dans un processus organisationnel ou sur la manipulation d'un être humain.

Naoki Hiroshima en a fait les frais il y a peu. Détenteur d'un compte Twitter très rare, @N, et qui vaut donc près de 50,000$, il est une cible de choix pour les attaques ciblées.

L’attaquant a mis en place une stratégie ingénieuse pour lui extorquer son compte :

  1. L’attaquant a réussi à obtenir de PayPal les 4 derniers numéros de la carte de crédit de Naoki
  2. Il a ensuite appelé le Registrar utilisé par Naoki, GoDaddy, pour prendre possession de son nom de domaine. Pour ça il a prétexté avoir perdu l’accès à ses mails, et a fourni pour prouver son identité... les 4 derniers numéros de la carte de crédit. Il lui ont aussi fait deviner les 2 premiers, qu’il a trouvés du premier coup. Pour information les 2 premiers chiffres sont conditionnés par le type de carte (Visa, MasterCard...) ce n’est donc pas un total hasard.
  3. prendre la main sur le compte GoDaddy lui a permis de prendre la main sur son email... Et oui Naoki utilisait une redirection pour son mail. Par exemple en possédant le domaine « monsite.fr », vous pouvez créer une adresse « contact@monsite.fr » qui n’est autre qu’une redirection vers « veritable-adresse@gmail.com » par exemple. L’attaquant a donc redirigé contact@monsite.fr vers « adresse-de-l-attaquant@gmail.com ».
  4. maintenant en possession de l’adresse utilisée par Naoki, il pouvait réinitialiser le mot de passe de n’importe quel compte souscrit par Naoki avec cette adresse.
  5. Naoki a eu le temps de changer le mail associé à son compte twitter, ce qui a empêché l’attaquant de le voler directement. Il a donc opté pour la prise du compte Facebook et a commencé à le faire chanter.
  6. Naoki a fini par céder, il a relâché le compte @N en échange de la restitution de ses comptes personnels.

Cette péripétie est une leçon importante. Nous l’avons déjà dit plusieurs fois, votre compte mail est le sésame ultime pour la quasi-totalité de vos comptes en ligne, car tous les services proposent une réinitialisation par envoi de mail. La subtilité supplémentaire ici était que l’adresse email dépendait du nom de domaine. La sécurité d’un système se résume à celle du maillon le plus faible. Dans ce cas c’était GoDaddy. Une authentification forte sur son compte Gmail ne l’aura pas protégé.

À l’heure actuelle, Naoki ne possède plus @N, Godaddy, PayPal et Twitter ont tous annoncé qu’une enquête était en cours auprès de leurs équipes. N’y voyez pas là la ferveur de dirigeants amoureux de la justice. L’article de Naoki a fait du bruit, ils cherchent juste à nettoyer leur réputation. Espérons que cela fonctionne pour Naoki.

_Source : [Naoki Hiroshima (Medium)](https://medium.com/p/24eb09e026dd)_
