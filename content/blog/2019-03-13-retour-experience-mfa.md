---
title: "Retour d’expérience sur le déploiement de l’authentification multifacteur Microsoft en entreprise"
authors:
  - jil
  - morgan
date: 2019-03-03
image:  /images/covers/2019-03-03-rex-mfa.jpg
categories:
  - Article
tags:
  - password
  - mfa
  - rex
  - azure
---

*Avant-propos : l’article est centré sur l’environnement Microsoft, mais ne soyez point révulsé ! Il y a des choses à prendre pour tout le monde ;) Que ce soit Duo ou Okta ou d’autres…*

La transformation numérique des entreprises s’accompagne souvent d’un élargissement des facultés de connexion : il est non seulement possible d’accéder aux services depuis les locaux de l’entreprise, mais aussi depuis n’importe quel lieu, sans VPN. Cette tendance se dégage clairement dans le déploiement d’Office 365, mais aussi d’autres solutions fournies *as as service*.

Dans ce contexte, et afin de limiter les conséquences d’une campagne d’hameçonnage réussie, l’authentification multifacteur (MFA) est de plus en plus généralisée à tous les utilisateurs. Cet article partage l’expérience des membres du Comptoir Sécu, et leur point de vue.

# Les licences, mon capitaine !

Vous pouvez activer l’authentification en deux étapes (formulation pour l’utilisateur final retenue par Microsoft) pour tous les utilisateurs d’Office 365. Cette faculté couvre uniquement l’authentification aux services O365 (ce qui est déjà pas mal ! Vous le payez, pourquoi s’en priver ?).

Pour étendre le MFA à d’autres applications, il vous faudra les licences Azure AD Premium P1 ou supérieur. Les packs E3 EMS et — il me semble — Microsoft 365, contiennent une AAD P1 ; le pack E5 EMS, une AAD P2. Cela quelle que soit la topologie de déploiement.

# Topologie : ADFS à la maison ou Azure dans les nuages ?

La topologie de déploiement est un premier sujet à traiter. Parce que c’est galère à déployer, il est essentiel que l’outil de MFA soit le même pour Office 365 et les autres applications qui en auront besoin. Parmi ces applications, certaines seront hébergées par un tiers tandis que d’autres resteront bien au chaud dans votre centre de données.

Vous pouvez faire d’Azure votre fournisseur d’identité (si quelqu’un a fait ça, venez nous dire ce que vous en pensez et ce que ça nécessite comme licences !) ou utiliser Active Directory Federation Services (ADFS) déployé en interne.

L’ADFS est, à ma connaissance, le seul moyen de gérer un deuxième facteur différent de celui d’Azure, comme un certificat délivré par votre PKI sur un jeton physique (Yubikey, Feitan, Key-id, RSA). Si quelqu’un a joué avec ça, nous sommes également preneurs d’un retour !

## Oui, mais mon ADFS n’est pas encore en 2016

L’intégration d’Azure MFA dans ADFS *on prem* est gérée convenablement à partir d’ADFS 2016. Votre ADFS interne décidera dans quels cas déclencher la demande du deuxième facteur, et délèguera le traitement à Azure AD. Cela signifie que si vous êtes à la ramasse avec votre version d’ADFS, vous pouvez déjà sécuriser l’accès à Office 365 en laissant Azure AD exiger le deuxième facteur pour l’utilisateur et, dans un deuxième temps, monter votre ADFS de version et étendre la même expérience de MFA aux autres services connectés à votre ADFS.

# La configuration

## Démarrer

Dans le centre d’administration d’Office 365, sous le nouveau menu `Settings > Services & add-ins`, vous retrouverez l’accès au panneau de contrôle du MFA et les liens vers les docs. Prenez le temps de lire les docs.

Je le répète : prenez le temps de lire les docs. Les FAQ, les incompatibilités, etc.

