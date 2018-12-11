---
title: "Retour d’expérience sur une politique de mot de passe zxcvbn en entreprise"
authors:
  - jil
date: 2018-12-13
image:  /images/covers/2018-12-11-rex-zxcvbn-rajeev-mog.jpg
categories:
  - Article
tags:
  - password
  - zxcvbn
  - rex
---

Avec la dernière [recommandation du NIST][nist], nous avions enfin une raison légitime de benner les anciennes politiques de mot de passe avec renouvellement périodique. Encore fallait-il oser le faire (et vérifier que son cadre réglementaire le permettait). J’ai tenté l’expérience, voici mes conclusions.

Avant de commencer, il est utile de préciser que le changement de mot de passe s’effectue exclusivement sur une page web.

# Situation

À l’origine, plus de 70 % des mots de passe de l’AD succombaient à une attaque. À la fin, en ciblant uniquement les mots de passe ayant appliqué la nouvelle politique (il y a toujours des comptes difficiles à mettre à jour dans un AD sans écrouler un pan du SI…), le taux tombe à 0,02 % (9 comptes sur 4 000) avec les derniers dictionnaires disponibles par la société qui avait fait le premier cassage. Si l’on peut légitimement arguer que le résultat dépend de la qualité des dictionnaires de l’attaquant, le gain est alléchant.

# Contexte utilisateur et plan de communication

L’expérience a été conduite dans une entreprise internationale où les sujets ne sont pas particulièrement sensibilisés à la protection du patrimoine informationnel. Ils savent simplement que l’on entend de plus en plus parler dans la presse d’attaques informatiques, et qu’il est cohérent que l’entreprise y fasse quelque chose. Plus de la moitié des sujets ont connu dans l’année un traumatisme : le déploiement de l’authentification à deux facteurs dans le cadre de la réaction à une intrusion. À la fin de cette crise, ils avaient été informés que la politique de mot de passe évoluerait quelques mois plus tard.

En septembre, la nouvelle politique de mot de passe est communiquée par courriel officiel et par le réseau social de l’entreprise. Le courriel est bref et renvoie à une page dédiée pour de plus amples informations. Deux points sont particulièrement mis en avant dans cette deuxième communication (la première étant celle de fin de crise) : 

* sauf perte ou compromission, ce sera le dernier changement de mot de passe ; 
* trouver un mot de passe conforme s’annonce être un véritable défi.

Le courriel pointe vers une page web interne qui permet aux sujets de tester leur prochain mot de passe. Il leur est vivement conseillé de chercher un mot de passe conforme avant l’entrée en vigueur de la nouvelle politique.

Le jour dit (un mois plus tard) une nouvelle communication officielle annonce l’entrée en vigueur et exige le changement de tous les mots de passes sous 30 jours.

À compter de ce moment, les agents de l’assistance informatique closent chacun des appels (quel que soit son motif) par un rappel à l’utilisateur de changer son mot de passe si ce n’est pas déjà fait, et l’assistent si nécessaire.

À l’occasion d’une prise de parole du DG, ce denier a glissé quelques mots sur la nécessité de changer son mot de passe, ce qui a poussé une nouvelle vague de sujets à l’acte.

À J+25, un rappel est fait aux retardataires. Ce sera ensuite un rappel quotidien, puis un encouragement fort avec la coupure de l’accès à l’internet (sauf pour quelques ressources en liste blanche).

À ce plan théorique se sont ajoutées plusieurs modifications de la page de changement de mot de passe pour clarifier l’accès à la page d’aide (avec des conseils pour la composition du mot de passe) et pour alerter de certaines incompatibilités avec des applications obsolètes.

# Paramétrage

L’approche est fondée sur [zxcvbn][zxcvbn] de Dropbox, qui a été complété avec des dictionnaires qui couvrent les langues principales en usage dans l’entreprise et quelques mots en liste noire liée au contexte de l’entreprise (nom des logiciels internes, des marques et des éléments qui ressortaient dans le dernier cassage de mot de passe). Le script qui en résulte est atroce, pesant plus d’un mégaoctet, mais le réseau le permet.

