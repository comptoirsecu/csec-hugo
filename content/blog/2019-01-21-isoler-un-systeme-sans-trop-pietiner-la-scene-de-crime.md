---
title: "Kit de survie : isoler un système sans trop piétiner la scène du crime"
authors:
  - jil
date: 2019-01-21
image:  /images/covers/2019-01-14-banniere-forensics.jpg
categories:
  - Article
tags:
  - incident
  - forensic
---

> Vas-y Robert ! Débranche la prise !

Les gestes qui sauvent en cas de contamination d’un système ne sont plus ceux de naguère. Si vous n’avez pas de procédure de réaction à un incident de sécurité impliquant la compromission d’un poste de travail ou d’un serveur, ou que celle-ci ne comporte pas la collecte des éléments vitaux pour une analyste post-mortem, cet article est pour vous.

Je me concentre sur les systèmes Windows. Certains outils sont portables.

# Vous ne serez pas seul

Quand on a la classe (peut-être avec un étendard [PRIS, annexe 2][PRIS]), ou qu’on s’en est déjà pris de belles dans les dents, on sait qu’identifier un système contaminé et envisager son isolement marque le début d’une nouvelle histoire. Le temps où il suffisait de débrancher la prise et de reformater l’ordinateur incriminé est révolu. L’époque est au rançongiciel à déplacement latéral automatisé, à l’intrusion pilotée par un serveur *Command & Control (C2/C&C)*, au maliciel qui réside uniquement en mémoire. Vous aurez donc besoin d’analystes, et les analystes travaillent sur des artéfacts.

# Pokéball, go !

Récupérer les artéfacts, c’est un peu comme réussir à capturer un pokémon rare du temps des premiers jeux vidéo Pokémon (c’est-à-dire l’époque où l’on attendait du joueur un tant soit peu de talent, qu’il fallait affaiblir le Pokémon, mais pas trop, sinon on obtenait un cadavre, ni trop peu, car il te renvoyait ta pokéball dans les dents) : il faut capturer les bonnes choses au bon moment. Voyons de suite les fausses bonnes idées :

## Couper l’alimentation électrique

Très pratique en cas d’incendie, mais très dangereux à l’époque des disques chiffrés et des maliciels qui hantent la mémoire vive. Vous allez perdre des traces.

## Procéder à une copie intégrale du disque dur 

Parfois nécessaire en cas de procédure judiciaire, mais souvent inutile et trop lent pour réagir à une attaque.

## Commencer par déconnecter le réseau

C’est un peu comme déterminer le côté duquel tombera une tartine. Ce sera de plus en plus une mauvaise idée, car le maliciel qui détecte une perte subite du réseau pourrait décider de saborder la mission. Bon courage pour le retrouver s’il disparaît de la mémoire.

Cela ne veut pas pour autant dire qu’il faudrait laisser toutes les connexions au réseau actives. Simplement, il est préférable de commencer par d’autres étapes. De toute façon, le temps que vous identifiiez la bonne machine, vous n’êtes probablement plus à ça près. Si vous faites face à un péril imminent en provenance d’une machine, utilisez le pare-feu qui se trouve devant vos ressources à protéger pour ne filtrer que les flux qui vous menacent. L’idée est de ne surtout pas rompre la communication entre le maliciel et son C2. (Et là, j’entends déjà les hurlements, alors venez arguer sur [Discord][disord], et faites votre propre analyse de la menace).

# *Same player, shoot again*

Quoi qu’on fait, alors ?


## Prérequis de longue date 

Tout d’abord, il vous faut un périphérique USB d’une taille démentielle : RAM + pagefile + 3 Go. Dans les faits, vous ferez avec ce que l’assistant du coin aura sous la main. Au pire, vous opérerez depuis un partage de fichiers en réseau.

Vous aurez un partage de fichiers que vous pourrez facilement rendre accessible de toute part pour que votre assistant improvisé puisse préparer sa clé USB sans réfléchir (un bon vieux copier/coller de tous les fichiers après avoir formaté la clé). Il y aura là vos outils de collecte de preuve. 

En outre, vous vous serez assuré que vos antivirus n’y voient pas d’inconvénient.