Il y a plusieurs manières d’exiger le MFA. Si vous avez les licences AAD Premium, vous pouvez définir ça dans la partie accès conditionnel d’Azure AD. Vous avez différents critères qui vous permettent de décider si vous exigez ou non le MFA. Toutefois, vous voudrez jouer avec ces politiques après avoir configuré vos utilisateurs (quelqu’un a-t-il réussi à déclencher un enrôlement depuis l’accès conditionnel ?)

> **Aparté sur la confiance dans le réseau local**
	
> Il est parfois dit qu’on ne devrait exiger le deuxième facteur que depuis un lieu de connexion non maîtrisé. Quiconque a déjà vécu un test d’intrusion physique sait qu’il n’y a aucune raison de croire qu’un couple login,password fourni sur le réseau de l’entreprise provient de l’utilisateur concerné. Une fois que vous avez réussi à faire comprendre à l’utilisateur que le MFA est un procédé normal et régulier, exigez-le, quelle que soit l’origine de la connexion. Vous pourrez réduire le nombre d’occurrences avec Windows Hello for Business et le modern management de la flotte, où il deviendra possible de faire un choix basé sur la présence d’éléments de confiance (c’est bien le terminal connu pour cet utilisateur, son téléphone est à portée de Bluetooth et appairé, etc.).


Politique conditionnelle ou non, le compte utilisateur peut être dans trois états :

 * *disabled*
 * *enabled* : à la prochaine connexion avec un client qui le supporte, la procédure d’enrôlement est déclenchée et l’utilisateur doit la réussir pour accéder au service. Entre temps, la connexion par couple login,password est autorisée (pour les clients mails et Skype notamment).
 * *enforced* : plus d’authentification possible sans MFA ou application password. Un utilisateur qui réussit son enrôlement passe automatiquement de *Enabled* à *Enforced*.

## Durée de vie du cookie navigateur

Il est possible de laisser l’utilisateur demander, quand il est dans son navigateur, de ne pas avoir à effectuer à nouveau l’authentification en deux étapes depuis ce navigateur pendant *n* jours. Dans un cas, 30 jours a été bien vécu par les utilisateurs.

À noter qu’à partir de Windows 10 1803, l’OS intègre mieux l’authentification à Office 365 et permet de ne pas exiger d’authentification tout court avec Edge.

## Numéro de téléphone de récupération ou uniquement la notification mobile ?

Pas de numéro de téléphone pour vos comptes sensibles. Les attaques par portabilité de ligne opérateur sont fréquentes en attaque ciblée.

Pour les autres, c’est mieux que rien. Surtout, cela permet de permettre d’authentifier l’utilisateur qui change de téléphone, car la configuration de l’app Microsoft Authenticator n’est pas sauvegardable. Elle est liée au *device*. Donc si l’utilisateur supprime par mégarde l’app, il sera bloqué jusqu’à ce que le service desk l’aide à se réenregistrer.

De plus, les restrictions chinoises en matière de chiffrement font que  Microsoft Authenticator n’est pas disponible en Chine. Or, l’option d’activer ou non les fonctionnalités de SMS ou appel téléphonique est globale.

Vous devez aussi prendre en compte les personnes qui n’ont pas d’ordiphone, voire pas de téléphone mobile tout court (ou qu’ils n’en ont pas de professionnel et ne veulent rien installer dessus pour le boulot, liberté que le droit français garantit). S’ils n’ont pas non plus de ligne de bureau joignable depuis l’extérieur, il ne sera pas possible de les déployer avec Azure MFA. Il faut donc prévoir un groupe d’utilisateurs qui n’auront pas de MFA et n’auront pas accès aux services depuis des zones non contrôlées.

La notification mobile devrait être positionnée comme moyen préféré, en particulier si vous envisagez de bénéficier du monde sans mot de passe que prévoit Microsoft, où l’app mobile devient le premier facteur, et le mot de passe, le second. Cf. Windows Hello for Business.

## Pour ou contre les mots de passe d’application

