---
title: "Faut-il tout journaliser ?"
authors:
  - jil
date: 2018-07-20
image:  /images/covers/2018-06-29-banniere-mailbox.jpg
categories:
  - Article
tags:
  - log
  - RGPD
  - GDPR
---

Quiconque s’est frotté à la réponse à un incident de sécurité sait l’importance des logs. Les journaux constituent la quasi-seule source d’informations sur ce qui s’est passé dans le système possiblement compromis. De leur complétude dépend la capacité à comprendre l’évènement et à définir la réponse la plus judicieuse à y apporter.

À l’heure où l’obligation de sécurité des données personnelles pèse en millions d’euros, et où la communication d’un incident portant sur elles doit intervenir dans les 72 heures ; à cette même heure où la collecte des journaux, qui sont eux-mêmes des données personnelles, doit rester proportionnée, que devrait-on faire ?

Ce billet n’est que le fruit de ma réflexion, il n’a pas de source juridique ; il vise simplement à étayer le questionnement de mes pairs. Il se concentre sur la protection du SI d’entreprise, pas des applications installées chez le consommateur.


# Acceptation sociale

Le contrat de travail, en droit français, subordonne le salarié à l’employeur. Cette subordination le soumet aux règles de l’entreprise, qu’il s’agisse du règlement intérieur ou des processus. En acceptant le contrat, le salarié accepte les contraintes du poste qu’il occupe ; on pourrait dire qu’il renonce à une partie de sa liberté. Dès lors, je pense que l’on peut dire qu’il consent à la collecte d’informations réalisée dans l’exercice de sa mission tant que cette collecte est loyale, proportionnée et qu’il en est dûment informé.

Il y a dix ans, dire à quiconque travaillait dans une PME que l’ensemble de ses actions effectuées avec l’outil informatique serait tracé aurait sans dout paru disproportionné. Aujourd’hui, en l’état de la menace et d’une définition réglementaire de l’obligation de sécurité de plus en plus précise (juridiquement parlant ; techniquement, ça reste flou), je crois que cela peut s’entendre.

Il n’en reste pas moins que le législateur reconnaît un droit d’usage à des fin personnelles de l’équipement fourni par l’entreprise. De la même manière, les ordiphones sont par nature à double usage. La protection de la vie privée demeure un droit fondamental. Toute atteinte qui lui est portée doit être justifiée. Heureusement, les entreprises bénéficient de la présomption d’usage professionnel du matériel qu’elle met à disposition. Il reste à savoir si cette présomption sera étendue au BYOD, qui consiste à intégrer un matériel personnel au contexte professionnel.


# Finalité

Avant même d’envisager l’ampleur de la collecte, il convient d’en définir la finalité. Cette question n’est pas sans poser problème. Il peut y avoir trois grandes finalités :

> la détection, l’analyse et le traitement des évènements susceptibles
> de porter atteinte à l’intégrité, à la disponibilité ou à la
> confidentialité des données de l’entreprise, dans le but de préserver
> ces données et de garantir la bonne marche de l’entreprise.

Cette définition, que je sors de mon esprit, me paraît couvrir les activités de sécurité dans leur dimension technique : tracer ce qui se passe pour comprendre s’il y a un problème et, le cas échéant, y remédier.

Or, dans le remède, on peut trouver, outre les mesures techniques, des mesures disciplinaires ou juridiques. Il y a donc une autre finalité au traitement, beaucoup plus sujette à débat quand on entre dans la question de la durée :

> l’engagement de procédures disciplinaires, contractuelles ou légales
> en cas de manquement par un salarié ou un sous-traitant au respect de
> la politique de sécurité de l’information ou à ses engagements
> contractuels.

Je vous laisse la relire plusieurs fois, on ne peut pas y mettre de virgules pour qu’elle ait toute sa portée juridique. Il s’agit là de pouvoir utiliser les journaux pour constater la faute. Il faudra toutefois mettre en œuvre les moyens techniques permettant de constituer une preuve légale (signature, horodatage, séquestre) si les journaux sont le seul fondement à la constitution de la faute.

Cette finalité permet de couvrir aussi bien le manquement à une clause de secret qu’à une obligation de télémaintenance (si votre contrat stipule une maintenance préventive mensuelle que vous n’avez aucune trace de connexion, selon la précision de votre contrat, vous avez de quoi agir).

