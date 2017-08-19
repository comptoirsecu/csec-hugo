---
title: "Le NIST signe-t-il l’arrêt de mort de nos politiques de mot de passe ?"
authors:
  - jil
date: 2017-08-21
image: /images/covers/2017-08-08-bitcoin-laundering.jpg
categories:
  - Article
tags:
  - NIST
  - Mot-de-passe
  - CNIL
---

# À bas le renouvellement périodique à intervalle fixe !

La nouvelle a défrayé la chronique : le dogme du mot de passe complexe de huit caractères s’effondre devant l’[analyse statistique de la Vérité][mesure]. Cela fait dix ans que l’on s’en doute et, pour reprendre la formule préférée des publicitaires : c’est prouvé scientifiquement. Plus besoin de se torturer l’esprit pour concevoir un mot de passe à la « 3v1L#P@ssword666 » et, au bout de trois mois, par un intense effort psychique, le transformer en « 3v1L#P@ssword667 ». Hourra !

Le NIST l’a annoncé, [Troy Hunt][troy] l’a relayé, mais vous trouverez un article très intéressant du côté de la [Federal Trade Commission][ftc], avec de nombreux renvois utiles. 

[![Dilbert 10 Sept 2005](http://assets.amuniversal.com/e47ff0606d5001301d7a001dd8b71c47)][dilbert]

# Mais pourquoi ? 

Il n’est aucunement question de remettre en cause la robustesse d’un mot de passe aléatoire de huit caractères contenant des minuscules, majuscules, chiffres, caractères spéciaux (dont les frimousses Unicode font partie), même si l’on porte aujourd’hui ce nombre à douze caractères. Il s’agit simplement de rappeler que la méthode de construction des mots de passe *n’est pas* aléatoire, car notre cerveau fonctionne par association de mémoire. [Des modèles se dégagent][hydraze], avec une majuscule au début et les chiffres plutôt vers la fin.

Il est surtout question de repositionner cette mesure de sécurité (le mot de passe) dans le cadre [des menaces][menace] auxquelles le système d’information est exposé, ainsi que vis-à-vis des mesures complémentaires de protection.

Bref, il s’agit de confronter l’avantage réel de la politique de mot de passe aux contraintes qu’elle induit.

# Les règles de composition disparaissent

Un mot de passe aléatoire et long, disons « qdztyomxpfei » serait plus sûr que « Pl@ton-428 » réutilisé sur trente-sept sites. Un scoop. 

Les règles de composition se retirent au profit de l’emploi systématique d’un gestionnaire de mots de passe pour avoir un secret différent par système. Il s’agit d’amoindrir les conséquences d’une compromission unitaire, qu’elle résulte d’une interception, d’une divulgation, d’une divination ou d’une inférence.

Il s’agit là d’une avancée considérable si l’on se place dans une stratégie de réduction du risque lié à l’usage des mots de passe. La vulnérabilité tirée de la paresse humaine est reportée sur l’accès au gestionnaire de mot de passe, donc notamment sur son exposition (en ligne ou hors-ligne), le secret qui le protège (clé, passphrase) et la robustesse du chiffrement utilisé. 

# Plus d’astuce, mon petit ! Tu te débrouilles maintenant.

Fini le temps où il était possible d’introduire une astuce qui permettait de retrouver son mot de passe (si si, ça existe sous Windows 7), les indications semblaient trop claires pour les attaquants. Exit aussi les questions secrètes auxquelles l’on peut répondre avec le profil public de la personne.

Puisqu’il n’y a plus de complexité inhérente à la composition, l’utilisateur devrait se souvenir de son mot de passe.

# Le débat de la liste noire

Troy et ses 330 millions de mots de passe à blacklister se heurtent à l’utilisabilité d’un tel système (quel client accepterait de se faire refouler plusieurs mot de passe de suite au motif qu’ils ont fuité ailleurs ?) ainsi qu’à [la preuve du faible gain][blacklist] procuré par les listes noires (au-delà des valeurs les plus usitées, qui sont [obligatoires][nist]).

Pour évaluer la robustesse d’un mot de passe issu du cerveau humain, l’on doit fournir à l’utilisateur une indication visuelle pour éclairer son choix. On pense alors à [zxcvbn][zxcvbn].

Cependant, pour en revenir à Troy, dans le cadre actuel où les mots de passe sont réutilisés entre les systèmes, savoir que son mot de passe a été découvert peut avoir un sens, dans la mesure où la personne concernée peut s’inquiéter de l’impact réel de la divulgation et prendre les mesures qu’elle estime appropriées. 

# Le renouvellement périodique perd la bataille

Le renouvellement périodique présente l’intérêt d’invalider un mot de passe au bout d’un temps défini et prévisible. Si l’on prend la politique largement répandue d’un renouvellement trimestriel, cela laisse à l’attaquant un maximum de trois mois pour créer un compte secondaire sous sa maîtrise ou, plus simplement, pour installer un shell résilient qui bénéficiera de la session valide de l’utilisateur, quand bien même ce dernier aurait renouvelé son mot de passe. Sans avoir trouvé de base statistique du temps nécessaire à un pentester expérimenté pour installer un shell en userspace avec une évasion HTTPS lors d’une attaque ciblée, je pense pouvoir avancer que même un renouvellement périodique d’un mois serait vain. J’avais croisé une publication de Microsoft, hélas perdue depuis, qui insistait sur l’immédiateté de l’usage d’un mot de passe découvert par un attaquant.

C’est l’occasion de rappeler que le renouvellement périodique ne doit pas être confondu avec la mesure de désactivation des comptes inactifs.

Toutefois, avec les bases de mot de passe qui s’enfuient par millions dans les profondeurs du deep dark web pour y mener une vie trépidante en compagnie des bitcoins sous la cape de Tor (ou comment placer les mots-clés de référencement en une phrase maligne), le renouvellement périodique demeure une bonne idée.

# Détecter et réagir

Pour envisager la suppression du renouvellement périodique, il est nécessaire d’être en mesure de détecter la compromission probable du mot de passe. 

Pour les clients d’Office 365 disposant de la licence EMS (E5), Azure Active Directory est votre allié, avec pour limite qu’il est aveugle à l’activité de l’ADFS (qui gère l’authentification) que vous avez pris soin de conserver chez vous sans synchroniser les condensats de mot de passe chez Microsoft. Un petit [RTFM][o365] vous fera le plus grand bien.

Il demeure essentiel d’être en mesure de forcer le renouvellement du mot de passe en cas de suspicion, chose que vous avez depuis belle lurette, en lecteur assidu de [l’OWASP][owasp], inclus dans les spécifications des applications qui réalisent une authentification interne.

# Protégé par le deuxième facteur

En outre, pour réduire l’impact de la divulgation, de l’inférence ou de la divination d’un mot de passe, il est conseillé de requérir un deuxième facteur d’authentification. 

Si vous disposez des capacités de détection qui vous permettent d’attribuer un niveau de risque à chaque authentification (nouveau terminal, accès via serveur mandataire, localisation inhabituelle, accès hors du réseau de l’entreprise, accès à une ressource critique, etc.), ce mécanisme un peu pénible peut être requis au-delà d’un niveau de risque pertinent. 

Sans cela, je vous déconseille de le rendre systématique sauf à disposer d’un jeton à clic, tel que l’envoi de la séquence par une pression sur une Yubikey (norme OATH) ou la validation d’une notification mobile du Microsoft Authenticator (ne fonctionne que pour les authentifications Azure).

Quoi qu’il en soit, comme le rappelle [Alex Weinert][alex], le deuxième facteur n’est second que lorsque le premier existe. Il demeure vital que le mot de passe demeure secret et de le renouveler en cas de compromission suspectée ou avérée.

# La réglementation dans tout ça ?

La CNIL a émis une [recommandation relative aux mots de passe][cnil] en janvier 2017, laquelle vient préciser l’obligation de protection des données personnelles de la loi. Autant dire qu’il vous faudra déployer de grands efforts pour justifier à un juge que vous y avez dérogé.

Dans l’hypothèse où le système d’authentification est pourvu d’un mécanisme d’anti-bruteforce, le mot de passe suivant est conforme à la réglementation : « J’suis trop fan du comptoir sécu ! » (vous noterez la vraie apostrophe inclinée, amateurs de [bépo](http://bepo.fr)). Vous avez donc l’obligation d’avoir trois types de caractères différents (minuscules, majuscules, chiffres, caractères spéciaux).

Pour en revenir au renouvellement : « La commission recommande que le responsable de traitement veille à imposer un renouvellement du mot de passe selon une périodicité pertinente et raisonnable, qui dépend notamment de la complexité imposée du mot de passe, des données traitées et des risques auxquels il est exposé. »

Le « recommande » s’interprète ici comme le « SHOULD » des [RFC][rfc]. Je ne peux pas prononcer d’avis juridique qualifié, mais si vous disposez des systèmes de détection d’activités suspectes et de réinitialisation forcée, cela me semble conforme.

D’autres réglementations propres à votre domaine d’activité peuvent aussi entraver la disparition du renouvellement périodique des mots de passe.

# TL;DR

En France, pour l’instant, le cas moyen (allez vraiment lire la [réglementation][cnil]) :

* Un mécanisme anti-bruteforce
* 8 caractères minimum
** 5 si authentification à deux facteurs
* 3/4 des classes de caractères : minuscules, majuscules, chiffres, caractères spéciaux 
** sauf si authentification à deux facteurs
* un renouvellement selon une périodicité pertinente et raisonnable

Pour limiter le renouvellement, il vous faut des capacités de détection des compromissions.


[alex]: https://blogs.microsoft.com/microsoftsecure/2017/06/05/three-basic-security-hygiene-tips-from-microsofts-identity-team/
[blacklist]: https://www.internetsociety.org/sites/default/files/usec2017_01_3_Habib_paper.pdf
[cnil]: https://www.legifrance.gouv.fr/affichTexte.do?cidTexte=JORFTEXT000033928007
[dilbert]: http://dilbert.com/strip/2005-09-10
[ftc]: https://www.ftc.gov/news-events/blogs/techftc/2016/03/time-rethink-mandatory-password-changes
[hydraze]: http://www.passwordresearch.com/stats/statistic376.html
[menace]: http://www.cerias.purdue.edu/site/blog/post/password-change-myths/
[mesure]: https://www.cs.unc.edu/~reiter/papers/2010/CCS.pdf
[nist]: https://pages.nist.gov/800-63-3/sp800-63b/sec5_authenticators.html#memsecretver
[o365]: https://docs.microsoft.com/en-us/azure/active-directory/active-directory-identityprotection
[owasp]: https://www.owasp.org/index.php/Password_Storage_Cheat_Sheet
[rfc]: https://www.ietf.org/rfc/rfc2119.txt
[troy]: https://www.troyhunt.com/passwords-evolved-authentication-guidance-for-the-modern-era/
[zxcvbn]: https://dl.dropboxusercontent.com/u/209/zxcvbn/test/index.html
