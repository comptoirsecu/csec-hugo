---
title: "Défense — NotPetya"
authors:
  - jil
date: 2019-10-08
image:  /images/covers/2019-07-27-defense-wannacry.jpg
categories:
  - Article
tags:
  - defense
  - notpetya
---

Voici le deuxième billet d’une courte série axée sur la défense contre certains types de menaces. Mis bout à bout, ils ambitionnent de permettre au néophyte de comprendre comment construire sa défense en fonction des moyens dont il dispose, tout en limitant les discussions théologiques, pour lesquelles nous vous invitons sur [Discord][discord].

# NotPetya, résumé de ce qui nous importe

NotPetya entre dans la catégorie des *wipers*, car sous l’image d’un rançongiciel, son seul dessein est la destruction à grande échelle du réseau contaminé. Il a été délivré par la mise à jour du logiciel de déclaration de taxes ukrainien MeDoc, sa chaîne de build ayant été compromise par un attaquant.

NotPetya escalade ses privilèges pour récupérer les identifiants des comptes administrateurs en mémoire et tente de se délacer latéralement avec (à coup de `psexec` et `wmic`). De plus, il ne se prive pas du déplacement par EternalBlue, pour laquelle je vous renvoie à l’article [Défense — WannaCry][wannacry].

# Contrer l’élévation ou le dump d’identifiants

Un grand nombre de logiciels malveillants requièrent le privilège [SeDebugPrivilege][debug] qui permet de s’attacher à un processus qu’on ne possède pas. Rares sont les cas où vous avez besoin de déboguer des processus système. Le durcissement des postes et serveurs passe par une GPO : 

	Computer Configuration\Windows Settings\Security Settings\Local Policies\User Rights Assignment\Debug Programs

Soit vous retirez purement et simplement le groupe Administrateurs, soit vous vous gardez une porte de sortie avec un groupe du domaine dédié à cet usage (pour éviter d’avoir à changer la GPO en cas de besoin). Ce groupe devient un groupe sensible (à surveiller et mettre hors de portée des admins, toussa toussa).

Il est important de noter que ce n’est pas magique : il existe d’autres techniques pour récupérer les identifiants.


# Contrer le déplacement latéral

Tout comme WannaCry, il ne faut pas laisser la possibilité aux postes clients de parler entre eux. Les ports d’administration ne doivent être ouverts qu’en provenance du réseau d’administration. Le détail est dans l’article [Défense — WannaCry][wannacry]


# Le problème des privilèges globaux

## Les administrateurs 

Les administrateurs d’une flotte de PC aiment pouvoir utiliser leur compte pour devenir administrateurs de tout poste sur lequel ils interviennent. 

À bon entendeur, rappelons que le compte utilisé au quotidien par un administrateur ne doit disposer d’aucun privilège. C’est un deuxième compte qui tient ses privilèges (et ceux-ci ne sont pas *Domain Admin* ni aucun groupe *built-in* d’Active Directory ; [plus d’infos ici][builtin]).

Généralement, on retrouve un groupe du domaine (du type « Admin_all_computers ») qui est poussé par GPO dans le groupe local Administrators de tous les postes clients du domaine. D’un point de vue technique, on gagne en souplesse en maintenabilité avec les [GPP et leur filtrage][gpp].

Le hic, c’est que si la machine contaminée contient dans son antémémoire (*cache*) les identifiants d’un tel compte, il sera utilisé pour se propager. Ce qui bloquera la propagation sera le filtrage réseau (cf. *Contrer le déplacement latéral*) et l’interdiction de se connecter aux machines d’administration avec ces comptes (à défaut, l’attaquant atteindra le réseau d’administration et pourra pivoter sur toutes les autres machines). 

Il existe plusieurs moyens de limiter ce risque.

L’option la plus brutale consiste à retirer les droits administrateur aux administrateurs de la flotte et à ne leur laisser que les privilèges nécessaires à la prise de main à distance. Quand ils auront besoin d’une élévation, ils utiliseront le compte administrateur local dont le mot de passe est géré par LAPS (on y reviendra plus bas). Ils ne conserveront sur l’Active Directory que les droits pour gérer le join/leave de la machine au domaine. Haro garanti, car il faudra chaque fois aller chercher le mot de passe administrateur local.

Pour limiter la capacité d’un attaquant à jouer avec les jetons Kerberos, AD 2008 a introduit l’option [*This account is sensitive and cannot be delegated*][nodelegation]. En gros, cela empêche d’utiliser le jeton auprès d’un autre service (sur la même machine ou ailleurs) qui, lorsqu’il accepte la session, va lancer un process qui s’exécute *en tant que* le compte indiqué dans le jeton. AD 2012 R2 a introduit un groupe [*Protected Users*][protectedusers] qui va plus loin : interdiction du NTLM, Digest, CredSSP, pas de mise en antémémoire des mots de passe, limitation des algorithmes de chiffrement utilisables, et les jetons ne sont valables que quatre heures. Ce paramètre engendre l’activation du premier (compte sensible ne pouvant être délégué).

## Les comptes de service

L’on trouve souvent les comptes de service des antivirus qui ont des privilèges exorbitants sur toutes les machines. Or, ils servent uniquement au déploiement de l’agent par la console centrale. Une fois l’antivirus installé, il s’exécute soit en *Local System* soit en *Network Service* et n’a plus besoin du compte. 

Il s’avère délicat de supprimer ces comptes tout en conservant le support de l’antivirus.

## Les comptes locaux 

Tous les comptes locaux devraient avoir des mots de passe propres à la machine. Pour le compte Administrateur, Microsoft propose la rotation périodique du mot de passe des comptes administrateurs avec [Local Admin Password Solution][laps] (l’article est même en français !). Pour les autres, il vous appartient de trouver une solution.

Comme c’est compliqué, AD 2012 R2 introduit [un SID spécial][sidlocal] à cet effet (ils ont été backportés sur Windows 2008 R2, KB2871997). Il permet d’interdire l’utilisation de comptes locaux à distance (*Deny access to this computer from the network*). L’article lié vous explique tout.


[builtin]: https://www.jasonfilley.com/display/JF/Active+Directory+Built-In+Groups+Self-Elevation
[debug]: https://docs.microsoft.com/en-us/windows/security/threat-protection/security-policy-settings/debug-programs
[gpp]: http://www.checkyourlogs.net/?p=22921
[laps]: https://blogs.technet.microsoft.com/arnaud/2015/11/25/local-admin-password-solution-laps/
[nodelegation]: https://blogs.technet.microsoft.com/poshchap/2015/05/01/security-focus-analysing-account-is-sensitive-and-cannot-be-delegated-for-privileged-accounts/
[protectedusers]: https://docs.microsoft.com/en-us/previous-versions/windows/it-pro/windows-server-2012-R2-and-2012/dn466518(v%3dws.11)
[sidlocal]: https://blogs.technet.microsoft.com/secguide/2014/09/02/blocking-remote-use-of-local-accounts/
[wannacry]: https://comptoirsecu.fr
