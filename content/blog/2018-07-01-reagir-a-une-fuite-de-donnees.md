---
title: "Kit de survie : réagir à une compromission de mailbox"
authors:
  - jil
date: 2018-07-01
image:  /images/covers/2018-07-01-banniere-mailbox.jpg
categories:
  - Article
tags:
  - incident
  - Office365
---

# Les plans, c’était pour avant !

Vos chantiers de sécurisation avancent bon gré mal gré, quand, tout à coup, survient le tant redouté incident de sécurité. Et galère ! Vous aviez prévu de vous préparer sérieusement au traitement des incidents quelques semaines plus tard. Un report louable, puisqu’il avait fallu parer au plus urgent. 

Vous y voilà donc ! Vous êtes sans le doute le seul à bord à ne pas paniquer, mais vous êtes stressé et oublierez des choses, c’est pourquoi je tente d’écrire ce kit de survie sur un incident type. Je rappelle que le blog est communautaire, vous pouvez donc enrichir cet article, ou venir en parler sur [Discord][discord] et nous le mettrons à jour pour être encore plus utile !


# Dois-je être pressé ?

> Ne confonds pas vitesse et précipitation, padawan

Si vous étiez dans le cas d’une détection d’attaque par votre équipe de détection (SOC par-ci, SOC par-là), vous ne liriez pas cet article. Vous arrivez donc trop tard. La probabilité que vos attaquants soient en train d’opérer quelque chose d’utile et qu’en déconnectant toute votre infrastructure, vous les en empêchiez est infinitésimale.

Respirez, la journée d’un salarié au forfait jour en France compte jusqu’à treize heures, vous avez le temps. Vous devrez :
	
	1. Cerner la situation
	2. Rassurer vos interlocuteurs
	3. Ouvrir le tiroir-caisse
	4. Analyser la situation
	5. Réagir selon la criticité de la situation
	6. Suivre l’enquête sur plusieurs semaines
	7. Tirer les enseignements
	8. Présenter vos conclusions

Ces étapes ne sont pas séquentielles. Vous alternerez entre les étapes 3 à 6.


# 0. Tout documenter

Avant toute chose, ouvrez un tableur avec deux onglets :

	* la timeline de compromission
	* la timeline de vos actions

Principe d’auditabilité : vous pourrez retrouver tout ce que vous avez fait pour voir si vous avez oublié quelque chose. Un auditeur externe pourra rapidement comprendre la manière dont vous avez traité l’incident et explorer d’autres pistes. Ce fichier va évoluer au fil de votre compréhension du problème. N’enlevez jamais rien, barrez tout au plus. Indiquez vos revirements dans la timeline de vos actions. Vous vous direz merci à vous-mêmes quand il faudra faire le rapport, cinq à six semaines plus tard.

*Exemple fictif de timeline de compromission:*

| Date       	| Heure 	| Fuseau 	| Source                          	| Acteur        	| Action                                                             	| Commentaires                            	|
|------------	|-------	|--------	|---------------------------------	|---------------	|--------------------------------------------------------------------	|-----------------------------------------	|
| 05/04/2018 	| 10:37 	| CEST   	| Copie eml depuis mailbox de FD. 	| F. LeComptable 	| Envoi du mail (legit) à P. QuiDoitPayer                              	| Sujet: facturation en retard, pénalités 	|
| 05/04/2018 	| 10:50 	| CEST   	| Copie eml depuis mailbox de FD. 	| P. QuiDoitPayer  	| Réponse (legit) pour éviter les pénalités contre paiement immédiat 	| Re: facturation en retard, pénalités    	|
| 05/04/2018 	| 10:59 	| CEST   	| Copie eml depuis mailbox de FD. 	| F. LeComptable 	| Réponse « ok » depuis smartphone                                   	| Re: facturation en retard, pénalités    	|
| 05/04/2018 	| 11:45 	| CEST   	| Forward mail de F. LeComptable   	| Attaquant      	| Demande de paiement sur RIB indiqué dans le mail                   	| Re: facturation en retard, pénalités    	|

*Exemple fictif de timeline de vos actions:*

| Date       | Heure | Acteur      | Action                                                                                |
|------------|-------|-------------|---------------------------------------------------------------------------------------|
| 05/04/2018 | 14:35 | M. Sécurité | Prise en charge de l’incident signalé par F. LeComptable                              |
| 05/04/2018 | 14:36 | M. Sécurité | Appel de F. LeComptable. Prise de main à distance, export eml des mails avec en-têtes |


