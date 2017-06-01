---

title: "Have I Been Pwned : Cachez-moi ce leak que je ne saurais voir."
date: 2014-01-24T09:30:38+00:00
author: morgan


aliases: /2014/01/have-i-been-pwned-cachez-moi-ce-leak-que-je-ne-saurais-voir/
views: 1044
image: /images/2014/01/passwordscloud.png
categories:
  - Article
tags:
  - Authentification forte
  - Dashlane
  - HaveIBeenPwned
  - mot de passe
  - otp
---
<ins datetime="2014-01-23T20:22:48+00:00">Cet article était précédemment édité sur [Medium](https://medium.com/le-comptoir-secu/99a298943f2c), j'ai profité de cet import pour mettre à jour certaines sections.</ins>

Le site « HaveIBeenPwned » a fait pas mal de buzz ces derniers jours. Le principe est simple : renseignez votre adresse email, il vous dira si vous êtes présent parmi une des (nombreuses) fuites d’informations identifiantes que l’on a eues ces derniers temps (Adobe, Gawker, Sony…).

Je trouve ce genre d’outil très bon pour faire un peu de sensibilisation, s’il s’avère que le site trouve un compte à votre nom. Et même les plus paranoïaques peuvent se faire avoir!

![havveibeenpowned](/images/2014/01/havveibeenpowned.png)

Il semblerait qu’un compte m’appartenant a leaké chez Adobe. Je n’ai aucun souvenir de ce compte et il est tellement vieux que mon gestionnaire de mot de passe ne le connait pas. Je suis allé sur le site d’Adobe, ai tenté de me loguer en entrant un vieux pass que j’utilisais par défaut sur tous les sites que j’estimais peu sensibles : Bingo…

Par chance Adobe a eu la bonne idée de forcer un reset du mot de passe sur l’ensemble des comptes leakés, le mien en faisant partie, j’ai donc pu utiliser mon password manager pour générer quelque chose d’un tant soit peu solide.

Quitte à enfoncer une porte ouverte, ne faites pas comme j’ai pu faire dans le passer et éviter à tout prix d’utiliser un même mot de passe à différents endroits. Le site que vous estimez peu sensible aujourd’hui ne le sera peut-être plus demain.

Je suis également assez peu partisan des « méthodes maison pour générer un mot de passe » consistant à mélanger par exemple un noyau commun avec le nom du site (ou un dérivé). Ce genre d’astuces ont longtemps été conseillées pour faire des mots de passe à la fois uniques et faciles à retenir… du moins facile à retrouver. Voici quelques exemples en prenant gmail, facebook, et twitter :

  1. le très naïf : **passwordgmail** / **passwordfacebook**/**passwordtwitter**
  2. le presque malin : **passwordgml** / **passwordfcbk**/**passwordtwttr**
  3. le faux difficile : **password714mg**/ **passwordk00b3c4f**/ **password53tt1wt**

![ad-hamlet](/images/2014/01/ad-hamlet.png)

Le problème avec cette méthode est que, bien qu’elle soit efficace pour empêcher une tentative de réutilisation bête et méchante ou automatisée, elle ne devrait pas tenir bien longtemps face à un être humain prêt à passer 30 secondes pour décrypter votre précieux algorithme maison. Je suis sûr d’ailleurs que vous avez très vite trouvé les algorithmes sous-jacents à mes exemples. On peut bien sûr continuer de corser les choses, utiliser la position dans l’alphabet à la place de la lettre, ne prendre que quelques lettres à des positions bien précises, mais de manière générale je trouve que ces méthodes sont une illusion de sécurité.

Le pire, c'est que de nombreuses personnes pensent que c'est effectivement une bonne façon de faire et l'évangélisent. Pas plus tard que cette semaine je suis tombé sur [un article de PcINpact](http://www.pcinpact.com/news/85439-les-25-pires-mots-passe-annee-histoire-se-repete-encore-et-toujours.htm), qui nous donnent en infographique ceci :

![sensibilisation-securite](/images/2014/01/131904.png)

Quand un attaquant va tester des mots de passe, s’il est malin son attaque par dictionnaire inclura les "mutations" les plus courantes, à savoir :

  1. S’il y a une majuscule elle est probablement en première position ou au début d'un mot de façon générale
  2. S’il y a des chiffres ils sont probablement isolés des mots ou alors ils remplacent des voyelles dans celui-ci plus deux trois consonnes bien identifiées comme le "l" (aka le [1337 Sp34k](http://fr.wikipedia.org/wiki/Leet_speak))
  3. S’il y a un caractère spécial, c'est très probablement un point d'exclamation à la fin du mot de passe.

Enfin pour des attaques "ciblées", si jamais un mot de passe vous concernant a fuité sur internet pour par exemple twitter et que c'était "S3cur3 Twitt3r!", vous pouvez être sur que les premiers mots de passe qui seront tentés sur votre Facebook seront &lsquo;S3cur3 Fac3book!", puis "S3cur3 Faceb00k!", ou "S3cur3 F4cebook!"...

Relisez l'affiche maintenant...vous la trouvez toujours aussi avisée ?

En attendant un monde où nous pourrons enfin nous authentifier sans utiliser ces maudits mots de passe, j’ai deux conseils à donner :

## Utilisez un gestionnaire de mot de passe, et utilisez-le pour absolument tout ce sur quoi vous vous connectez.

J’utilise personnellement [la version 2.10 portable](https://www.dashlane.com/fr/cs/3ba2769c" >Dashlane </a>, parce qu’ils sont français et donc soumis aux lois européennes. Mais dans l’absolu il y en a plein d’autres, connus et reconnus, qui feront parfaitement leur affaire, pêle-mêle je vous laisse googler lastpass, onepassword, keepass… J’ai tendance à préférer les versions synchronisées en ligne pour des raisons de confort. Les plus paranoïaques se limiteront au local avec keepass, de préférence en se limitant à la version auditée par l’ANSSI, c’est-à-dire <a href=" http://www.ssi.gouv.fr/fr/produits-et-prestataires/produits-certifies-cspn/certificat_cspn_2010_07.html).

Je reste confiant avec les produits en ligne, les [zero Knowledge Privacy](http://www.techrepublic.com/blog/it-security/how-safe-are-online-password-managers/" >articles sur le sujet</a> vont aussi dans ce sens. Ils avancent tous le même mécanisme de confidentialité : la clé de chiffrement, dérivée de votre mot de passe maître, n’est jamais stockée ni même envoyée sur leurs serveurs. Tout le travail de chiffrement et de déchiffrement se fait en local par l’intermédiaire de leur plug-in ou, sur la version web, de JavaScript. Ce mécanisme s'appelle le Z<a href="http://zeroknowledgeprivacy.org/). Maintenant, il est vrai que je ne me suis pas amusé à auditer les fameux outils pour vérifier leurs dires.

L’avantage de ce type d’outils et qu’il permet de générer très rapidement des mots de passe complexes à « craquer » (à savoir, aléatoire, avec des majuscules, chiffres et caractères spéciaux, et plus de 12 caractères) sans jamais n’avoir ne à les retenir ni même les taper, l’add-on se chargera de ça pour vous.

## Pour les sites vraiment sensibles, complétez si cela est possible par un deuxième facteur d’authentification<figure id="attachment_399"  >

![Bon, ok, mauvais exemple vu l'affaire RSA révélée par Snowden !](/images/2013/12/SID700.gif)

L’authentification « forte » (2 facteurs), autrefois réservée aux banques en ligne, commence à se généraliser sur les sites sensibles (Twitter, Gmail, même Facebook s’y mettent!).

Vous avez beau avoir choisi un mot de passe impossible à retrouver (en un temps raisonnable) par une attaque par recherche exhaustive ou par dictionnaire, vous n’êtes pas à l’abri de le taper sur une machine infectée contenant un keylogger. Vous pouvez aussi, par inadvertance bien sûr, l'avoir saisi sur navigateur configuré pour retenir (en clair…) l’ensemble des mots de passe utilisés!

_Un keylogger est outil malveillant « écoutant » tout ce qui est tapé sur le clavier dans l’espoir de reconnaître des saisies de mot de passe pour ensuite les envoyer à son commanditaire. Certains keyloggers vont encore plus loin et retiennent la position de la souris lors des clics, prennent une capture d'écran régulière ou de la zone autour du clic,..._

C’est là que le token entre en jeu, votre mot de passe a beau être révélé, il est inutilisable sans l’ajout du mot de passe temporaire/à usage unique généré par votre précieux second facteur.

Le biométrique est quasi inexistant en dehors du passage de frontière et du monde de l’entreprise, ce que l’on retrouve le plus souvent est un token, qui prend la forme :

  1. Soit d’un petit porte-clés type RSA SecureID
  2. Soit d’un logiciel de type Google Authenticator.

Le premier est bien entendu le plus « inviolable ». En effet, le token est totalement indépendant, on ne risque pas d'installer un logiciel qui compromettrait la clé secrète qu'il utilise pour générer les mots de passe à usage unique comme c'est le cas avec le token logiciel. On a tendance à préférer, mois le premier, le pendant logiciel pour éviter d’avoir un trousseau de clés qui ne rentre plus dans la poche…

Pendant longtemps j'aurais conseillé les yeux fermés Google Authenticator. Il est gratuit, gère plusieurs OTP, et est entièrement en local ce qui rassure pas mal de gens. Oui seulement l'absence de backup, quand un produit survient, on le sent passer. Une mise à jour sur iPhone il y a peu à [involontairement supprimé tous les comptes enregistrés dans l'application](https://winauth.com/2013/09/04/iphone-google-authenticator-wipe-backup/). Ceux qui n'avaient pas de backup manuel via iTunes n'ont eu que leurs yeux pour pleurer.

J'ai récemment découvert [Authy](https://www.authy.com/). Je ne vous cache pas que la première fois où je suis tombé sur leur site, je me suis dit :

> "Ok, un énième logiciel d'OTP, qui n'apporte rien de plus que Google Authenticator, et qui en plus stock mes clés chez eux ? Kthxbye."

![authy](/images/2013/12/authy.jpg)

Depuis j'ai un peu creusé, et [Ce genre d'article ](http://blog.authy.com/backups" >ils stockent les clés de la même façon que les password manager</a>, ils n'y ont donc pas accès. Je pars bien sûr du principe que ce qu'ils disent sur leur site est vrai je n'ai pas audité l'application. <a href="http://shinynightmares.wordpress.com/2013/10/24/authy-cracking-encrypted-authenticator-backups/)rassure aussi sur la sécurité du dispositif. Il rappelle bien sur une évidence : même les meilleurs algorithmes de chiffrement se font casser en un rien de temps quand la clé de chiffrement est faible.

On y trouve aussi des fonctionnalités intéressantes, comme la possibilité d'intégrer une authentification forte sur votre [OpenVPN](http://blog.authy.com/wordpress" >WordPress</a>, [SSH](http://blog.authy.com/two-factor-ssh-in-thirty-seconds), ou votre <a href="http://blog.authy.com/openvpn).

Je reste frileux sur certaines fonctionnalités qui sacrifient un peu trop la sécurité sur l'autel du confort, je pense notamment au fait de [compte Authy du téléphone avec le navigateur via bluetooth](http://blog.authy.com/multi-device" >mettre son compte sur plusieurs appareils</a>. Ou d'interfacer le <a href="http://vimeo.com/71272779).



Cependant, il faut reconnaître qu'ils font les choses bien, toutes ces fonctionnalités sont en "opt-in" et ils font ce qu'ils peuvent pour réduire les risques associés.

**Ah oui, une dernière chose, tout gestionnaire de mot de passe digne de ce nom gère l’authentification forte, de la même manière, la quasi totalité des sites vous permettent de réinitialiser un mot de passe par le simple envoi d’un email.

Si l’utilisation d’un token vous fait soupirer et que vous souhaitez le mettre sur le moins de comptes possible, mettez-les à minima sur ces deux comptes.</strong>