Pour le cas le plus simple (et qui serait déjà un progrès), nous allons mettre deux outils. C’est un choix tout à fait arbitraire, vénéré par d’aucuns du monde de l’au-delà (pardon, de l’analyse post-mortem), mais qui minimise l’impact sur la cible :

* [dumpIt.exe][dumpIt] pour récupérer la RAM
* [CyLR.exe][cyLR] pour collecter les artéfacts


## Prérequis *in situ*

Vous avez besoin du compte administrateur local. Pour ceux pour qui il serait opportun de le rappeler, on n’utilise jamais un compte disposant de privilèges globaux sur un système suspecté d’avoir été compromis. Changer un mot de passe ne révoque pas les tickets Kerberos.

Profitez-en pour récupérer la clé de déchiffrement du disque si vous l’avez en central.

## Action !

1. Ouvrir votre chrono des opérations et noter les étapes au fur et à mesure, en mettant tous les commentaires qui traduisent ce qui se passe (du genre : pas de clé USB dispo, on va chercher les outils via `\\serveurTruc\muche` ; cela évitera d’égarer l’analyste sur de fausses pistes)
1. Connecter la clé (à défaut, ouvrir le partage réseau)
1. Exécuter DumpIt, qui va récolter la RAM. On commence par lui pour minimiser les changements qui interviennent sur le système dont on fait l’image.
1. Si vous avez un IT sous la main : ouvrir une invite de commande en tant qu’administrateur. Sinon, faites une prise de main à distance, on a déjà la RAM (mais gare à ne pas divulguer vos identifiants, [`mstsc /restrictedAdmin`][restrictedadmin])
1. Exécuter CyLR. Il va copier tous les fichiers indiqués dans sa configuration (registre, profil utilisateur, journaux d’évènement, etc.)
1. Éjecter proprement la clé
1. Déconnecter le système du réseau 

## Et maintenant ?

Et maintenant, vous avez une clé USB que vous n’osez pas connecter ailleurs. Dans le meilleur des mondes, vous avez un assainisseur tel que le [CIRCLean - USB key sanitizer][CIRCLean], ce qui vous permettra de transférer les données sur un autre disque avant de les transférer pour analyse. 

Dans la vraie vie, vous connectez la machine à un wifi qui dispose juste d’un accès internet et vous transférez tout sur un espace en ligne partagé, de préférence qui n’exige pas d’identifiants ou des identifiants dédiés à cette capture-ci (et, dans les deux cas, révoqués incontinent après le transfert).

Dès réception, vos analystes vont vérifier les condensats et se mettre au travail. Gardez la machine cible sous la main, car ils pourraient revenir vers vous. Par exemple, si le maliciel a vidé les journaux d’évènements, ils vous demanderont de regarder ce qui se trouve dans les *volume shadow copy*, ne sait-on jamais. 

Les artéfacts de Windows pourront quelques fois (de rares fois ?) permettre de corroborer la méthode d’infection. C’est assez probable s’il y a eu interaction avec l’utilisateur, et bien moins s’il n’y a que des appels système. D’où l’absolue nécessité de collecter la mémoire pendant que le maliciel est actif. C’est en la désossant qu’un analyste pourrait trouver des choses utiles.

Vous pouvez également augmenter vos chances de succès en établissant une [politique de journalisation][eventlog].

Si le succès de l’analyse n’est pas garanti, vous aurez au moins mis toutes les chances de votre côté pour comprendre ce qui se tramait réellement, voire vous faire une idée de qui vous en veut, et peut-être même ferez vous profiter la communauté de nouveaux IOCs.

[PRIS]: https://www.ssi.gouv.fr/actualite/le-nouveau-referentiel-pris-pour-les-prestataires-de-reponse-aux-incidents-de-securite-est-maintenant-disponible/
[discord]: http://discord.comptoirsecu.fr 
[restrictedadmin]: https://social.technet.microsoft.com/wiki/contents/articles/32905.remote-desktop-services-enable-restricted-admin-mode.aspx
[circlean]: https://www.circl.lu/projects/CIRCLean/
[cyLR]: https://github.com/orlikoski/CyLR/releases
[dumpIt]: https://blog.comae.io/your-favorite-memory-toolkit-is-back-f97072d33d5c
[eventlog]: https://acsc.gov.au/publications/protect/windows-event-logging-technical-guidance.htm