# 1. Cerner la situation : la première analyse

Prenez connaissance des éléments à votre disposition. Appelez le lanceur d’alerte, mesurez son niveau de tension, rassurez-le, interrogez-le en douceur **pour comprendre les faits**. Il ne s’agit pas de trouver la cause, le vecteur de compromission ou que sais-je encore, mais de mesurer la gravité de l’incident. 

S’il s’agit d’une fraude financière : a-t-on payé ? Si des documents ont été publiés : à vue de nez, est-ce que, pour le métier, ils renfermaient des informations sensibles ? (À moins d’une excellente culture de la classification des données, demandez plutôt ce qu’il y a dans les documents, est-ce qu’il y a des noms de clients, le contenu de votre golden parachute, des coordonnées, des intitulés de projet, l’identifiant de votre cryptoportefeuille, etc…)

Dès que vous savez de quoi il retourne, vous pouvez déduire s’il est nécessaire de déclencher une cellule de crise. En bref, s’il y a des risques que l’information sur l’attaque sorte de l’organisation et puisse être manipulée par les médias ou par un concurrent, s’il y a un impact sur des vies, s’il y a un impact client, s’il y a un impact financier non négligeable, il est de bon aloi de déclencher la cellule de crise dans laquelle vous allez retrouver un représentant de la direction générale, de la direction juridique et de la communication. En l’absence d’un processus défini, appelez votre chef et demandez-lui de vous aider : qu’il s’occupe d’alerter pendant que vous enquêtez.


# 2. Rassurer

L’usurpation d’identité peut-être un traumatisme pour une personne qui n’y est pas préparée, et qui se demande si on a utilisé son ordinateur, si le fait qu’il y ait son nom risque d’engendrer son licenciement, qu’est-ce qu’elle doit faire, etc.. Votre connaissance de l’organisation, couplée à votre formation en ingénierie psychosociale (oui, justifiez-la pour la réponse aux incidents), devrait vous permettre de trouver les bonnes paroles pour ne pas paraître blasé (« on en voit tous les jours, vous savez ») tout en écartant les craintes vis-à-vis de l’avenir. Peu importe si vous pensez que cette personne a une part de responsabilité, car la première chose dont vous avez besoin, c’est de sa coopération. Il sera toujours temps, plus tard, d’établir ce qui aurait pu être fait. Veillez donc à ne prendre aucun engagement que vous ne pourriez tenir en cas de procédure disciplinaire ou pénale qui, par nature, échapperont à votre contrôle.


# 3. Ouvrir le tiroir-caisse

Si l’incident est sérieux, vous aurez besoin d’assistance extérieure. Partagez la qualification de l’incident avec votre chef et confirmez les montants de dépense possible. La recherche de preuves est une activité chronophage, les tarifs journaliers sont élevés.

(Si on pouvait glisser une estimation ici, ce serait bien. Sur un incident mail avec les logs, 5j; pour une analyse forensic d’un poste 5-8j ?)


# 4. Analyser la situation (simulant le chaos d’une réaction non procédurée)

Les impacts étant connus (dans notre exemple, la fraude a été détectée à temps), il s’agit d’allier réaction et compréhension ce qu’il s’est passé.

Je ne fais pas partie des ayatollahs du changement immédiat de mot de passe, d’une part car il y a de grandes chances que l’attaque soit terminée ou suspendue, et d’autre part, car le vecteur d’intrusion s’affranchit souvent des mots de passe. Déroulons notre cas.

## Origine de la fraude

À 11 h 45, P. QuiDoitPayer reçoit un message frauduleux avec un autre RIB pour régler la facture en retard.

Dans le cas typique, ce message a été émis par vos serveurs de messagerie. Vous ne retrouvez pas le mail en question dans la boîte mail de F. LeComptable, qui nie l’avoir écrit. Vous avez néanmoins sa trace dans les journaux de la passerelle de messagerie. Vous remontez alors les journaux de connexion à la messagerie et trouvez une connexion suspecte depuis un autre pays. L’authentification n’étant pas protégée par un deuxième facteur : phishing, accès à distance, recherche de messages intéressants, mise en place d’une règle de redirection et d’une règle de tri pour supprimer les messages gênants, exploitation, fin. Il ne reste plus qu’à changer le mot de passe et vérifier les règles.

Pour rendre cet article plus savoureux, notre exemple est différent : 

* le message reçu par P. QuiDoitPayer a été émis par un serveur tiers avec un nom de domaine avoisinant;
* P. QuiDoitPayer fait partie de l’organisation (pour éviter d’avorter l’analyse à cause d’un système hors périmètre);
* le message contient l’historique de la conversation.