zxcvbn découpe le mot de passe qui lui est soumis et calcule le temps théorique requis pour un cassage par dictionnaire et force brute. Il est configuré pour n’accepter que les mots de passe avec un log à 15 (voir la page de détail des calculs dans la doc du script), donc censé résister à une semaine d’attaque à raison de 10 milliards de tentatives par seconde. La longueur minimale du mot de passe est de 12 caractères, mais personne n’arrive à la barre verte avec 12 caractères ; 14 semble le minimum viable.

# Réactions des utilisateurs

Certains se sont pris au jeu et un concours est né dans les *open spaces* pour trouver un mot de passe vert. Au fil des sollicitations, l’un des conseils majeurs, outre de trouver une phrase courte qui fait sourire, a été de puiser dans l’argot ou dans des langues rares. 

D’aucuns se sont plaints de la difficulté à trouver un mot de passe qui soit accepté, et parmi ceux-là, certains ont trouvé un bug qui permettait de valider des mots de passe plus faibles que la politique. Dans l’accompagnement du changement, j’ai décidé de ne pas corriger le bug et d’attendre le verdict de l’attaque. Bien qu’ils aient saisi un mot de passe plus faible, il s’est révélé incomparablement meilleur que le précédent. Le paramétrage indiqué plus haut est donc trop exigeant.

J’ai été surpris de ne pas être confronté à une fronde plus forte des sujets conduisant à une réduction des critères d’acceptation du mot de passe. Je m’y étais préparé, mais comme me l’a indiqué un utilisateur : « nous avions été prévenus, nous avons eu le temps de chercher, rien de bloqué le temps d’y parvenir, donc c’était imposable ».

Il m’a néanmoins été reproché de mieux protéger nos accès que les banques, et que cette robustesse ne devrait être exigée que des services les plus sensibles. La notion de maillon faible n’est comprise que des informaticiens qui savent ce qu’est l’Active Directory, et le sentiment de devoir protéger ses données pour protéger la compétitivité de l’entreprise, donc son emploi, relève d’une culture à développer. Il est intéressant de noter que l’on m’oppose régulièrement ces arguments lorsqu’il s’agit de déployer des mesures avant-gardistes et effectives, comme si une entreprise devait être innovante sur ses marchés et surtout pas sur la protection de ses informations. Pour couper court à la discussion, je rétorque usuellement que si les autres sont mal protégés, ce seront eux qui seront attaqués, et que nous poursuivrons une croissance tranquille (en moins diplomate : je fais mon job, je le fais bien, alors fais le tien).

# Limites

Je suis stupéfait par la longueur de certaines phrases de passe que les utilisateurs ont choisies. Même si j’ai surpris certaines personnes expliquer à leur collègue ronchon que ça ne leur prenait pas plus de temps que le mot de passe précédent (avec lequel il se demandait toujours où étaient la majuscule et le chiffre), cela conforte mon ambition de réduire le nombre de fois où ce mot de passe est exigé. C’est chose faite par le recours à Windows Hello sur Windows 10 (code pin et reconnaissance faciale ; je suis impatient de tester la connexion par *companion device*), et par l’authentification unifiée que nous exigeons depuis plusieurs années à l’intégration de toute nouvelle application (SAML, OAuth2, OpenID Connect).

Il est aussi important de noter que ces mots de passe longs ne sont pas adaptés à toutes les populations. Ainsi, le personnel des ateliers, où les opérateurs passent leur temps à le saisir, ne peut l’appliquer sans compromettre sa productivité et son humeur (ce dernier point faisant craindre une démission dans certaines zones géographiques à forte concurrence). D’autres mécanismes, tels que la restriction d’utilisation du compte à un périphérique donné, ou l’authentification à courte portée par bracelet, doivent être préférés.


# Conclusion

Toutes les expériences que nous menons dans le domaine de la sécurité, où notre travail consiste à trouver des parades aux attaques, ne sont pas couronnées de succès. Aussi, pour une fois que nous avons une avancée qui semble vraiment porter ses fruits, en attendant que le mot de passe disparaisse, profitons-en. Je n’ai détaillé là qu’une mesure, qui s’intègre dans un ensemble plus large, et qu’il ne faut pas perdre de vue pour ne pas exiger trop de choses sur la porte quand la fenêtre est grande ouverte.

[nist]: https://www.comptoirsecu.fr/blog/2017-08-27-le-nist-signe-t-il-la-mort-de-nos-politiques-de-mot-de-passe/
[zxcvbn]: https://github.com/dropbox/zxcvbn
