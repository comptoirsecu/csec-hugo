---
title: "DFIR SANS Summit 2020"
authors:
  - jil
date: 2020-07-17
image:  /images/covers/2020-07-17-sans-dfir-2020.jpg
categories:
  - Article
tags:
  - forensic
  - conférence
---

Me voici sorti de deux jours de conférences en ligne sur le thème de l’inforensique et de la réponse à incident. Le SANS DFIR Summit s’est tenu les 16 et 17 juillet 2020, en partie sur Zoom et en débordement sur YouTube. J’avoue avoir vite basculé sur YT, pour la flexibilité quant à la taille des fenêtres. L’interaction entre les participants avait lieu sur Discord, et les présentations étaient en direct. Plus de dix mille personnes en ligne, deux mille sur Discord.

J’ai pu suivre la majorité des deux tracks simultanément, ajustant le son selon l’intérêt de le présentation. Sans plus attendre, mes notes en vrac :

Les vidéos devaient être publiées sous deux semaines. Les supports ne sont téléchargeables qu’avec un compte SANS (gratuit, mais lisez les conditions quand même), comme d’habitude. Vous avez aussi le [fond d’écran de 24,9 Mio.](https://www.sans.org/images/misc/Zoom-Background-DFIR.png)

# COVID, Black Lives Matter, Community Matters, Common Good

La conférence est diffusée depuis les États-Unis d’Amérique, bousculés par la pandémie et les mouvements sociaux. Les keynotes parlent de richesse culturelle, de la relative rareté des analystes forensics et de la nécessité d’apporter ces compétences où l’argent n’afflue pas : les ONG, les prévenus sans le sou pour leur défense, les écoles, etc. Lee Whitfield, dans un talk dédié aux conséquences de l’absence d’analyste forensic sur la liberté des prévenus (Track 2), basé sur des cas réels et son expérience, revient sur la tentation de laisser passer l’émoi du sentiment d’injustice et de revenir à sa vie. L’appel à la contribution volontaire pour une meilleure société est lancé. Pour protéger des vies.

Il est question aussi de renforcer les méthodes scientifiques dans la pratique. De contribuer davantage au partage des connaissances en structurant les publications. De limiter les biais.

Et sinon, vous ai-je dit qu’il fallait contribuer à la communauté ? 

Le SANS Institute lance son premier camp d’été pour les ados, filles comme garçons. L’[agenda][summer-camp] devrait nous inspirer en Europe également.

Mais vraiment, contribuez ! (oui, ils ont insisté, insisté, insisté)

# Jour 1

## You Need a PROcess to Check Your Running Processes and Modules.

*Écouté en fond*

De la nécessité de surveiller les processes et leur usage de la mémoire est de plus en plus nécessaire pour détecter les activités malveillantes. J’ai survolé la solution et les hints pour sysmon, mais je pensais à Comæ.

## Kansa for Enterprise Scale Threat Hunting

Magnifique, extraordinaire, splendide travail sur Kansa pour collecter les artefacts sur plusieurs dizaines de milliers d’endpoints. Avec la gestion de la contention et des secrets temporaires, et de la suppression du bruit généré en s’appuyant sur le SOAR. Très riche en contenu. Par contre, une partie ne semble pas déployable lors d’une intervention à l’emporte-pièce (le cas : « au secours ! Des logs ? Quels logs ? Un SIEM, c’est quoi ? »).

Une bonne partie a été fusionnée sur le repo de Kansa. Une autre partie (dont la GUI pour éviter les lignes de commandes trop longues) n’est pas encore disponible.

## Data Science for DFIR – The Force Awakens

*Écouté en fond*

Application des techniques de big data pour visualiser différemment les logs et détecter visuellement les points où enquêter. Puis utilisation du machine-learning pour prétraiter les logs et sortir les écarts. Une bonne introduction à la manière dont ces techniques vont changer le métier des analystes, qui sont encore restés : « à l’ère d’Altavista, avant que Google débarque avec son *ranking algorithm* ».

## Making Memories: Using Memory Analysis for Faster Response to User Investigations

*Écouté en fond*

Excellent parcours des artéfacts clés sur Windows pour retracer l’activité utilisateur. J’ai eu l’impression de survoler une bonne partie des sections de la formation FOR500.

## Using Big DFIR Data in Autopsy and Other Tools

*Écouté en fond*

Présentation des évolutions d’Autopsy, de la conviction que nous sommes encore à l’âge de pierre pour rendre les analyses accessibles. Dévoilement de leur nouvel outil DFIR [Cyber-Triage][cybertriage] qui a l’air vraiment accessible et pourrait permettre un premier triage à quelqu’un qui ne serait que sommairement formé à l’inforensique. 


## Healthy Android Exams: Timelining Digital Wellbeing Data

Plongée dans les artéfacts des applications de bien-être de Google, et un peu de Samsung. On y trouve des trucs sympas, mais certains, comme l’historique de navigation, sont bloqués jusqu’au consentement de l’utilisateur. La rétention est limitée (7 jours, à vérifier). On peut tracer l’activité du terminal (écran déverrouillé) ou encore la gestion de la playlist. Il me semble avoir vu passer une liste d’applications supprimées. Complet, parfait.

## Just Forensics, Mercifully

Cf. l’intro. Il avait peur que les keynotes enlèvent le caractère particulier de sa présentation (il aime bien faire des présentations non techniques de temps à autre). Visiblement un grand moment émotionnel. À la suite de ce talk, le fondateur d’Autopsy a décidé d’offrir sa formation en ligne aux *défenseurs civils* comme il le fait pour les forces de l’ordre.

## Capa

Par des reversers de FireEye/Mandiant. Outil d’analyse statique d’un binaire pour sortir un résultat humainement lisible des capacités utilisées (par ex : télécharger depuis une URL). Permet à un analyste de voir si ça colle ou cloche avec son investigation sans avoir à appeler d’office le reverser. L’outil fonctionne mieux que prévu, sauf contre les binaires qui découpent une fonctionnalité dans plusieurs fonctions, si j’en crois les discussions sur Discord.

Un grand appel à la communauté pour enrichir les modules d’analyse. [Blog post ici][capa].

# Jour 2, main Track

## Help! We Need an Adult! Engaging an External IR Team

Par Lizz Waddell, Incident Commander de Talos. Inclure le « quand appeler de l’aide » dans le plan de réponse à incident et les raisons pour lesquelles les appeler (voir slides). Faites de préférence vos courses avant pour préparer le jour où ils vont débarquer et savoir combien ça va vous coûter. Définir la portée de leur engagement et leur objectif (root cause, lateral movement, data exfiltration, etc). En cas de déplacement sur site, quel site ? Quel niveau de vérification avant d’autoriser l’accès ? Préparer aussi la relation avec l’assureur (encore faut-il savoir que la société est assurée, hein, les RSSI…). Les logs, c’est bien, mais nom d’une pipe ! En UTC, vos logs ! Ça vous coûtera très cher sinon (car tout prendra plus de temps, et des erreurs seront commises par les analystes). 

Cette présentation est un trésor pour préparer la réponse à incident. Parmi les discussions Discord, un rappel du plan pour éviter que toute l’équipe IR soit off à cause du COVID, savoir ne pas s’engager quand on n’a pas ce qu’il faut pour travailler. Bonne discussion sur les playbooks également.

## Forensic Analysis of the Apple HomePod and the Apple HomeKit Environment 

*Écouté en fond*

Un très bon déroulement du désossage du HomePod pour savoir ce qui y est stocké et où, incluant les interactions avec les autres équipements connectés.

## Hunting Bad Guys that Use TOR in Real Time

Retraçant une attaque dont le C2 utilisait Tor, il trébuche sur de la compromission de busybox et met un honeypot en place. Sur un an, 74% de SQL Injection. En seconde place, les RCE. Je n’ai pas capté le rapport avec Tor, ça avait l’air d’être basé sur les exit nodes (j’ai du mal à ne pas décrocher sur de l’indian-english, et il semblait y avoir un souci de bande passante avec son canal audio). Une chose intéressante : quand les beacons cessent d’être émis, c’est que le rançongiciel est en train de chiffrer ses cibles. 

## Using Storytelling to Be Heard and Remembered

Rien de plus que le titre, c’est dommage. Il y avait des exemples sur l’utilité de raconter des histoires pour devenir mémorable, mais j’aurais aimé une histoire de forensics. 

## From Threat Research to Organizational Threat Detection

Modélisation de la menace avec les frameworks : Microsoft STRIDE or PASTA. C’est fantastique, mais on en vient vite à la réalité : on prend un schéma d’architecture pour modéliser le système susceptible d’être attaqué et on liste ses frontières comme les TTPs susceptibles d’être utilisées contre lui. 

## DFIR To Go

Discussion avec les organisateurs sur le déroulement de l’évènement (magnifiquement géré, il faut le dire).

## Cyber Sleuth: Education and Immersion for the Next Generation

Démarche d’éveil des adolescents aux métiers du DFIR. Ça fonctionne bien, et c’est suffisamment tôt pour avoir aussi les filles. En jouant un peu, on arrive à introduire les maths, la prise de notes et l’écriture des rapports, ce qui permet d’entrer dans le cursus. Le pilote est achevé et la démarche va s’étendre pour tenter de couvrir l’intégralité des États-Unis. Par contre, il faut des profs jeunes, donc ils vont tenter de faire ça avec des étudiants de master (enfin, un truc comme ça).

## The DFIRlympics

C’est l’heure du défi des DFIRlympics ! 8 participants. 

Duels. Un mot est donné, ils doivent rédiger quatre vers qui font sens en DFIR, en rimes.

*C’est là que ça commence à merder dans les présentations et l’utilisation de Zoom :D*

## Forensics Awards

Célébration avec animation virtuelle du bonhomme sur la scène. Sacré délai entre Discord et YT. 

Éric Zimmerman et son KAPE a été détrôné par Autopsy !

Double awards pour David Cowen !

Quadruple awards pour Heather Mahalik également, qui nous a mis l’ambiance pendant les deux jours en lisant les blagues (pourraves) des participants sur Discord ! C’est une première dans l’histoire des DFIR Awards.

Nouvel award pour les ressources. Le digital forensic Discord server prend la palme ! Accès réservé aux professionnels de la matière.


# Jour 2, Solutions Track

## Putting Big Data to Work in DFIR

*Zappé, la conf de Talos ne se prêtait pas à une double écoute.*

## How Not to Ruin Your Day: Avoiding Common Threat Hunting Mistakes

Les erreurs courantes des Threat Hunters : tenter de tout regarder, s’appuyer que sur des IOCs, ne pas prêter attention au contexte, rester dans sa zone de confort, ignorer les fichiers connus et signés, ne pas actualiser ses connaissances sur les dernières attaques. Les (un peu) moins évidentes : ne chercher que des APTs, ou chercher exactement la même chose que dans les rapports, ne pas poser des questions aux gens qui exploitent le SI, ne regarder que les fichiers au détriment de la mémoire, s’entêter dans une direction au lieu de bifurquer vers un autre scénario. Ne pas croire que les analystes ont vu ce qui s’est passé, donc toujours donner du contexte.

Puis la semi-automatisation des tâches avec le produit.

## Profiling Threat Actors in DNS

*Écouté en fond*

Un tour des fonctions de DomainTools.

## Completing the Triad, The Case For Leading With NDR 

*Écouté en fond*

NDR utile pour la visibilité sur l’IoT, sur les endpoints où l’attaquant a désactivé les logs (mais si tu as un bon heartbeat, tu devrais le voir), et pour éviter certaines contentions en IO des SIEM (pas trop compris pourquoi, mais j’écoutais peu). Utile aussi en réponse à incident, quand l’EDR a été contourné et qu’on ne sait pas qui est infecté (cas Ryuk). Un peu d’inspection peut faire le boulot (à condition d’avoir les IOCs…)

## Empowering DFIR Through Automation and Orchestration – Enhancing Your Artifacts with Threat Intelligence 

*Écouté en fond, j’ai décroché parce que sur la track 1, on racontait des blagues :’)* 

## Accelerate Your Threat Hunting and IR with Next-Gen NDR+EDR 

Démo de Defender ATP et BlueHexagon pour la détection de nouvelles menaces. Pas suivi plus que ça.

## Dig Deeper: Acquisition and Analysis of AWS Cloud Data

Démo de Magnet pour AWS et Azure, avec une intro sur les services de base AWS.

# Conclusion

Sans qu’on ait eu le plaisir des évènements SANS dans un lieu unique, cela a été l’occasion de réunir beaucoup de monde, parmi lesquels de nombreuses personnes qui n’auraient pas pu faire le voyage ni payer l’entrée. Les prochains summit, pandémie oblige, seront probablement virtuels, jusqu’à pouvoir revenir en vrai, et peut-être conservera-t-on les serveurs Discord éphémères.

[summer-camp]: https://www.sans.org/cyber-camp
[cybertriage]: https://cybertriage.com
[capa]: https://www.fireeye.com/blog/threat-research/2020/07/capa-automatically-identify-malware-capabilities.html