Enfin, vous avez la finalité imbattable :

> l’exécution des obligation de collecte imposées par voie réglementaire

C’est l’occasion de relire la Loi pour la confiance dans l’économie numérique, le Code des postes et des télécommunications, les différentes lois de programmation militaire et autres joyeusetés de même acabit.


# Durée

Il y aura probablement une durée différente par finalité. Ainsi, les actions disciplinaires sont prescrites après deux mois (de mémoire), les journaux sont souvent exigés pendant 6 à 12 mois, et pour le traitement de vos incidents, vous pourriez être intéressés par 2 à 3 ans de rétention (partant du principe que vous découvrirez l’incident un an et demi après sa survenue). 

La clarification des durées est néanmoins fondamentale pour garantir la sérénité des personnes vis-à-vis de la collecte.


# Étendue

L’état de la technique en la matière a de quoi faire peur. Les métadata des communications suffisent à connaître avec précision la vie l’un individu et détecter les changements de comportements. [retrouver la vidéo du documentaire]

## Périmètre de surveillance

Pour que la collecte soit proportionnée, elle ne peut s’effectuer que sur les composants qui jouent un rôle dans les données à protéger. Il sera difficile de conserver le journal d’utilisation de la cafetière.

## Journaux d’accès

L’ensemble des connexions et tentatives de connexion, couronnées de succès ou non, à chaque système entrant dans le périmètre de surveillance. Pour chaque connexion : les identifiants utilisés, l’adresse d’origine et de destination, les caractéristiques connues ou devinables, potentiellement obtenue par croisement auprès de sources externes, et qui permettent d’idenfier de manière unique le terminal source.

Il faut bien stocker le user agent, l’adresse mac, le fingerprint d’un bon nmap des familles autoexécuté sur l’appelant (ah, proportionné, a-t-on dit ?) Cette rubrique me semble peu sujette à débat (sauf sur le croisement, éventuellement).

## Comportement d’un système

L’ensemble des caractéristiques techniques permettant de situer l’état d’un système à un instant donné et l’utilisation qui en est faite. Dont : l’activité des processus, le détail des connexions aux réseaux (depuis le système ou de l’extérieur vers le système), l’attribution des privilèges, l’état de la configuration, l’altération d’une configuration, le détail des fichiers consultés et leur altération, etc.

Qui a dit *endpoint protection* ? Ça couvre aussi auditd, sysmon, les event logs, les journaux des pare-feux… Tout ce qui permet, après coup, de comprendre comment l’attaque est passée.

Ici, la question est de savoir, pour les EDR en praticulier, si tout le parc doit être concerné ou seulement une population à risque ? Je crois que la réponse dépend de votre schéma de défense (capacité de déplacement latéral sur le parc des ordinateurs). Imaginons que vous soyez confiant dans votre segmentation, qu’un poste ne peut pas parler à son voisin et que vous avez, du côté des serveurs applicatifs, toute la traçabilité système. Vous devriez être en mesure de déterminer avec précision les systèmes impactés et la machine utilisée en rebond de l’attaque. Est-il alors nécessaire (donc, proportionné) d’avoir le détail d’activité de cette machine en temps réel ?

Autrement dit : l’immixtion dans l’activité de l’utilisateur lambda, avec le risque de surveiller ce qui relève de son usage privé de la machine, est-elle proportionnée par rapport à ce que vous pourriez obtenir en analysant vos serveurs de messagerie (pour confirmer la réception d’un courriel malveillant) puis en réalisant une analyse post-mortem de ladit machine ? Donc avec un bon processus de séquestre et le risque d’avoir des eventlogs effacés.

La réponse du professionnel de la sécurité, c’est « évidemment ! » comme ça, je peux expliquer ce qui s’est passé. Mais dans un monde où le smartphone prévaut, quand les EDR seront déployés sur votre téléphone, avec la géolocalisation et la trace détaillée de votre usage privé du terminal, trouverez-vous toujours cette collecte proportionnée ? N’y a-t-il pas un risque que, pour en instruction pénale extérieure à l’activité de l’entreprise, ces données soient réquisitionnées et que vous trahissiez ainsi un sujet loyal ?

Renseignez-vous sur le système de score social chinois et reposez-vous la question.