L’attaquant a eu accès à la messagerie. Ce n’est pas juste un bon timing fondé sur un coup de téléphone. Il y a eu intrusion.

Chaque piste doit être explorée.

## Piste 1 : compromission de la messagerie de l’expéditeur (F. LeComptable)

### Journal (intime)

Mauvaise nouvelle : vous êtes client d’Office 365 et n’avez pas l’abonnement Azure AD Premium (ou l’Entreprise Mobiliy Suite); vous ne pouvez pas accéder aux journaux détaillés de connexion côté Azure AD. Vous disposez toutefois des évènements UserLogIn et UserLogInFailure. Or, comme vous avez activé l’authentification à deux facteurs, les échecs correspondent à du brute-force par des botnets et les succès, à une authentification approuvée depuis Microsoft Authenticator sur le téléphone. Ce qu’il vous manque c’est le détail des connexions à la boîte mail.

Crowdstrike a publié [un article et un outil][apio365] sur une API non documentée qui vous donne accès à l’historique des actions effectuées dans une boîte mail O365. Vous y trouverez les connexions avec leurs IP, les lectures des courriels, les suppressions, etc.

Avec l’abonnement Azure AD Premium (mais aussi la version d’essai 30 jours, je pense), vous avez accès aux 30 derniers jours de logs Azure AD avec, pour chaque authentification, le détail du protocole : est-ce un token valide, une authentification MFA, un application password… Ce qui laisse apparaître de nouvelles IPs, dont celle de l’attaquant. C’est sans doute la première chose à débloquer et le journal le plus utile.

Ces informations ne figurent pas dans les audits logs, parce que les audits logs ne consolident pas tous les journaux de la plateforme (ne négligez jamais les interfaces d’admin dédiées à un module). Pour remonter davantage d’informations, il faut paramétrer spécifiquement chaque boîte mail en Posh.

### M’as-tu vu ?

Vous ne constatez aucune création de règles de tri (New-InboxRule). Vous ne trouvez, dans l’historique des mails sortants sur le serveur de messagerie, aucun *forward* évident ni aucune communication vers une adresse suspecte.

A priori, vous ne constatez aucune chose anormale. F. LeComptable ne se souvient pas avoir eu de demande d’approbation louche qu’il aurait validée sur son téléphone, ni d’appel suspect.

Dès lors, la compromission de la boîte interne nécessiterait un accès à distance aux ordinateurs de F. LeComptable (son pc ou son téléphone). En vertu du principe de récolte du fruit mûr à portée de main, par lequel un attaquant prend le chemin le plus rapide et le plus facile, cette piste est suspendue.

Tout cela est consigné dans le tableau de suivi. Les journaux sont exportés et stockés (il m’est déjà arrivé de voir les évènements n’apparaître qu’après 48h dans les journaux O365, l’export vous rassure quant à votre santé mentale après coup).

## Piste 2 : compromission de la messagerie du destinataire (P. QuiDoitPayer)

Pas mieux que la piste 1. Vous avez même pensé aux délégations de boîte mail, mais il n’y en avait pas. (Vous remarquerez que, souvent, dans ces situations d’urgence, vous vous apercevez plus tard que vous avez oublié un autre vecteur d’accès, d’où l’absolue nécessité du tableau avec ce que vous avez récolté. Je radote, non ?)

## Piste 3 : contournement du multifacteur

### Zone de confiance (mouhaha !)

Une fonctionnalité existe, qui n’a pas été implémentée dans cet exemple, mais c’est l’occasion d’une digression susceptible de sauver des SI. Imaginons que vous n’ayez pas envie de vous mettre tout le monde à dos et que vous rendiez le deuxième facteur d’authentification non nécessaire lorsque quelqu’un est connecté au réseau d’entreprise. Vous n’auriez alors aucune certitude quant à l’usage légitime des connexions émanant de votre IP d’accès à l’Internet.

Votre réseau peut être pénétré par un cheval de Troie, par une compromission du VPN, ou vous pouvez faire face à une fraude interne (fondée sur les utilisateurs qui prêtent leur mot de passe à leur collègue à cause d’une séparation des tâches imposée dans les systèmes, mais incompatible avec la réalité du métier). 

Or, suspecter une compromission du réseau, c’est amenuiser ses chances de trouver quoi que ce soit. Parce que si vous aviez déployé une protection de tous vos terminaux (Endpoint Detection & Response), vous auriez déjà préparé les moyens pour réagir aux nombreuses alertes ; vous sauriez réagir à un incident ; vous ne liriez pas ce kit de survie ;)

