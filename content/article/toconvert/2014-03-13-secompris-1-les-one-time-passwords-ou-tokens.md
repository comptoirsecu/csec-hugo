---

title: "SECompris #1 : Les One Time Passwords, ou 'tokens'"
date: 2014-03-13T09:25:05+00:00
publisher: morgan


aliases: /2014/03/secompris-1-les-one-time-passwords-ou-tokens/
views: 7772
image: /images/covers/2014-03-secompris-token.jpg
categories:
  - Article
tags:
  - Authentification forte
  - Authy
  - keylogger
  - otp
  - SECompris
  - Token
---
_SECompris est une nouvelle série de billets ayant pour objectif d’expliquer un terme, une technologie, ou un principe de sécurité au plus grand nombre._

Vous avez sûrement déjà entendu parler de **One Time Passwords** (mot de passe à usage unique), d’OTP ou encore de tokens dans le cadre d’une authentification forte. Cette option de sécurité est proposée de plus en plus couramment sur les comptes internet considérés comme « sensible », en complément du mot de passe.

![battlenet_authenticator](/images/2014/03/battlenet_authenticator.png)

Il vous est demandé de fournir à chaque connexion un code, souvent sous forme d'une suite de 6 a 8 chiffres, fourni par un logiciel sur votre téléphone ou par un porte clé souvent appelé "token".

  * Google a été un des pionniers à le proposer sur Gmail avec [Google Authenticator](https://play.google.com/store/apps/details?id=com.google.android.apps.authenticator2&hl=fr)
  * Blizzard a été un des premiers à le proposer dans le monde du jeu vidéo avec son [Battle.Net authenticator](https://eu.battle.net/support/fr/article/faq-authenticator-battlenet), l’initiative est maintenant très courante dans le monde du MMO, très sujet aux hacks de compte
  * Les banques en sont aussi friandes, elles ont par contre tendance à préférer vous proposer la version « matériel »

## Matériel ? Logiciel ?

On retrouve deux formes de token, la version matérielle et la version logicielle.

La version matérielle prend généralement la forme d’un porte-clés avec un écran LCD. Le token a l’avantage d’être totalement déconnecté et à priori inaccessible par un attaquant (en dehors d’un vol de l’objet en question). Certains demandent d’ailleurs un code PIN avant d’afficher le fameux sésame temporaire pour pallier à ce risque.

Un des leaders sur le marché des tokens physique est RSA SecureID.

![Bon, ok, mauvais exemple vu l'affaire RSA révélée par Snowden !](/images/2013/12/SID700.gif)

Leur gros désavantage est leur encombrement. Un token, ça va, quand vous commencez à avoir 5 accès en authentification forte, voir plus, vous comprendrez vite que ce n’est pas très « scalable ».

La version logicielle est, comme son nom l’indique, sous la forme d’une simple application, généralement pour smartphone. Certains éditeurs comme Blizzard ont leur propre logiciel avec le Battle.Net authenticator, d’autres respectent les standards et fournissent uniquement les informations nécessaires à l’ajout d’un nouvel OTP dans votre token logiciel préféré. Ces applications sont par exemple Google Authenticator et Authy.

![authy](/images/2013/12/authy.jpg)

Les avantages et désavantages sont exactement l’opposé des versions matérielles. Plus aucun problème pour posséder 2 , 50 , 200 tokens. Par contre, vous vous retrouvez avec un OTP potentiellement accessible par des malwares, et donc plus vulnérable.

Je reste tout de même partisan de l’utilisation de ce type de dispositifs, tout simplement parce que le token physique est pour moi trop contraignant/couteux pour l’utilisateur, je préfère que l’utilisation de tokens se généralise de façon « non optimale » plutôt que cela reste marginalisé.

**Je pense que l’avenir sera de déporter ce type de produit sur les « objets connectés »**, toujours sans être parfait, ils seront encore un peu plus isolés, ce qui réduit d’autant la surface d’attaque.

## Comment ça marche ?

Les principes cryptographiques sur lesquels l’algorithme se repose sont complexes, mais le principe général est extrêmement simple. Il y a deux acteurs, « _Token_ », le programme ou le matériel vous fournissant le mot de passe à usage unique, et « _Serveur_ », le serveur d’authentification vous demandant le fameux OTP.

Lors de l’initialisation, _Serveur_ et _Token_ ont partagé un secret. Pour les tokens physiques, cela se fait avant qu’il vous soit livré. Pour les tokens logiciel, cela consiste généralement en une phase d’enrôlement sur le site du compte à protéger qui vous affiche un QR code que vous devez prendre en photo via l’application d’OTP. Au final ce QR code n’est rien autre que le secret plus parfois quelques informations telles que l’algorithme qui sera utilisé pour générer les OTP.

![totp-example](/images/2014/03/totp-example.png)

_Token_ et _Serveur_ possèdent donc maintenant un même secret « _S_ ». Lors de votre prochaine connexion sur « _Serveur_ », il va falloir vérifier que vous possédez bien « _Token_ ». On ne peut pas se contenter d’envoyer _S_, cela reviendrait à tout simplement posséder deux mots de passe, _S_ pourrait donc être volé par un attaquant au même titre que votre mot de passe habituel, tout l’intérêt serait perdu.

Mais comment faire savoir à Serveur que nous possédons bien _S_ dans notre _Token_ sans l’envoyer ? Il suffit de poser une question auquel il n’est possible de répondre que si l’on connait _S_.

Il est difficile de trouver ce genre d’exemple sans partir dans des mathématiques. Imaginons une fonction F, prenant deux arguments A et B. Lors de l’initialisation, Token et Serveur ont convenu de ce que serait la fonction F. À chaque connexion, Serveur va donc nous envoyer le paramètre A en le choisissant au hasard. Et nous allons lui répondre le résultat de _F (A, S)_.

  1. Comme le serveur connait aussi le secret _S_ et la fonction F, il lui suffit de refaire l’opération de son côté et de versifier la réponse. Si la réponse est la même, il est sûr que nous connaissons _S_, et que nous sommes donc la bonne personne
  2. Si un attaquant écoute l’échange, il ne verra que A et le résultat. Même s’il connait la fonction F utilisée (qui n’a aucunement besoin d’être secrète), il n’a pas S et il ne peut le retrouver à partir du résultat (la fonction a été choisie pour empêcher cela)
  3. Il ne pourra pas non plus réutiliser le résultat. Lors d’une prochaine tentative de connexion, le serveur enverra une autre valeur, disons Abis, et le résultat attendu ne sera donc pas le même.

![totp-function](/images/2014/03/totp-function.png)

Si vous voulez un exemple plus terre à terre, il y a par exemple la fonction « Grid » de LastPass. Celui-ci vous fournit en guise de token une grille remplie de nombre, chaque case est identifiée par ses coordonnées, un peu comme à la bataille navale.

![lastpass-grid](/images/2014/03/lastpass-grid1.png)

Lors d’une connexion, le serveur vous demande de fournir les chiffres de certaines cases, disons **A5, F9, C6 et Z3**.

Si vous répondez **bfnh**, c’est donc que vous possédez la grille.

Ce système, plus simple, n’est pas parfait, il est techniquement possible à l’attaquant de reconstituer petit à petit la grille au fil du temps, jusqu’à en avoir assez pour un jour pouvoir répondre à la question à votre place.

Si vous avez compris ceci, vous avez compris comment fonctionne un « challenge response » !

Les COTP (Counter-based OTP) se basent exactement là-dessus, sauf que le serveur n’envoie même plus la variable A, celle-ci est un compteur, incrémenté à chaque utilisation du token à la fois chez le Token et chez Serveur. La première connexion se basera donc sur F (1, S), puis F (2, S), etc. Un exemple de produit basé sur le COTP est par exemple la [Yubikey](http://www.yubico.com/products/yubikey-hardware/yubikey/), ce qui lui permet d'être totalement dépourvu de batterie.

![YubiKey-NEO-+-finger](/images/2014/01/YubiKey-NEO-+-finger.jpg)

Les TOTP (Time-based OTP) ne se basent pas sur un compteur, mais sur l’heure. Le mot de passe à usage unique est donc valable pendant une courte fenêtre de temps, disons 1 minute pour l’exemple. Si vous vous connectez le 10 mars 2014 à 14h35 il enverra le résultat de F (10/03/2014 – 14 :35 :00) au serveur. À 14h36, la variable utilisée aura changé et le résultat précédent ne sera plus valable.

Ce dernier système est le plus couramment utilisé. Le TOTP le plus connu est sûrement RSASecureID :<figure id="attachment_399"  >

![Bon, ok, mauvais exemple vu l'affaire RSA révélée par Snowden ](/images/2013/12/SID700.gif)<figcaption >Bon, ok, mauvais exemple vu l'affaire RSA révélée par Snowden !</figcaption></figure>

## Quel est l’intérêt ?

Rappelons les fondamentaux, il existe **3 facteurs d’authentification possibles** :

  1. Ce que l’on connait : un mot de passe, un code PIN, appelez-le comme vous voulez
  2. Ce que l’on est : la biométrie, comme l’empreinte digitale ou la reconnaissance rétinienne
  3. Ce que l’on possède : les fameux « tokens »

L’extrême majorité de nos accès informatique est régie par ce que l’on connait. Le problème étant que le mot de passe est très faillible :

  * Nous sommes submergés par les comptes à créer et retenir, et il est donc très tentant de réutiliser toujours les mêmes mots de passe, et de mettre des mots de passe court, facile à retenir et surtout à saisir. La mode des smartphones n’a rien amélioré, c’est même pire, il est très désagréable de taper un mot de passe complexe avec un clavier tactile
  * Un mot de passe est très facile à capturer par un malware, ceux qui le font sont tellement courants qu’ils ont même un nom propre, on les appelle les keylogger, généralement classés dans la grande famille des trojan, ou cheval de Troie. Un keylogger « surveille » les touches du clavier et retient les mots de passe saisis pour les envoyer à l’attaquant. Il suffit donc d’une connexion sur un ordinateur infecté pour fournir nos précieux mots de passe.

C’est là que vient l’authentification « forte », ne pas demander un, mais deux facteurs différents. Dans le cas présent : un mot de passe **ET** un token.

Comme on l’a vu au-dessus, le token est conçu pour que la preuve de possession change à chaque connexion. De ce fait, le keylogger va obtenir notre mot de passe et le mot de passe à usage unique (OTP) généré par le token lors de la connexion. Seulement lorsqu’il essayera de se connecter de son côté, le mot de passe à usage unique ne sera plus valide.

Le seul moyen pour lui d’assurer une connexion à volonté sur notre compte est de posséder notre token, ou, plus techniquement, la clé secrète qu’il utilise pour générer ces mots de passe à usage unique. Pour un token matériel, c’est conçu pour être mission impossible sans recourir à un vol du fameux porte-clés.

Pour un token logiciel, cela implique de pirater l’ordinateur utilisé par la connexion ET l’ordinateur (oui, votre smartphone est un ordinateur) hébergeant le token. On se rend tout de suite compte que la tâche est bien plus ardue.

## Conclusion

L’OTP est un mécanisme simple et éprouvé permettant de renforcer efficacement la sécurité d’un contrôle d’accès. C’est une fonctionnalité qui nécessite peu de moyens pour être mise en place du côté du fournisseur de service et qui ne nécessite pas de matériel spécifique pour s’identifier à la différence de la biométrie.

Enfin, contrairement à la biométrie, un OTP se change très facilement, comme un mot de passe, si jamais votre token était compromis, il est très simple pour l’utilisateur de réinitialiser le secret et donc d’invalider les bénéfices du vol.

L’OTP est un superbe palliatif à un mot de passe faible, même si je ne conseillais jamais à quelqu’un d’utiliser un mot de passe facile à deviner, si la personne ne peut se résoudre à retenir un mot de passe complexe, l’ajout d’un OTP réduira très fortement les risques. De la même manière, se connecter sur un équipement que l’on ne maîtrise pas, comme l’ordinateur d’un cybercafé, peut se faire bien plus sereinement.

Bien sûr, le token n’est pas parfait, sa perte est souvent très problématique pour l’utilisateur, comment prouver son identité auprès du fournisseur ? Si l’opération de réinitialisation est trop peu exigeante, l’attaquant se contentera de l’exploiter. Si la réinitialisation est trop simple voir impossible, l’utilisateur sera traumatisé par la perte définitive de ses accès et, croyez-moi, il ne sera pas prêt de recourir de nouveau à une solution d’authentification forte.

C’est pour cette raison que je soutiens des initiatives comme [la synchronisation du token sur plusieurs terminaux comme le propose Authy](http://blog.authy.com/multi-device). Même si les puristes crieront au scandale, je préfère un bon compromis au rejet de la sécurité.

J’espère que ce billet aura été informatif et compréhensible, peut-être vous a-t’il convaincu à activer l’authentification forte sur votre compte mail ? Votre gestionnaire de mot de passe ? votre banque ? Si vous avez un token sur ces 3 services, vous faites partie du haut du panier des utilisateurs.
