---

title: "Sécurité et full disclosure, doit-on tout dévoiler ?"
date: 2014-01-23T20:30:58+00:00
author: morgan


aliases: /2014/01/securite-et-full-disclosure-doit-on-tout-devoiler/
views: 1029
image: /images/2014/01/snapchat.jpg
categories:
  - Article
tags:
  - full disclosure
  - responsible disclosure
  - snapchat
---
> Quand la communication responsable ne fonctionne plus, un cas d’école avec SnapChat

<ins datetime="2014-01-23T21:58:37+00:00">Cet article était précédemment édité sur [Medium](https://medium.com/le-comptoir-secu/49473ef0a475), j’ai profité de cet import pour mettre à jour certaines sections.</ins>

## Avant propos : Qu’est-ce que Snapchat

Ce billet a été motivé par les dernières actualités sur [Snapchat](http://www.snapchat.com/). Pour ceux qui ne connaissent pas, Snapchat est la dernière application pour smartphone populaire auprès de nos petites têtes blondes. L’application permet d’envoyer à ses amis une photo ou une vidéo de moins de 10 secondes accompagnée si désiré d’un texte en guise de commentaire. Jusque-là on se demande la différence avec un simple logiciel de chat ou Instagram ou un Vine. Tout réside dans un petit détail : la photo/vidéo n’est visible qu’une fois par le destinataire, après quoi elle est détruite au lieu d’être stockée à vie comme le ferait Facebook ou Internet de façon générale.

![snapchat-pres](/images/2014/01/snapchat.png)

Ce semblant de vie privée fait fureur auprès des jeunes, notamment pour les choses que l’on risque de regretter avec l’âge : sexto avec une photo coquine, vidéo d’un ami ivre mort avec une forme phallique immortalisée au marqueur sur le front….

Je dis bien « un semblant de vie privée », car il faut pour cela faire une confiance aveugle dans la gestion des médias par Snapchat.

  1. Sont-ils vraiment détruits ? L’ajout récent d’une fonctionnalité permettant de [rejouer un snapchat une fois par jour](http://gigaom.com/2013/12/23/snapchat-update-lets-you-replay-one-snap-every-24-hours/) ne fait rien pour nous rassurer
  2. Il est bien sûr possible de prendre une capture d’écran de l’image et de la distribuer ailleurs. Le fait que l’application prévienne l’envoyeur qu’une capture a été faite ne change rien au problème, le mal est fait
  3. On peut récupérer le média sans que l’envoyeur ne le sache avec un peu de malice. Il suffit d’accéder au système de fichier et de récupérer le média avant son ouverture dans Snapchat.

Ce dernier point vous semble compliqué pour l’utilisateur moyen ? Ne vous en faites pas, des applications [rendent cela trivial](https://play.google.com/store/apps/details?id=com.appztastic.snapchatsavepics&hl=fr).

Bref, l’application est tellement populaire qu’elle a reçu des offres de rachat de p[lusieurs milliards de dollars](http://www.lepoint.fr/chroniqueurs-du-point/guerric-poncet/pour-snapchat-3-milliards-de-dollars-ne-suffisent-pas-14-11-2013-1756906_506.php) par Facebook et Google.

### L'affaire Snapchat

Comme toute application de startup qui se respecte, le développement a été fait le plus rapidement possible et n’a sûrement pas bénéficié d’une revue sécurité par des spécialistes.

Un groupe d’étudiants passionnés, réunis sous le nom [GibsonSec](http://gibsonsec.org/), a décortiqué l’API de l’application. Ils ont trouvé de nombreuses faiblesses et les ont gracieusement fournies en privé à Snapchat pour qu’ils puissent les corriger. Ils sont même allés jusqu’à postuler à une offre d’emploi de la Startup pour les aider à corriger leurs vulnérabilités.

_API signifie « Application Programming Interface ». C’est un peu le dialecte qui permet aux clients (l’application SnapChat) de communiquer avec les serveurs qui hébergent et redistribuent les images, enregistrent les comptes utilisateurs, etc…_

Ça, c’était début juillet. Les retours de Snapchat n’étaient pas encourageant, l’entreprise qualifiant les attaques de « théoriques ». En août, GibsonSec commence à s’agacer et fournit un premier avertissement par le biais d’ une [communication publique](http://gibsonsec.org/snapchat/index.html). Cette communication met en évidence les vulnérabilités sans fournir suffisamment d’informations pour les rendre directement exploitables.

![gibsonsec](/images/2014/01/gibsonsec.png)

Quatre mois plus tard, toujours aucune réaction de la part de Snapchat. GibsonSec craque et passe à ce qu’on appelle le « [Snapchat a essayé de « noyer le poisson »](http://gibsonsec.org/snapchat/fulldisclosure/" >full disclosure</a> ». Cela consiste à fournir au grand public l’ensemble des informations nécessaires pour exploiter les vulnérabilités découvertes. Comme cette fois-ci l’information a fait un peu de bruit sur la toile, <a href="http://blog.snapchat.com/post/71353347590/finding-friends-with-phone-numbers) peu de temps après, sous-entendu que le problème était pris en compte.

Il faut croire que les mesures prises étaient totalement insuffisantes, car pour fêter la nouvelle année [snapchatdb.info](http://www.theverge.com/2014/1/1/5262740/4-6-million-snapchat-phone-numbers-and-usernames-leaked" >un groupe d’individus s’est servi des informations fournies par GibsonSec</a> et a mis en ligne sur le site <a href="snapchatdb.info) une base de données de 4.6 millions de noms de comptes Snapchat associés au numéro de téléphone de l’utilisateur. Si vous êtes un utilisateur de Snapchat, ne vous sentez pas rassuré par le fait que les deux derniers chiffres du numéro de téléphone ne soient pas visibles. À ma connaissance la faille est toujours exploitable et par conséquent n’importe qui peut obtenir les mêmes résultats, sans censure. Ils pourront même obtenir l’information sans se fatiguer, il suffit de demander aux auteurs de Snapchatdb, qui fournissent volontiers la base de données non modifiée, ils l’ont d’ailleurs déjà communiquée à plusieurs reprises.

_Le site Snapchatdb n’est plus disponible, aucun ordre de justice n’a été donné pour sa fermeture. Cependant, l’hébergeur semble avoir pris peur et a fait justice lui-même._

> “Our motivation behind the release was to raise the public awareness around the issue, and also put public pressure on Snapchat to get this exploit fixed,” they say. « Security matters as much as user experience does. »

![snapchadb-exploit](/images/2014/01/snapchadb-exploit.jpg)

Les responsables de cette diffusion, ou en tout cas un groupe de personnes prétendant être les auteurs, [Snapchat stories](http://www.theverge.com/2014/1/1/5263156/alleged-snapchat-hackers-explain-how-and-why-they-leaked-data-on-accounts" >ont communiqué auprès des médias</a> les raisons de cette divulgation. Pour faire court, la diffusion des résultats de cette exploitation était la dernière étape possible pour alerter le grand public et secouer Snapchat, qui depuis des mois se moque royalement des trous béants dans son protocole, et par conséquent de la vie privée de ses utilisateurs. La société préfère ajouter de nouvelles fonctionnalités comme le « <a href="http://blog.snapchat.com/post/62975810329/surprise) ».

Snapchat a enfin réagi, [demander aux experts sécurité de les contacter à l’avenir](http://valleywag.gawker.com/snapchat-ceo-wont-say-sorry-for-hack-1494117437" >confirmé que ça leur passait au dessus</a>, et [promis une mise à jour](http://www.pcinpact.com/news/85188-snapchat-repond-a-attaque-et-promet-mise-a-jour-ses-applications.htm) qui ne semble pas corriger le cœur du problème. Le tout en ayant le culot de <a href="http://www.numerama.com/magazine/27944-snapchat-intervient-apres-une-fuite-touchant-46-millions-d-usagers.html).

<ins datetime="2014-01-23T19:38:09+00:00">Mise à jour du 23 janvier 2014 : Les dirigeants de SnapChat ont enfin compris leur totale absence de talent concernant la communication et [fournir un programme répondant tout seul au test](http://thehill.com/blogs/hillicon-valley/technology/194481-snapchat-hires-first-lobbyist" >ont embauché des lobbyists</a>. Malheureusement, ils n’ont toujours pas embauché de gens compétents en sécurité. La dernière mise à jour en date pour contrer les spammeurs, demande de [trouver les images avec un fantôme pour prouver que l’on est un humain](http://techcrunch.com/2014/01/21/snaptcha/). Grosse erreur, ce type d’exercice, le pattern matching, est extrêmement facile à faire pour un ordinateur, il a fallu moins d’une heure à un quidam sur internet pour <a href="http://gizmodo.com/how-i-hacked-snapchats-dumb-anti-robot-security-in-les-1506890048). C’est pitoyable.</ins>

## Le full disclosure : une approche critiquée

Comme à chaque affaire impliquant un full disclosure, les avis sont fermement partagés entre les partisans et ceux qui préfèrent le « coordinated disclosure », également appelé « responsible disclosure ». En français dans le texte : communiquer la vulnérabilité en privé à l’éditeur, attendre sagement que celui-ci prenne les mesures nécessaires et, seulement lorsque celui-ci donne son accord, publier l’information auprès du public. Le plus souvent, la communication sera faite en concert avec l’éditeur en question, rappelant que tout ceci n’est bien sûr plus que du passé et qu’il n’y a pas lieu de s’inquiéter.

Cette approche peut paraître idéale, car l’information sur la vulnérabilité n’est accessible que lorsqu’elle n’est plus exploitable. Bien qu’elle puisse fonctionner sans accroc avec des éditeurs responsables, les situations sont souvent moins roses.

### Quand l’éditeur maîtrise la communication, il n’est pas pressé par le temps

![Sleepy-at-work](/images/2014/01/Sleepy-at-work.jpg)

> À quoi bon se presser et débloquer un budget en urgence ? Le public n’est pas au courant, nous aurons tout le temps d’allouer des ressources au budget de l’année prochaine.

J’exagère le trait, mais l’idée est là. À part dans certains secteurs dont la sécurité est le cœur de métier, la correction des vulnérabilités est souvent le cadet des préoccupations des dirigeants. La priorité est à l’ajout de fonctionnalités demandées par l’utilisateur ou qui optimisent les bénéfices pour faire plaisir aux investisseurs. Sans pression, que ce soit au niveau du public avec la réputation, ou au niveau du porte-monnaie avec la justice et les régulateurs, les corrections traînent, parfois elles ne viennent jamais.

Ce n’est pas pour rien que les États mettent souvent en place des audits obligatoires, avec obligation d’amélioration, sous peine de lourdes sanctions financières ou pénales auprès des dirigeants. Cela pour forcer des secteurs comme le monde financier ou l’industrie (OIV) à maintenir un niveau de sécurité convenable. Sans cela la situation serait encore bien pire qu’elle ne l’est aujourd’hui.

Une fois l’information publiée, l’image de l’éditeur en pâtit et chaque jour qui passe sans correction augmente la chance de voir son nom dans les journaux associé à une grosse fuite d’information. C’est malheureusement parfois le seul moyen pour que l’entreprise se décide à traiter la vulnérabilité comme elle le mérite : une urgence.

### Certains éditeurs prennent le problème à l’envers

![Legal-scales-books-gavel-Image](/images/2014/01/Legal-scales-books-gavel-Image.jpg)

> Pourquoi perdre du temps et de l’argent dans une correction quand une simple lettre rédigée en 15 min par notre département juridique menaçant de poursuites judiciaires en cas de divulgation suffit à étouffer l’affaire ?

Ce genre de situation est déjà arrivé par le passé. Des professionnels ou passionnés de la sécurité découvrent une vulnérabilité « suivent les règles » en contactant l’entreprise concernée…et se font attaquer en justice.

Tout ceci m’écœure, rappelons les faits :

  1. Un expert en sécurité audite de façon bénévole le produit
  2. Il trouve des vulnérabilités et les fournit gracieusement à l’éditeur
  3. Dans la plupart des cas, l’expert applique même ce qu’on appelle le « responsible disclosure » et ne fait aucune communication publique avant correction
  4. …l’éditeur l’attaque en justice.

Si vous avez écouté notre [ponts d’or](https://www.comptoirsecu.fr/2013/06/podcast-episode-2-les-failles-0-day/" >épisode 2 du ComptoirSécu parlant des failles 0-day</a>, vous savez à quel point ce genre de découverte peut valoir cher sur le marché noir, ou comme des sociétés telles que Vupen sont prêtes à vous faire des <a href="http://www.vupen.com/english/careers/) pour vous intégrer dans ses rangs. Malgré tous, certains individus avec encore un peu d’éthique sont prêts à vous fournir tout ça gratuitement, même sans aucune reconnaissance pour leur travail, et ils se font attaquer.

Il ne faut pas s’étonner après ce genre d’affaires de voir de plus en plus d’experts [Google commence à serrer la vis](https://www.schneier.com/essay-146.html" >prôner le full disclosure</a>, souvent sous couvert d’anonymat pour éviter tout retour de flamme. Bigre, même <a href="http://googleonlinesecurity.blogspot.co.uk/2013/05/disclosure-timeline-for-vulnerabilities.html) et applique le full disclosure si après 7 jours l’éditeur n’a toujours pas corrigé une faille qu’ils estiment critique ou déjà utilisée.

### Alors, full ou responsible disclosure ?

![business turning back time](/images/2014/01/time-management1.jpg)

Je suis personnellement pour une version hybride, [similaire à celle que préconise Google](http://thehackernews.com/2013/05/google-sets-7-day-deadline-for_30.html). Je trouve dommage d’offrir en pâture aux médias une entreprise qui serait prête à jouer le jeu et traiter la correction de la vulnérabilité dans les délais qu’elle mérite. Ainsi une divulgation en privé auprès de l’éditeur me semble une bonne approche.

Si l’entreprise s’implique et communique en bonne intelligence, le responsible disclosure peut rester une solution viable.

Le full disclosure resterait donc un « plan B », dernière arme à utiliser quand l’éditeur :

  1. Est connu pour son hostilité envers ce genre de requête.
  2. N’offre aucun moyen de contact adéquat
  3. Ne répond pas à nos sollicitations
  4. Refuse de traiter la vulnérabilité en accord avec son niveau de criticité

## Le traitement des « responsible disclosure » par une entreprise

Si j’avais carte blanche dans une entreprise sur le traitement des failles de sécurité découvertes par un tiers, voilà ce que je préconiserai :

### Faciliter la prise de contact

![standard-telephonique-sacastar](/images/2014/01/standard-telephonique-sacastar.jpg)

Parfois, même avec toute la bonne volonté du monde, il est tellement difficile d’obtenir un interlocuteur compétent, voir un interlocuteur tout court, [que les experts jettent l’éponge et s’en remettent au full disclosure](http://qnrq.se/full-disclosure-american-express/).

Il est, je pense, important :

  1. De mettre en place un **moyen de contact dédié** aux évènements de sécurité
  2. Une adresse email à destination des équipes sécurité IT en interne (ou du responsable de la sécurité de l’entreprise) est amplement suffisante
  3. Fournir quelque part sur le site institutionnel une référence vers ce moyen de contact
  4. Former l’équipe responsable du support traditionnel à rediriger les sollicitations sur ce point de contact

### Disposer d’un processus de mise à jour d’urgence

Une vulnérabilité exploitée peut avoir des conséquences catastrophiques pour l’entreprise, ne serait-ce que pour l’image de marque. Beaucoup se souviennent encore des fuites d’[Sony](http://krebsonsecurity.com/2013/10/adobe-breach-impacted-at-least-38-million-users/" >Adobe</a> ou de <a href="http://krebsonsecurity.com/2011/04/millions-of-passwords-credit-card-numbers-at-risk-in-breach-of-sony-playstation-network/).

À partir du moment où une vulnérabilité est découverte, on peut supposer qu’elle est déjà connue et peut être même exploitée par d’autres individus. Il faut donc réagir dans les plus brefs délais.

![red-button](/images/2014/01/siteimg_large_276.jpg)

Une bonne pratique consiste à mettre en place une politique de patching. Cette politique définit différents niveaux de criticité pour les patchs, avec pour chaque niveau un délai de traitement maximal. Pour les patchs les plus critiques, on peut alléger les processus de tests et réduire le temps nécessaire à sa mise en production. On peut par exemple fixer un délai de 6 mois pour les patchs les moins sensibles, et 7 jours pour les plus critiques.

### Entretenir la relation avec les experts sécurités

Snapchat est un parfait exemple, l’équipe GibsonSec a fait son possible pour aider la startup, mais celle-ci n’a pas su leur donner confiance dans leur capacité à comprendre le potentiel des failles découvertes et à les corriger.

Mettre les experts en relation avec du personnel adéquat en interne est la première chose à faire. Ce genre d’échange ne peut pas se faire avec un responsable communication ou un support client standard. Il faut une personne compétente en sécurité informatique, avec un pouvoir de décision ou au moins des capacités d’escalades appropriées. Il est, je pense, important de garder contact avec l’expert et de le tenir informé de l’avancement sur la correction des vulnérabilités. Certaines entreprises comme [Google ](http://technet.microsoft.com/fr-FR/security/dn425036" >Microsoft</a>ou <a href="https://www.google.com/about/appsecurity/reward-program/)vont même jusqu’à offrir des récompenses pour ce genre de services, ce qui est selon moi une très bonne approche.

![shaking_hands1](/images/2014/01/shaking_hands1.jpg)

**Le responsible disclosure est un véritable cadeau fait aux entreprises. La conscience et l’éthique ne devraient pas être la seule motivation disponible, la reconnaissance, financière ou publique, devrait à minima être de la partie.

Sans cette reconnaissance, et avec l’importance croissante des outils de « cyber guerre », cet acte de charité se fera de plus en plus rare. L’éthique sera sacrifiée au profit d’une vente juteuse du 0-day en question, ou d’un débauchage par des professionnels du milieu.</strong>
