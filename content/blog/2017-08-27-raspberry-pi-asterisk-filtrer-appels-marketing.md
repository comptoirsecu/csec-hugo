---
title: "Filtrer les appels marketing avec un raspberry-pi et un peu de config Asterisk"
authors:
  - ptro46
date: 2017-08-27
image: /images/covers/2017-08-25-raspberry-pi-asterisk-filtrer-appels-marketing.jpg
categories:
  - Article
tags:
  - RaspberryPi
  - Raspbian
  - Asterisk
  - NoMarketing
---

# Le besoin


## Filter les appels des boites de Marketing

Tout le monde à recu des appels de service de marketing, parfois cela frise le harcellement, l'idée est d'utiliser un Asterisk pour essayer de filtrer ces appels indésirables. Bien sur il est possible de s'inscrire sur bloctel [*http://www.bloctel.gouv.fr*](http://www.bloctel.gouv.fr) lancé dans le cadre de la loi Hamon sur la consommation, mais celui ci est confié à un opérateur privé (Opposetel), le fait qu'un organisme privé puisse donc en toute légalité collecter les informations nominatives et se constituer ainsi un fichier de mise en relation avec des numéros vérifiés peut poser quelques soucis à certains.
Même si vous vous êtes inscrit sur bloctel celui ci ne filtre pas 100% des appels de services de tele-démarcharge puisque ces sociétés peuvent tout à fait refuser d'appliquer les filtres de bloctel, vous constatez donc que certains appels continuent à vous déranger.
Nous allons donc voir un premier temps sur quel principe s'appuyer pour les filtrer, une fois la méthode connue celle ci est vérifiable avec un simple téléphone quand vous recevez un de ces appels, nous verrons donc dans un second temps comment automatiser cette méthode en la confiant à un Asterisk, nous verrons donc comment installer Asterisk sur un Raspberry Pi 3.

# La méthode

La prochaine fois que vous recevez ce genre d'appel écoutez bien ce qui se passe au moment ou vous décrochez, vous allez constater que vous n'avez pas directement l'interlocuteur mais qu'en premier il y a une mise en relation, d'ou la petite attente et le cliquetis de la mise en relation.

## Vous n'etes pas directement appelé par un humain

C'est parce que vous n'etes pas directement appelé par un humain, ces sociétés optimisent le temps de leur personnel d'appel, il est donc hors de question qu'un opérateur soit mis en relation avec un répondeur, pour cela ils utilisent un robot d'appel qui va essayer de déterminer si la ligne est décrochée par un répondeur ou par un humain, pour cela il va analyser ce qui se passe au moment ou la ligne est décrochée.

***Si un humain décroche il y avoir une séquence qui ressemble à***

   -   décrochage de la ligne

   -   petite attente de l'ordre de 1 à 2 secondes

   -   l'humain va dire : Oui Alo ?

***Si un répondeur décroche il y avoir une séquence qui ressemble à***

   -   décrochage de la ligne

   -   Message du répondeur

Le rôle du robot d'appel est donc d'essayer d'identifier une séquence qui va lui permettre de décider si il est en présence d'un humain ou d'un répondeur, pour cela il va analyser ces quelques secondes qui suivent le décrochage, c'est a dire la petite attente entre le décrochage et la voix humaine, si il pense être en relation avec un humain il fera la mise en relation, sinon il va raccrocher et passer à son numéro suivant.

## Comment tromper le robot d'appel

Sur le prochain appel faites le test, vous décrochez et vous ne dites rien, il y a de fortes chances que dans les 3 secondes le robot vous raccroche au nez ! C'est simple non ? Nous allons donc maintenant voir comment automatiser cette fonction grâce à Asterisk, en l'améliorant un peu pour vraiment tromper le robot, pour cela nous allons faire jouer une séquence a Asterisk qui va vraiment dire au robot qu'il n'a pas un humain en face de lui, voici ce que nous allons lui faire faire.

***A l'instant ou l'appel entrant est détecté***

   -   décrochage de la ligne

   -   Renvoyer à l'appelant une sonnerie d'attente

   -   Attendre 5 secondes

   -   Faire sonner les postes SIP

Placons nous dans le cas d'un humain pour comprendre ce qui va se passer pour l'appelant, et dans le cas d'un robot d'appel.

### Si un humain appelle

L'humain ne va pas détecter le décrochage de la ligne puisque nous renvoyons tout de suite une sonnerie d'attente de 5 secondes, donc il a l'impression que la personne qu'il appelle met 5 secondes avant de décrocher. Comme au bout des 5 secondes nous faisons sonner les postes SIP si le poste est vraiment décroché la mise en relation est faite, sinon on renvoie sur un répondeur. Donc pour l'humain c'est une séquence tout a fait normal parce que son oreille ne capte pas le décrochage.

### Si un robot appelle

Le robot lui va détecter le décrochage, à cet instant il lance son analyse en attendant une voix humain, comme on lui renvoie le son d'une sonnerie il en conclue qu'il n'a pas un humain de l'autre coté, et il raccroche ! vous verrez que cela se joue en moins de 3 secondes. ***Attention à ce moment les postes SIP n'ont pas encore sonné, donc le robot aura bien raccroché AVANT que vos postes SIP sonnent, vous ne serez plus dérangé !***

