---
title: "Défense — WannaCry"
authors:
  - jil
date: 2019-09-08
image:  /images/covers/2019-07-27-defense-wannacry.jpg
categories:
  - Article
tags:
  - defense
  - wannacry
---

Voici le premier billet d’une courte série axée sur la défense contre certains types de menaces. Mis bout à bout, ils ambitionnent de permettre au néophyte de comprendre comment construire sa défense en fonction des moyens dont il dispose, tout en limitant les discussions théologiques, pour lesquelles nous vous invitons sur [Discord][discord].

# WannaCry, résumé de ce qui nous importe

WannaCrypt est un vers de type rançongiciel qui s’est propagé à travers l’Internet par une vulnérabilité dans le protocole SMBv1 (EternalBlue). Cette version du protocole d’échange de fichiers de Windows est obsolète (utilisée jusqu’à Windows 2003 et XP), au point que depuis Windows 10 1709, elle est désactivée [après 15 jours d’inutilisation][smbv1-disable] en cas de mise à niveau et elle n’est plus installée lors d’une *fresh install*.

Pour se propager, WannaCry tente de scanner des IPs au hasard sur Internet ainsi que celles du réseau local de la machine infectée.

# Maîtriser ses frontières

On ne le répétera jamais assez : moins il y a de services en écoute sur Internet, moins vous aurez d’ennuis. 

## Vérifier ou découvrir son exposition

Pour cette étape, vous avez besoin de connaître les adresses IPs de vos connexions Internet. Une tâche plutôt simple en PME et bien galère dans une multinationale (boxes Internet installées en scred et connectées au réseau, bonjour !). 

### Si j’suis sans le sou, mais que j’ai le temps

Ensuite, pour rester dans les clous de la législation française, vous devez vérifier qui gère les Address Spaces par un whois et obtenir l’autorisation de scan auprès du gestionnaire. Vous devez aussi obtenir l’autorisation de scan auprès du fournisseur de votre point de sortie (qui émettra le scan). Ce serait dommage de perdre la connexion parce que votre fournisseur vous aura considéré comme un client compromis qui émet des scans. De plus, réalisez ce scan depuis un système tiers pour éviter toute blague de vos routeurs et pare-feux qui vous laisseraient passer parce que vous venez du réseau interne.

Dans les faits, si vous émettez un scan ciblé, personne ne s’en inquiétera. Par contre, si vous scannez comme un goret des plages d’IP, vous risquez d’avoir des surprises.

Dans le cas qui nous intéresse, vous allez scanner les [100 ports les plus courants][nmap-top100] :

	nmap -F 127.0.0.1
	nmap -sT -sU 127.0.0.1 -p 135,139,445,3389

En remplaçant l’IP par vos cibles. 

### Si j'ai de l’argent (ou que je suis prévoyant)

