---

title: "World of Warcraft et authentification forte : encore insuffisant pour décourager les pirates"
date: 2014-01-15T13:30:50+00:00
publisher: morgan


aliases: /2014/01/world-of-warcraft-et-authentification-forte-encore-insuffisant-pour-decourager-les-pirates/
views: 877
image: /images/covers/2014-01-battle.net-authenticator.jpg
categories:
  - Article
tags:
  - Authentification forte
  - Token
  - World of Warcraft
---
Une attaque sur les comptes World of War

[Cette attaque est intéressante](http://bit.ly/1gF9OHt), car elle est souvent ignorée ou sous-estimée. Lorsqu’on présente l’authentification forte, on explique souvent que la protection est maximale, car il faut posséder les deux facteurs pour s’authentifier. Ce n’est que partiellement vrai. Lorsque les deux facteurs d’authentification sont renseignés depuis le même équipement, il suffit de compromettre l’équipement en question.

Comment est-ce possible ? Voici le cheminement :

  1. Votre machine est infectée par un malware.
  2. Vous voulez vous connecter à WoW avec votre mot de passe plus un code à usage unique généré par l’application Battle.net Authenticator.
  3. Vous entrez votre mot de passe dans WoW, le malware l’enregistre et l’envoie à l’attaquant.
  4. WoW vous demande maintenant le code à usage unique, mais cette fois c’est le malware qui l’intercepte, l’envoi à l’attaquant.
  5. L’attaquant dispose de quelques secondes pour se connecter depuis chez lui avec le mot de passe et le code à usage unique fraichement reçu.
  6. Pendant ce temps le malware à mis un code erroné dans la fenêtre de WoW et vous avez donc une erreur au moment du login.
  7. Vous ressayez plusieurs fois, pensant faire une erreur de frappe, le malware continue de changer le code, l’authentification ne marche pas, vous pensez avoir un problème avec votre compte.
  8. Pendant ce temps l’attaquant vide votre inventaire et surtout vos précieuses pièces d’or</ol>
L’exemple est tiré de l’article, mais peut être adapté pour voler le code que vous recevez par votre banque par SMS lorsque vous faites un paiement 3DSecure.

Pour mitiger ce type d’attaque, il faut que le second facteur soit envoyé à minima par un autre équipement, le mieux étant qu’il soit carrément envoyé depuis un autre canal de communication. Par là j’entends, s’il est envoyé par deux ordinateurs différents, mais que les deux utilisent la même connexion réseau, un man-in-the-middle reste envisageable.

Si par contre on vous demande par exemple de répondre à un SMS en plus de renseigner votre mot de passe sur l’ordinateur, il faut maintenant compromettre votre ordinateur **ET** votre téléphone, ou votre connexion internet ET votre liaison GSM. Bien entendu, la sécurité absolue n’existe pas, cela reste théoriquement réalisable, mais beaucoup, beaucoup plus difficile à mettre en place.