# Le matériel

* Raspberry Pi Carte Mère 3 Model B Quad Core CPU 1.2 GHz 1 Go RAM ([*https://www.amazon.fr/dp/B01CD5VC92/*](https://www.amazon.fr/dp/B01CD5VC92/))

{{< img src="2017-08-25-raspberry-pi-3b.jpg" desc="Raspberry Pi 3 Model B" >}}



* Raspberry Pi 3 Étui – Noir/Gris ([*https://www.amazon.fr/dp/B01F1PSFY6/*](https://www.amazon.fr/dp/B01F1PSFY6/))

{{< img src="2017-08-25-raspberry-pi-boitier.jpg" desc="Raspberry Pi 3 Boitier" >}}



* Carte Mémoire microSDHC SanDisk Ultra 32GB (Nouvelle Version) Vitesse de Lecture Allant jusqu'à 80MB/S, Classe 10 FFP ([*https://www.amazon.fr/dp/B013UDL5RU/*](https://www.amazon.fr/dp/B013UDL5RU/))

{{< img src="2017-08-25-carte-sd.jpg" desc="Carte SD" >}}


* Aukru Micro USB 5v 3000mA Chargeur Adaptateur Alimentation Pour Raspberry Pi 3, Pi 2 modele b et Modele B+ (B Plus) ,Banana pi ([*https://www.amazon.fr/dp/B01566WOAG/*](https://www.amazon.fr/dp/B01566WOAG/))

{{< img src="2017-08-25-alimentation.jpg" desc="Alimentation" >}}

# Installation

## Debian Stretch

La distribution utilisée est la raspbian-pi que vous pouvez récupérer directement sur leur site, bien sur une fois téléchargé le zip on vérifie son sha256 puisqu'ils ont pris la peine de l'indiquer !

* [*https://www.raspberrypi.org/downloads/raspbian/*](https://www.raspberrypi.org/downloads/raspbian/)

Ensuite il suffit de suivre la documentation pour installer l'image sur la carte SD

* [*https://www.raspberrypi.org/documentation/installation/installing-images/README.md*](https://www.raspberrypi.org/documentation/installation/installing-images/README.md)

Sur Mac ou linux il suffit d'utiliser la commande dd apres avoir identifié la device du lecteur SD, voici un exemple dans lequel le lecteur SD est sur la device /dev/rdisk2

`sudo dd bs=1m if=2017-08-16-raspbian-stretch.img of=/dev/rdisk2 conv=sync`

Une fois l'opération terminée il suffit d'insérer la carte SD dans le raspberry pi, le brancher et c'est parti, il faudra mettre à jour par un petit upgrade afin d'être à jour et ensuite nous pouvons passer à l'installation de Asterisk et sa configuration.

## Configuration de ligne Free en SIP

Avant d'aller plus loin il est nécessaire de configurer sa ligne free afin qu'elle puisse être utilisée avec la VoIP, pour cela il faut aller dans la configuration de la freebox (sur le site de free) et activer le compte SIP.

{{< img src="2017-08-26-freephonie.jpg" desc="Activation SIP Free" >}}

Cela peut fonctionner aussi avec d'autres opérateur (aucun soucis avec une ligne de VoIP d'OVH par exemple), mais bien vérifier si c'est possible avant de se lancer, attention en ce qui concerne Free ils bloquent les appels sortant vers les mobiles, donc si vous voulez pouvoir continuer à appeler les mobiles il sera nécessaire d'ajouter un second compte SIP dans Asterisk qui sache le faire, un compte SIP OVH fera l'affaire par exemple.

## Asterisk

Aller il est temps de s'occuper de la configuration d'Asterisk, pour cela nous allons devoir modifier les fichiers suivants

   - ***sip.conf*** : Contient la configuration de connexion au serveur freephonie, ainsi que la celle des postes clients qui se connectent en SIP

   -   ***extensions.conf*** : Contient les informations sur comment traiter les appels entrant ou sortant, c'est celui ci qui va nous intéresser pour bloquer les appels marketing

   -   ***voicemail.conf*** : Contient les informations sur la boite mail pour l'envoie des messages du repondeur par mail comme fait Free.


### Fichier sip.conf

```
[general]
language=fr
context=default
allowoverlap=no                 ; Disable overlap dialing support. (Default is yes)
udpbindaddr=0.0.0.0             ; IP address to bind UDP listen socket to (0.0.0.0 binds to all)
tcpenable=no                    ; Enable server for incoming TCP connections (default is no)
tcpbindaddr=0.0.0.0             ; IP address for TCP server to bind to (0.0.0.0 binds to all interfaces)
transport=udp                   ; Set the default transports.  The order determines the primary default transport.
srvlookup=yes                   ; Enable DNS SRV lookups on outbound calls

defaultexpiry=1800
dtmfmode=auto
qualify=yes

register=>09XXXXXXXX:SUPERPASSWORD@freephonie.net/fromfree

[freephonie-out]
type=peer
host=freephonie.net
username=09XXXXXXXX
fromuser=09XXXXXXXX
secret=SUPERPASSWORD
qualify=yes
nat=force_rport,comedia
disallow=all
allow=alaw
musiconhold=freephonie
canreinvite=no
callcounter=yes
restrictcid=no
insecure=invite
amaflags=default
dtmfmode=auto
context=maison
```


#### Commentaires

Pas grand chose à ajouter, il suffit de reprendre ce fichier de configuration en faisant bien attention d'adapter les lignes qui suivent :

| Ligne | Modifications |
|---|---|
| register=>***09XXXXXXXX***:***SUPERPASSWORD***@freephonie.net/fromfree  | La il faut indiquer le numéro de votre ligne SIP et le mot de passe que vous avez parammétré dans l'étape (Configuration de ligne Free en SIP)  |
| username=***09XXXXXXXX***  | A nouveau le numéro de votre ligne SIP  |
| fromuser=***09XXXXXXXX*** | A nouveau le numéro de votre ligne SIP  |
| secret=***SUPERPASSWORD*** |A nouveau le mot de passe |




### fichier sip.conf suite

```
[001]
type=friend
username=001
secret=MDPPOSTESIP
musiconhold=freephonie
host=dynamic
context=maison
canreinvite=no
dtmfmode=auto
disallow=all
allow=alaw
allow=g722
allow=h263
mailbox=001@default
callcounter=yes
nat=force_rport,comedia
qualify=yes
```

#### Commentaires

C'est la configuration d'un poste SIP, vous pouvez en configurer autant que vous souhaitez, voici quelques points a surveiller


| Ligne | Détails |
|---|---|
| 001, 002 ... | N'essayez pas d'utiliser autre chose que des numéros pour cette configuration, cela risque de perdre Asterisk. |
| secret=***MDPPOSTESIP*** | Ce mot de passe sera a configurer dans le poste SIP |
| context=maison | Vous pouvez mettre ce que vous voulez comme context mais il faut faire attention à etre coherent à tous les endroits ou il est utilisé |


### Fichier extensions.conf

Dans le fichier extensions.conf il faut laisser tout ce qu'il a placé par defaut et insérer sa configuration à la fin

```
[maison]
exten=>fromfree,1,Answer()
exten=>fromfree,n,Ringing()
exten=>fromfree,n,Wait(10)
exten=>fromfree,n,Dial(SIP/001,30,rt)
exten=>fromfree,n,VoiceMail(001@default,su)


exten => _0[1-9].,1,Dial(SIP/freephonie-out/${EXTEN})


exten=>**1,1,Answer()
exten=>**1,n,Wait(2)
exten=>**1,n,VoiceMailMain(s001@default)
exten=>**1,n,HangUp()
```

Examinons la première partie, c'est à dire les 5 premières lignes, je n'expliquerai pas chaque commande Asterisk (par exemple Answer()) pour cela il suffit d'aller lire la documentation, mais on va plutot s'arreter sur la séquence.

1. Sur un appel en entrée, tout de suite la première action est de décrocher.

2. Ensuite renvoyer une sonnerie, ainsi l'appelant a l'impression que le poste n'est toujours pas décroché mais qu'il continue à sonner.

3. On attend 10 secondes (a peut pret 3 sonneries)

4. Puis on déclenche la sonnerie sur le poste SIP (ici le 001)

5. Si pas de décrochage du poste alors on envoie le répondeur et le message sera envoyé par mail.

Pour la compréhension du pourquoi de cette séquense se repporter au début de l'article.

Ensuite nous trouvons le plan de dial sur un appel sortant, et sur l'appel en local du **1 qui permet de consulter la messagerie**.

### Fichier mailbox.conf

Dans le fichier sip.conf nous avons paramétré un envoie de mail (ligne `mailbox=001@default`) sur le compte SIP 001, il faut juste ajouter la configuration correspondante dans le fichier mailbox.conf

Se placer dans la section [default] et ajouter la ligne suivante

`001 => 4242,PRENOM NOM,NOM.PRENOM@gmail.com`

Bien sur il faut remplacer par les bonnes informations pour votre mail, attention cette fonction s'appuie sur la fonction mail de linux il faut donc l'installer et paramétrer un postfix local pour que cela puisse partir, mais ca c'est de la configuration du mail c'est une autre histoire.

# Conclusion

Voila, c'est terminé, il ne reste qu'a installer un logiciel de type SIP-Phone sur votre PC ou acheter un poste SIP compatible Asterisk, comme on dit souvent la seule limite de ce que vous allez pouvoir faire avec Asterisk c'est votre imagination.

Je tiens à remercier très chaleureusement **Mickael K.**  qui m'a aidé pour la configuration Asterisk, si jamais vous avez des questions ou besoin d'aide pour une configuration il suffit de faire une mention sur **@ptro46** sur le [discord du Comptoir Sécu](http://discord.comptoirsecu.fr).