Direction [Shodan.io](https://www.shodan.io/search). Il y a une braderie par an (au Black Friday, il me semble) durant laquelle le prix de l’abonnement de base descend sous les 10$. Pour avoir les résulats sur tout un réseau, il existe l’opérateur net :

	net:10.1.2.3/26

Si vos IPs ne sont pas contigües, entrez juste une IP. 

N’hésitez pas à chercher aussi des mots-clés liés à l’entreprise.

[Censys][censys] est une alternative à Shodan.

## Limiter l’exposition

Si vous trouvez quelque chose d’ouvert alors que, à tout hasard, vous êtes chez vous, vous avez trois options :

1. Rien à faire là tout court
2. C’est un service utile, mais il devrait être restreint
3. C’est un service légitime qui doit être ouvert aux quatre vents (un accès VPN par exemple)

Pour l’explication détaillée de comment exposer des services sur Internet, vous pouvez vous référer au [guide de l’ANSSI][anssi-dmz]. Ma conception de la chose, c’est que le réseau de votre entreprise est trop précieux pour qu’il héberge des services web destinés au monde entier, à l’exception du VPN et de l’Identity Provider (dans le cas où vous autorisez l’accès à des applications en SaaS sur des périphériques sans VPN). Les Wordpress et consorts sont très bien chez des hébergeurs sans lien réseau avec votre infrastructure. Il y a des exceptions, mais ici, on parle de WannaCry. Donc un contrôleur de domaine ou un serveur de fichiers, ça ne s’expose pas sur Internet, et si c’était vraiment nécessaire, ce serait avec un filtrage sur l’origine de la connexion.

Pour réduire l’exposition, pensez bien aux deux étapes : d’abord fermer le flux sur le pare-feu frontal, mais surtout éteindre le service, voire le serveur concerné.

# Contrer le déplacement latéral

WannaCry se déplace en scannant le réseau local. Comme on finit toujours par trouver des machines qui ont plusieurs interfaces réseau, on bondit de réseau en réseau et c’est la cata.

Il n’y a aucune raison d’avoir du trafic SMB entre les postes clients. La solution ne coûte pas cher, mais elle est un peu casse-pied à implémenter compte tenu du peu de soin que Microsoft apporte à son pare-feu.

L’approche supportée consiste à créer une règle de pare-feu complémentaire de type `Deny` en GPO sur les ports 135 (RPC), 139 (Netbios), 445 (SMB), 3389 (RDP) en TCP et UDP. En faisant cela, vous bloquez l’utilisation de toute administration à distance. Radical mais peu réaliste. Toutefois, lorsqu’on veut maintenir ces ports ouverts depuis le réseau d’administration (ou, à défaut, depuis les deux IP fixes des postes des admins), le support Microsoft recommande de créer un `deny` avec tous les réseaux internes (mouhahaha).

La version non supportée consiste à coupler deux choses : 

1. Une règle de type `Allow` sur les ports 135, 445, 3389 depuis les IPs d’administration
2. Un [script Powershell][fix-fw] qui désactive les règles installées de base sur Windows 10

Avec le temps, il est possible qu’il y ait des trous, mais c’est déjà un bon début. Rien ne vous empêche de nmap votre réseau interne pour vérifier.

## Si je suis bien lotti

Si vous disposez d’une gestion centralisée du pare-feu par votre agent d’antivirus HIDS qui tabasse les 0-day mais pas WannaCry, vous devriez disposer d’une option élégante pour réaliser ce filtrage.

# Est-ce tout ?

Si ces deux mesures (réduction d’exposition et limitation de la connexion interpostes) sont implémentées, les seules cibles restantes sont vos serveurs. De base, vous pouvez désactiver SMBv1 sur tout ce qui n’est pas un serveur de fichiers. Pour ces derniers, vérifiez si vous avez des clients XP ou 2003. Si ce n’est pas le cas, vous pouvez aussi le couper.

Pour rappel, l’application des GPO requiert l’accès au partage de fichiers SYSVOL sur les contrôleurs de domaine. Un serveur 2003 aura besoin de SMBv1 pour trouver ses GPO. Si vous vivez dans un tel environnement, vous aurez intérêt à créer un réseau spécifique pour les serveurs grabataires afin de leur assigner dans les *Sites and services* un contrôleur de domaine unique, le seul à avoir SMBv1 d’activé (de préférence un [*read-only domain controller*][rodc]), et que vous rendrez inaccessible des autres postes clients.


[anssi-dmz]: https://www.ssi.gouv.fr/administration/guide/definition-dune-architecture-de-passerelle-dinterconnexion-securisee/
[censys]: https://censys.io
[discord]: http://discord.comptoirsecu.fr
[fix-fw]: https://gist.github.com/Jil/c2b6f957e01dcbe3f1f1f0e99cf8a1cc
[nmap-top100]: https://nmap.org/book/nmap-services.html
[rodc]:https://docs.microsoft.com/en-us/previous-versions/windows/it-pro/windows-server-2008-R2-and-2008/cc772234(v=ws.10)
[smbv1-disable]: https://support.microsoft.com/en-us/help/4034314/smbv1-is-not-installed-by-default-in-windows