### Tolérance avec les vieux

Puisque les clients natifs des téléphones ne supportent pas l’authentification multifacteur, il existe un moyen de la contourner : le mot de passe d’application. Celui-ci est généré et donné à l’utilisateur, qui doit entrer le mot de passe dans son téléphone. Sauf que ce dernier, bourrinant à 4 auth/minute, est banni pour brute-force quand l’utilisateur saisit le mot de passe ; ce dernier apparaît alors comme invalide. L’utilisateur (et l’agent de l’assistance aussi) va stocker ce mot de passe aléatoire… sur son bureau. Aller à la case accès à distance.

## L’impasse, le scénario APT 

En somme, il n’y a pas de compromission apparente, mais il y a eu intrusion. Ce qui laisse les hypothèses suivantes :

* compromission d’un accès administrateur
* prise de main à distance d’un des ordinateurs (c’est pour ça qu’on dit ordiphone)
* cheval de Troie 
* interception de trafic réseau
* photo du message si le poste n’était pas verrouillé
* sortilège de divination de niveau 9

## Piste 4 : compromission d’un accès administrateur

Il y a de fortes chances que vous n’ayez pas ce qu’il faut dans les journaux. C’est sur des journaux incomplets ou inexistants que s’arrêtent la plupart des investigations.

Dans le cas d’Office 365, l’ensemble des actions de comptes administrateurs sont journalisées, mais accessibles selon votre abonnement. Vous pouvez néanmoins voir qu’il n’y a pas eu d’attribution de droit de délégation, ni de e-Discovery (fonctionnalité permettant de rechercher dans l’ensemble des données hébergées sur la suite Office). Ouf ! Pas de compromission des comptes admins.

Dans le doute, vous vérifiez les journaux de connexion de vos comptes admins (et vous pestez sur leur nombre).

## Piste 5 : prise de main à distance

Soyons clairs : si vous en arrivez au scénario sur l’ordinateur, ça pue. Parce si vous lisez cet article, c’est que vous n’étiez pas prêt à faire face. Subséquemment, votre maturité en matière de journalisation de poste utilisateur avoisine le néant. Et non, dans Windows 10, la configuration par défaut du journal d’évènement n’améliore pas substantiellement la rétention.

> Si le doute est raisonnable, les terminaux, tu séquestreras.

Le meilleur moyen de conserver les traces d’une attaque qui ne réside pas exclusivement en mémoire, c’est de mettre l’ordinateur en lieu sûr et de ne plus y toucher jusqu’à la copie des disques. Copie qui, en cas de procédure judiciaire, doit se faire dans les règles de collecte de preuve légale (ce n’est pas le sujet de l’article).

Si vous ressentez le besoin d’une analyse post-mortem (*forensic*), n’essayez pas par vous-mêmes : sans les logiciels adéquats, c’est la misère. Ouvrez le tiroir-caisse ;)

Cela étant dit.

Consultez les journaux de Teamviewer. Ce n’est pas parce qu’il ne figure pas à votre liste de logiciels qu’il n’est pas installé.

Consultez les journaux d’accès de votre serveur mandataire (*proxy*). S’il opère des classements, regardez les flux non catégorisés qui pourraient trahir une communication vers un serveur de Command & Control.

Sur le téléphone, vérifiez la liste des certificats de confiance (bon courage pour trouver les menus selon les versions), la liste des apps et s’il est rooté. Cela même si vous avez un MDM qui empêche l’accès aux données d’entreprise en cas de configuration inacceptable du téléphone.

Et on ne trouve rien. Toujours rien. Je vous fais grâce des pistes 5 (l’homme du milieu) et 6 (intrusion sur site ou complicité interne).

## Piste 42 : dépité, l’enquêteur partit élever des chèvres en Ardèche

Toujours rien ? Vous avez sans doute raté quelque chose. C’est dans les journaux que vous n’avez pas.

Ne vous inquiétez pas, ça se termine souvent ainsi : avec les hypothèses de ce qui a pu se passer, sans preuve. Pour autant, ce n’était pas vain : allez à la case « enseignements ».


# 5. Contremesures

## Tracfin

Si un virement a été effectué par une banque française, appelez le service fraude de la banque et exigez l’émission d’un signalement Tracfin pour avoir une chance de bloquer les fonds.

## Dépôt de plainte

