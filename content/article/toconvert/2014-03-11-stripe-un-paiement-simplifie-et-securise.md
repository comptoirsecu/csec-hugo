---

title: "Stripe, Un paiement simplifié et sécurisé ?"
date: 2014-03-11T20:49:27+00:00
author: justin


aliases: /2014/03/stripe-un-paiement-simplifie-et-securise/
views: 16538
image: /images/2014/03/stripe.png
categories:
  - Article
tags:
  - banque
  - paiement
  - paypal
  - stripe
---
Aujourd'hui je vais vous parler d'un concurrent au très célèbre Paypal: Stripe.

C'est toujours assez complexe et pas vraiment sécurisant d'être redirigé vers un autre site web lors d'un paiement sur Internet. Il faut toujours vérifier le certificat SSL et le numéro de carte bleue peut être récupérer si on a un key logger sur sa machine. De plus, payer sur son mobile peut être assez fastidieux.

Stripe permet avec son téléphone de payer via Navigateur web, mobile et tablette sans aucune redirection vers un site tierce. Le tout est réalisé en _Node.js_ ou via un plugin si c'est sur une application mobile.

Le concept est simple, le client voulant payer se connecte sur son compte via [OAuth2.0](http://oauth.net/2/) sur le site du vendeur. S’il n'a pas de compte, il est possible d'en créer un dans la foulée. Un sms contenant un code à usage unique est ensuite envoyé sur le téléphone de l'utilisateur. Il suffit de rentrer ce code sur le site pour payer. Le sms envoyé est lié à la commande et à la clé API du marchand.

[<img  alt="598px-Oauth_logo.svg" src="/images/2014/03/598px-Oauth_logo.svg_-300x300.png"     />](/images/2014/03/598px-Oauth_logo.svg_.png)

Évidemment, ils prennent une commission: 2.9% de la transaction plus 30 centimes. Cela peut paraitre beaucoup mais Paypal utilise exactement la même grille tarifaire. D'ailleurs un des principaux investisseurs est Peter Thiel qui n'est autre qu'un des fondateurs du géant du paiement électronique. Paypal s'intéresse beaucoup à cette solution de checkout puisqu'ils ont [annoncés](http://techcrunch.com/2014/01/13/paypal-debuts-a-simpler-native-checkout-experience-for-merchants-and-expand-beacon-internationally/) eux aussi développer quelque chose de similaire .

Du côté de la sécurité, la société est certifié [PCI Service Provider Level 1](http://www.visa.com/splisting/searchGrsp.do?companyNameCriteria=stripe). Ce qui indique un bon niveau de maturité de ce côté-là. Pour rappel, PCI/DSS est une certification qui impose beaucoup de règles de sécurité avec des audits réguliers. Les numéros de cartes bleues sont chiffrés en AES-256 et les clés de déchiffrements sont stockées sur des machines séparées. Il n'y a malheureusement pas suffisamment de description sur leur site pour se faire une réelle idée des mécanismes de sécurité employés de leur côté. Toutefois, il possible de signaler les failles trouver sur leur système via une adresse mail: security@stripe.com avec un engagement de réponse dans les 24h. Cependant, ils demandent de ne pas publier la description de la faille avant leur avoir communiqué et aucune contrepartie n'est annoncée.

[<img  alt="stripe mobile" src="/images/2014/03/stripe-mobile-260x300.png"     />](/images/2014/03/stripe-mobile.png)

Une attaque par l'homme du milieu me parait cependant faisable. C'est pour cela qu'ils conseillent aux vendeurs de faire attention aux autres javascripts qui seraient présent sur leurs sites. On peut très bien imaginé remplacer la clé API du vendeur par celle d'un attaquant afin de récupérer l'argent par exemple. Pour cela, il faudrait s'enregistrer sur leur site afin d'avoir une clé et le subterfuge serait, à mon avis, vite découvert.

Malheureusement, l'application n'est pas encore disponible en France mais cela ne saurait tarder. Attention, je vous déconseille de rechercher stripe sur le store android, j'ai eu des surprises NSFW 🙂



Sources:

<http://techcrunch.com/2014/03/05/stripe-debuts-a-new-checkout-experience-with-one-click-payments-for-mobile-and-web/?ncid=rss>

<https://stripe.com/help/security>
