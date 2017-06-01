---

title: "Tri-PIN : Une solution ingénieuse contre le shoulder surfing ?"
date: 2014-01-28T09:30:06+00:00
author: morgan


aliases: /2014/01/tri-pin-une-solution-ingenieuse-contre-le-shoulder-surfing/
views: 1317
image: /images/2014/01/tri-pin.png
categories:
  - Article
tags:
  - banker
  - keylogger
  - pin
  - screen nudging
  - shoulder surfing
---
Je suis tombé récemment sur une présentation du "futur" du code PIN.

Le problème d'un code PIN est que le pavé numérique est statique. La disposition des chiffres est toujours la même. Ainsi, si vous jetez un coup d'oeil par dessus l'épaule de quelqu'un tapant un code PIN sur son téléphone ou sur un distributeur automatique de billets (DAB), il est enfantin de voir le code saisi et de le retenir. C'est ce qu'on appelle une attaque par "shoulder surfing".

Dans le monde numérique, le problème n'est pas l'attaquant qui regarde par-dessus votre épaule, mais plutôt le keylogger qui enregistre les touches pressées sur votre clavier. C'est pourquoi les banques en lignes essayent de corser le travail des cybercriminels en vous obligeant à taper votre code sur un clavier virtuel dans lequel l'ordre des touches aura été mélangé de façon aléatoire.

Cela n'est malheureusement pas encore suffisant, des malwares de plus en plus malins (appelés "bankers") vont jusqu’à prendre une capture d'écran à chaque clic de souris réalisé dans la page web d'un site de banque en ligne. L'analyse manuelle des captures d'écran permet de retrouver le code PIN.

Une dernière étape consiste donc à vous donner un code PIN beaucoup plus long, disons 8 chiffres. Mais au lieu de vous demander les 8 à chaque fois, de n'en demandez que 4 au hasard parmi les 8. Le malware aura maintenant besoin de plusieurs séances d'espionnage pour avoir une chance de voir passer l'intégralité des 8 chiffres. Cela ne fait que retarder l'échéance, mais c'est toujours une bonne chose !

Tri-PIN est une variante de ce concept. Au lieu de demander 4 chiffres parmi 8, chaque touche peut correspondre à 3 informations. En effet chaque touche est une combinaison entre :

  1. un chiffre
  2. une couleur
  3. un symbole

Votre code PIN n'est plus 1418, mais par exemple "Rouge", 3, "Rond", 0. À chaque nouvelle connexion, les couleurs et symboles sont mélangés et disposés aléatoirement sur le clavier.



J'aime le concept, et le trouve suffisamment simple pour être accepté par la majorité des utilisateurs. Je ne comprends pas trop pourquoi par contre les chiffres ne sont pas également redistribués aléatoirement. J'imagine que c'est pour rester agréable à utiliser après l'authentification dans le cas d'un DAB lorsque l'on fait un retrait avec une somme non disponible par défaut sur l'écran.

Le (très) gros problème de cette méthode est qu'il implique :

  1. un changement des codes PIN des usagers
  2. un changement des distributeurs de billets ! Il faut un pavé numérique avec un écran LCD dans les touches ou bien un pavé tactile pour le remplacer.

C'est pourquoi je doute que l'on retrouve cette technologie implémentée de si tôt pour le grand publique. D'autant plus qu'il faudrait imposer une mise à jour globale. En effet, imaginons que Société Générale adopte le système et mette à jour tous ses DAB, ce qui est déjà improbable. Comment allez-vous faire pour retirer de l'argent dans un distributeur Caisse-D’épargne qui ne possède pas ce dispositif ?

Par contre, je serai ravi de voir ce genre de dispositif mis en place sur nos terminaux mobiles. Cela est facile à mettre en oeuvre, une simple mise à jour logicielle suffit. Le "shoulder surfing" et "screen nudging" (retrouver le code ou le schéma grâce aux traces de doigt sur l'écran) sont des problèmes courant sur ce type d'appareil.

Et vous, qu'en pensez-vous ?

Source: [Gizmodo](http://gizmodo.com/the-future-of-pin-could-involve-color-and-shape-buttons-1509670607)