Selon la nature du préjudice, ou la possibilité d’en subir, le dépôt de plainte peut être incontournable. Il sert à déclencher les assurances, quelques fois à notifier le renseignement intérieur, et à avoir l’avis d’enquêteurs spécialisés.

Si vous êtes en province, c’est la panade. Du moins tel est mon ressenti.

Vous perdrez probablement 10 jours à trouver le bon contact si votre attaque sort du cadre de l’hameçonnage classique.

## Décider du niveau de visibilité

Si vous n’êtes pas sûr de l’étendue de la compromission, la question se pose de la publicité à en donner. Imaginons que les attaquants puissent toujours lire les mails. Il serait regrettable de leur communiquer de nouvelles informations.

Le sujet est épineux, car la dépendance de l’entreprise aux courriels est telle que vous n’avez peut-être aucun médium alternatif de communication pour diffuser une alerte dans l’organisation. Cherchez tout de même du côté d’une passerelle SMS si vous avez un annuaire à jour des portables, ou du côté des mécanismes de notification de votre gestionnaire de flotte mobile. Dans une expérience récente, j’ai découvert qu’une grande partie des utilisateurs d’iPhone ne savaient pas ce qu’était une notification et, de fait, n’avaient pas pris connaissance de la notification du MDM. Eh oui. Remarquez : combien de personnes connaissent le chemin d’évacuation le plus rapide pour rejoindre le point de rassemblement en cas d’alerte incendie ? Qui y parvient à temps pour éviter la suffocation ? Même question dans un hôtel. Eh oui.

## Tomber le domaine d’attaque

Il y a des services spécialisés pour faire tomber un domaine ou des adresses mail utilisés par une fraude. Cela peut prendre du temps, et surtout, cela informe les attaquants de la détection de leur attaque. Vous perdez la capacité à détecter d’autres tentatives qui utiliseraient la même infrastructure.

À terme, il peut être utile de les faire tomber pour protéger vos clients de telles attaques. Tout est question de timing.

## Changer les mots de passe, révoquer les tokens

Sauf à avoir démontré que l’attaque est passée par ailleurs, il faudra révoquer les jetons d’accès et changer les mots de passe.

## Purge

Selon l’étendue de votre expertise, vous pouvez souhaiter détruire le compte AD et en créer un tout neuf pour perdre tout ce qui pourrait avoir été lié au profil. Vous en profitez pour réinstaller complètement le poste et le téléphone.


# 7. Enseignements

Peut-être avez-vous pu expliquer l’attaque. Peut-être pas. En tout cas, vous avez constaté des manques dans la journalisation des actions, voire dans les solutions à déployer pour être en mesure de réagir plus efficacement en cas d’incident.

À ce point de l’histoire, il est important de relativiser. De mettre votre frustration de côté. De raison garder.

La fraude est une réalité. Le risque d’une intrusion est réel. Les conséquences sont prédictibles. C’est cette prédiction, dans l’analyse de risques, qui permet de décider du budget à allouer à la sécurité. Le rôle d’une direction générale, c’est de prendre les bons risques.

Pour couvrir un risque, il existe des assurances. Le principe de l’assurance est de lisser le coût engendré par un dommage. Vous vous êtes fait attaquer, vous avez subi un préjudice. La question est de savoir s’il est acceptable pour l’organisation. Autrement dit, peut-elle l’encaisser ?

Si elle le peut sans broncher, ne vous prenez pas la tête. Poursuivez votre feuille de route comme si rien ne s’était passé. Tout au plus, vous rappelez l’évènement pour appuyer un propos. N’allez pas remuer ciel et terre pour la survenance d’un évènement qui n’a pas bouleversé l’entreprise.

Si vous avez failli y rester, ou qu’un attaquant un poil plus doué aurait pu mettre en péril l’existence d’une ou plusieurs activités, vous avez entre les mains un atout pour redéfinir la feuille de route. Si votre direction a eu peur, que vous n’avez pas cherché à l’entretenir, vous serez écouté.

Faut-il pour autant tout tracer ? 

Ce sera l’objet d’un autre article, celui-ci est déjà trop long.

Les lecteurs scrupuleux auront noté qu’il manque des sections par rapport à la liste des étapes. Il ne me paraissait pas utile de nous y attarder (et oui, cette phrase est syntaxiquement correcte ;) ).

Vogue la galère ! comme disaient les anciens.


[discord]: http://discord.comptoirsecu.fr 
[apio365]: https://www.crowdstrike.com/blog/hiding-in-plain-sight-using-the-office-365-activities-api-to-investigate-business-email-compromises/
