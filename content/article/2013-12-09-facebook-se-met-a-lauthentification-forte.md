---
title: Facebook se met à l’authentification forte
date: 2013-12-09T09:30:51+00:00
author: morgan
date: 2013-12-09
aliases: /2013/12/facebook-se-met-a-lauthentification-forte/
image:  /images/2014/01/url12.jpg
categories:
  - Articles
tags:
  - Biométrie
  - Duo Security
  - Facebook
  - otp
  - Yubico
---

Facebook annonce la possibilité de renforcer sa sécurité en intégrant les systèmes d’authentification forte [Yubico](http://www.yubico.com/) et [Duo Security](https://www.duosecurity.com/).

Petit rappel de rigueur, on appelle authentification forte une authentification se basant sur à minima deux facteurs parmi les 3 suivants :

  1. ce que l’on connaît, par exemple un mot de passe
  2. ce que l’on possède, l’exemple le plus parlant est la clé de votre porte d’entrée
  3. ce que l’on est, comme une empreinte biométrique

![illustration](https://www.comptoirsecu.fr/wp-content/uploads/2014/01/two-factor-authentication_Header-and-footer.png)

Yubico est maintenant depuis plusieurs années sur le marché, j’ai d’ailleurs encore dans un de mes tiroirs une de leurs premières “Yubikey”, sorte de petite clé USB se faisant passer pour un clavier auprès du système d’exploitation, et donc fonctionnant sans la nécessité d’installer le moindre driver sur tout ce qui peut gérer un clavier USB.

Dans cette clé réside un générateur d’OTP (One Time Password ou mot de passe à usage unique) basé sur un compteur et un secret partagé avec l’autorité d’authentification. Ce système a l’avantage de ne pas nécessiter de batterie et avoir une très grande durée de vie.

![Yubikey neo](https://www.comptoirsecu.fr/wp-content/uploads/2014/01/YubiKey-NEO-+-finger.jpg)

Je vous avoue avoir un peu délaissé ces token physiques ces derniers temps, au profit de leur pendant logiciel tel que Google Authenticator. Tout simplement car cela permet d’enlever un objet de ma poche et que de toute façon, comme tout bon drogué à la technologie du 21siecle, je ne me sépare jamais de mon smartphone à moins d’y être contraint !

Je suis très sceptique de leur nouveau format nano, dépassant à peine du port USB et nécessitant juste d’être effleurer pour cracher son précieux sésame. C’est le genre d’objet qu’on va très facilement “oublier”, volontairement ou non, dans le port de la machine. C’est certes moins risqué qu’un simple mot de passe robuste tel que “password123», mais pas idéal.

![Yubikey Nano](https://www.comptoirsecu.fr/wp-content/uploads/2014/01/YubiKey-Nano-+-lanyard.jpg)

Duo security est une approche assez différente, ou ici le téléphone à bien un rôle central. L’outil peut être un token software comme Google authenticator, ou peut vous simplifier encore plus la vie avec un système de notification ou il faut appuyer sur un gros bouton vert au lieu de rentrer à la main un code à 6 chiffres. Il sait également jouer dans la cour des facteurs biométrique en vous identifiant par la voix. Il peut s’adapter aux gens n’étant pas encore passé à l’être du smartphone par l’envoi du SMS, et peut bien entendue combiner tout ça pour faire de l’authentification à “3 facteurs”.

![iphone android push](https://www.comptoirsecu.fr/wp-content/uploads/2014/01/iphone-android-push-for-blog1.png)

Je suis plus attiré par ce genre de solution, j’avais d’ailleurs il y a environ 2 ans eu l’idée de ce genre de produit, avant de me décourager en voyant que cela n’avait rien de révolutionnaire et était déjà entrepris par des entreprises comme phonefactor, qui a d’ailleurs été racheté par Microsoft il y a peu !

Quoi qu’il en soit, je continue de penser que la vérité est encore ailleurs. Que le mot de passe ne doit pas être complété, mais totalement remplacé, au profit de solution à minima base sur le token (comme la nfcring) ou le token complète de la biométrie (comme le nymi). Certains puristes vous diront qu’il ne faut pas supprimer ce facteur du mot de passe, que cela permettrait d’authentifier quelqu’un a son insu en se rapprochant discrètement de son token authentifiant. Je répondrai a ça qu’un simple système de confirmation consciente ou de notification pallierait à la grande majorité des attaques, et que de nos jours le mot de passe est tellement peu apprécié et mal exploité par l’utilisateur, qu’il apporte plus de risque que de sécurité à nos systèmes.