Un mot de passe d’application est une chaîne de lettres aléatoires qui permet de contourner le MFA. C’est nécessaire pour la compatibilité de certains clients mails natifs des téléphones. Théoriquement, le mot de passe doit être saisi par l’utilisateur dans l’appli et il doit l’oublier. Sauf que, dans un souci de service, il pourrait venir y l’idée des agents du desk de le faire sauvegarder dans un fichier sur le bureau. Pourquoi ?

Parce qu’une messagerie native se connecte toutes les 15 secondes pour vérifier la présence de nouveaux messages. Ce faisant, dès que l’utilisateur passe en statut *Enforced*, se produit une attaque par force brute qui verrouille le compte pour l’IP du téléphone. D’où il s’en suit que quand l’utilisateur entre le mot de passe d’application, ça ne fonctionne pas, et il pense l’avoir mal tapé, puis comprend que le système est bogué et que tout ça, ça lui casse les pieds. Il finit par noter ce foutu mot de passe. Et jamais ne le supprime. 

La messagerie native d’iOS 11.x+ supporte l’authentification moderne d’Azure et n’a plus besoin de ce mécanisme. La dernière fois que j’ai regardé, ce n’était pas le cas d’Android. Vous pourriez prendre la décision d’imposer l’utilisation de l’app Outlook (les gens s’y font si vous leur laissez un moyen d’accéder à leur agenda sans Outlook).

# Engagez-vous qu’ils disaient !

## Dieu le veut !

Il s’agit d’activer le processus d’enrôlement depuis l’interface d’admin. (Si quelqu’un a réussi à le déclencher en Powershell sans effet de bord, on est preneurs !). Les utilisateurs, qui n’auront pas lu votre communication géniale, vont se prendre la procédure dès le démarrage d’Outlook. Les rares (10-15%) qui savent lire le contenu de la fenêtre vont se débrouiller tout seuls. Les autres vont submerger le sevice desk. Il est donc essentiel de lottir votre déploiement.

## À la main de l’utilisateur

Il s’agit d’envoyer l’utilisateur sur https://aka.ms/setupMFA . Si vos utilisateurs ne jettent pas ça à poubelle comme lien dangereux, ils pourront choisir le moment de leur enrôlement. Cela pose un problème de gestion de la relance. Il faudra donc indiquer un ultimatum auquel vous activerez de fait l’enrôlement. Nous vous conseillons cette méthode.

# Le temps des regrets

On aurait aimé proposer aux utilisateurs une période de transition, disons d’un mois, durant laquelle l’utilisateur aurait été sollicité à chaque connexion pour s’enrôler tout en gardant la faculté de reporter l’action.

Il est nécessaire de jouer avec l’ADFS ou l’accès conditionnel pour restreindre l’enrôlement (*proof up*) au réseau interne de l’entreprise. Souvent, les solutions déclenchent l’inscription à la prochaine connexion, quelle qu’en soit l’origine. Par conséquent, si vous n’activez le MFA que depuis l’extérieur, il faudrait attendre que l’utilisateur se connecte une première fois dans une situation de nomadisme pour qu’il s’enrôle. Autant dire qu’un attaquant qui arrive entre temps enregistrera le MFA sur son device (ce qui permettra de récupérer après coup son device ID et de tenter de le poursuivre, mais ce sera trop tard).

Il existe des personnes surmotivées qui, en survolant la communication globale, veulent activer tout de suite la protection. Plutôt que leur répondre d’apprendre à lire et d’attendre leur tour, il est préférable de leur offrir l’expérience en bêta, tout en veillant aux conséquences pour le service desk (s’assurer que les éventuels problèmes sont immédiatement remontés aux niveaux deux ou trois).

Les applications mobiles ne fonctionnent pas, ou mal, sur les anciens OS. Si vos téléphones ont trois ans d’âge, veillez à bien tester !

De manière générale, testez bien tous vos cas possibles sur votre périmètre, il y a forcément quelque chose qui va aller de travers.
