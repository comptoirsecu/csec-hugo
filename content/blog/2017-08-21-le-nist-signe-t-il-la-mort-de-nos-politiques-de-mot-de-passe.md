---
title: "Le NIST signe-t-il l’arrêt de mort de nos politiques de mot de passe ?"
authors:
  - jil
date: 2017-08-21
image: /images/covers/2017-08-26-password-policy-nist-medium.jpg
categories:
  - Article
tags:
  - NIST
  - Mot-de-passe
  - CNIL
---

# Qué qui dit ?

La nouvelle a défrayé la chronique : le dogme du mot de passe complexe de huit caractères s’effondre devant l’[analyse statistique de la Vérité][mesure]. Cela fait dix ans que l’on s’en doute et, pour reprendre la formule préférée des publicitaires : c’est prouvé scientifiquement. Plus besoin de se torturer l’esprit pour concevoir un mot de passe à la « 3v1L#P@ssword666 » et, au bout de trois mois, par un intense effort psychique, le transformer en « 3v1L#P@ssword667 ». Hourra !

Le NIST l’a annoncé, [Troy Hunt][troy] l’a relayé, mais vous trouverez un article très intéressant du côté de la [Federal Trade Commission][ftc], avec de nombreux renvois utiles. 

[![Dilbert 10 Sept 2005](http://assets.amuniversal.com/e47ff0606d5001301d7a001dd8b71c47)][dilbert]

# Ce que le NIST dit vraiment

Il n’est aucunement question de remettre en cause la robustesse d’un mot de passe aléatoire de huit caractères contenant des minuscules, majuscules, chiffres, caractères spéciaux (dont les frimousses Unicode font partie), même si l’on porte aujourd’hui ce nombre à douze caractères. Il s’agit simplement de rappeler que la méthode de construction des mots de passe *n’est pas* aléatoire, car notre cerveau fonctionne par association de mémoire. [Des modèles se dégagent][hydraze], avec une majuscule au début et les chiffres plutôt vers la fin.

Il est surtout question de repositionner cette mesure de sécurité (le mot de passe) dans le cadre [des menaces][menace] auxquelles le système d’information est exposé, ainsi que vis-à-vis des mesures complémentaires de protection.

Si la presse dessine une oasis de la fin de l’enfer des mots de passe, la réalité de la norme est plus terne. Le NIST en enlève d’un côté pour en ajouter de l’autre. On pourrait ainsi dire que :

1.	Les règles de composition se retirent au profit de l’emploi d’un composant graphique qui indique en temps réel la robustesse du mot de passe à l’utilisateur.

	Ce composant doit prendre en compte une liste noire de mots de passe basée sur les principaux mots de passe utilisés, sur les mots de passe faibles et probables dans le contexte de l’application, sur les mots de passe compromis. Si la composition de la liste noire est laissée à l’argumentation du détenteur du système, sa présence et le refus des mots de passe qu’elle contient sont obligatoires.

2. 	Le renouvellement périodique cède place au renouvellement lors de la suspicion de fraude. Il est donc nécessaire de s’équiper de systèmes qui détectent ce risque et provoquent le blocage du compte suspect ainsi que le renouvellement du mot de passe.

3.	Les sessions doivent expirer, le délai dépend de la sensibilité des informations (30 jours, 30 minutes, 15 minutes). Certaines expirations peuvent solliciter l’utilisateur deux minutes avant l’expiration pour qu’il confirme son activité et n’aie pas à se réauthentifier, d’autres sont systématiques, auquel cas elles ne doivent pas être inférieures à une heure.

4.	Le NIST rappelle aussi que les jetons (type OAuth) ne doivent pas être interprétés comme signalant la présence de l’utilisateur, donc qu’il faut les compléter par un autre moyen (deuxième facteur ou saisie du premier facteur) avant de donner accès à (ou de permettre des actions sur) des données sensibles.

Cette mise à jour du NIST vient s’ajouter aux recommandations du Royaume-Uni, où le [NCSC][ncsc] invite les entreprises à mettre à la disposition des utilisateurs des gestionnaires de mots de passe. C’est l’occasion d’avoir un secret différent par système et d’amoindrir les conséquences d’une compromission unitaire, qu’elle résulte d’une interception, d’une divulgation, d’une divination ou d’une inférence.

Il s’agit là d’une avancée considérable si l’on se place dans une stratégie de réduction du risque lié à l’usage des mots de passe. La vulnérabilité tirée de la paresse humaine est reportée sur l’accès au gestionnaire de mot de passe, donc notamment sur son exposition (en ligne ou hors-ligne), le secret qui le protège (clé, passphrase) et la robustesse du chiffrement utilisé. 

# Ce que la loi française permet

La CNIL a émis une [recommandation relative aux mots de passe][cnil] en janvier 2017, laquelle vient préciser l’obligation de protection des données personnelles visée par la loi. Autant dire qu’il vous faudra déployer de grands efforts pour justifier à un juge que vous y avez dérogé.

Dans l’hypothèse où le système d’authentification est pourvu d’un mécanisme d’anti-bruteforce, le mot de passe suivant est conforme à la réglementation : « J’suis trop fan du comptoir sécu ! » (vous noterez la vraie apostrophe inclinée, amateurs de [bépo](http://bepo.fr)). Vous avez donc l’obligation d’avoir trois types de caractères différents (minuscules, majuscules, chiffres, caractères spéciaux).

En ce qui concerne le renouvellement, « la commission recommande que le responsable de traitement veille à imposer un renouvellement du mot de passe selon une périodicité pertinente et raisonnable, qui dépend notamment de la complexité imposée du mot de passe, des données traitées et des risques auxquels il est exposé ».

Le « recommande » s’interprète ici comme le « SHOULD » des [RFC][rfc]. Je ne peux pas prononcer d’avis juridique qualifié, mais si vous disposez des systèmes de détection d’activités suspectes et de réinitialisation forcée, cela me semble conforme.

D’autres réglementations propres à votre domaine d’activité peuvent aussi entraver la disparition du renouvellement périodique des mots de passe.

# Le débat de la liste noire

Troy et ses 330 millions de mots de passe à blacklister se heurtent à l’utilisabilité d’un tel système (quel client accepterait de se faire refouler plusieurs mots de passe de suite au motif qu’ils ont fuité ailleurs ?) ainsi qu’à [la preuve du faible gain][blacklist] procuré par les listes noires (au-delà des valeurs les plus usitées, qui sont [obligatoires][nist]). Avec ironie, cette étude sur le faible intérêt est citée dans les sources du NIST, qui a tout de même laissé la possibilité d’interdire l’usage de millions de mots de passe. Certaines voies sont impénétrables.

Pour évaluer la robustesse d’un mot de passe issu du cerveau humain, l’on doit fournir à l’utilisateur une indication visuelle pour éclairer son choix. On pense alors à [zxcvbn][zxcvbn].

Cependant, pour en revenir à Troy, dans le cadre actuel où les mots de passe sont réutilisés entre les systèmes, savoir que son mot de passe a été découvert peut avoir un sens, dans la mesure où la personne concernée peut s’inquiéter de l’impact réel de la divulgation et prendre les mesures qu’elle estime appropriées. L’on préfèrera l’information à la contrainte.

# Le renouvellement périodique perd la bataille

Le renouvellement périodique présente l’intérêt d’invalider un mot de passe au bout d’un temps défini et prévisible. Si l’on prend la politique largement répandue d’un renouvellement trimestriel, cela laisse à l’attaquant un maximum de trois mois pour créer un compte secondaire sous sa maîtrise ou, plus simplement, pour installer un shell résilient qui bénéficiera de la session valide de l’utilisateur, quand bien même ce dernier aurait renouvelé son mot de passe. Sans avoir trouvé de base statistique du temps nécessaire à un pentester expérimenté pour installer un shell en userspace avec une évasion HTTPS lors d’une attaque ciblée, je pense pouvoir avancer que même un renouvellement périodique d’un mois serait vain. J’avais croisé une publication de Microsoft, hélas perdue depuis, qui insistait sur l’immédiateté de l’usage d’un mot de passe découvert par un attaquant.

C’est l’occasion de rappeler que le renouvellement périodique ne doit pas être confondu avec la mesure de désactivation des comptes inactifs.

Toutefois, avec les bases de mot de passe qui s’enfuient par millions dans les profondeurs du deep dark web pour y mener une vie trépidante en compagnie des bitcoins sous la cape de Tor (ou comment placer les mots-clés de référencement en une phrase maligne), le renouvellement périodique demeure une bonne idée.

# Protégé par le deuxième facteur

À toutes fins utiles, comme le rappelle [Alex Weinert][alex], le deuxième facteur n’est second que lorsque le premier existe. Il demeure vital que le mot de passe demeure secret et de le renouveler en cas de compromission suspectée ou avérée.

En outre, pour réduire l’impact de la divulgation, de l’inférence ou de la divination d’un mot de passe, il est conseillé de requérir un deuxième facteur d’authentification. 

Si vous disposez des capacités de détection qui vous permettent d’attribuer un niveau de risque à chaque authentification (nouveau terminal, accès via serveur mandataire, localisation inhabituelle, accès hors du réseau de l’entreprise, accès à une ressource critique, etc.), ce mécanisme un peu pénible peut être requis au-delà d’un niveau de risque pertinent. 

Sans cela, je vous déconseille de le rendre systématique sauf à disposer d’un jeton à clic, tel que l’envoi de la séquence par une pression sur une Yubikey (norme OATH) ou la validation d’une notification mobile du Microsoft Authenticator (ne fonctionne que pour les authentifications Azure).

[![Commitstrip password memory](http://www.commitstrip.com/wp-content/uploads/2015/08/Strip-Mots-de-passe-650-finalenglish1.jpg)][commitstrip]

# Je fais quoi avec tout ça ?

« J’y pense… et puis j’oublie… » Pas tout à fait, quand même. Nous faisons face à quelques difficultés :

1.	Il faudra expliquer aux utilisateurs que ce n’est pas la fête du slip pour autant, que la simplification s’accompagne d’un renforcement de la surveillance, donc une traçabilité soutenue de leur activité, ce qui n’est pas sans poser question selon le milieu (pensez aux salariés protégés, aux avocats, aux journalistes pour qui toute trace présente un risque).

2.	Il faudra mettre en œuvre la sécurité opérationnelle et s’assurer de son efficacité pour identifier et bloquer les activités suspectes, ce qui signifie que les attaquants vont se faire encore plus discrets (mais bon, aujourd’hui, on ne les voit déjà pas…).

3.	Il faudra jongler avec la législation, ou espérer qu’à l’occasion des détails d’interprétation du règlement général de protection des données (personnelles et privées), l’Union européenne s’aligne sur ce nouveau dogme.

4.	A-t-on les capacités techniques de forcer la réinitialisation d’un mot de passe dans l’ensemble des applications qui réalisent une authentification interne ? Évidemment, en lecteur assidu de l’[OWASP][owasp], cela fait belle lurette que c’est dans vos spécifications :D

5.	On parle des mots de passe des utilisateurs. Vous pouvez toujours planquer ceux des comptes à privilèges dans un bastion avec un mécanisme de renouvellement automatique à brève échéance, sans oublier de renouveler en temps utile les comptes de service (ceux qui sont planqués en clair dans les fichiers de configuration des applications).

Si d’aventure vous êtes client d’Office 365 et disposez de la licence EMS (E5), Azure Active Directory est votre allié, avec pour limite qu’il est aveugle à l’activité de l’ADFS (qui gère l’authentification) que vous avez pris soin de conserver chez vous sans synchroniser les condensats de mot de passe chez Microsoft. Un petit [RTFM][o365] vous fera le plus grand bien.

Bref, ce sujet pourrait mériter une note d’information de la part du RSSI.


[alex]: https://blogs.microsoft.com/microsoftsecure/2017/06/05/three-basic-security-hygiene-tips-from-microsofts-identity-team/
[blacklist]: https://www.internetsociety.org/sites/default/files/usec2017_01_3_Habib_paper.pdf
[cnil]: https://www.legifrance.gouv.fr/affichTexte.do?cidTexte=JORFTEXT000033928007
[commitstrip]: http://www.commitstrip.com/en/2015/08/19/password-memories/
[dilbert]: http://dilbert.com/strip/2005-09-10
[ftc]: https://www.ftc.gov/news-events/blogs/techftc/2016/03/time-rethink-mandatory-password-changes
[hydraze]: http://www.passwordresearch.com/stats/statistic376.html
[menace]: http://www.cerias.purdue.edu/site/blog/post/password-change-myths/
[mesure]: https://www.cs.unc.edu/~reiter/papers/2010/CCS.pdf
[ncsc]: https://www.ncsc.gov.uk/guidance/password-guidance-simplifying-your-approach
[nist]: https://pages.nist.gov/800-63-3/sp800-63b/sec5_authenticators.html#memsecretver
[o365]: https://docs.microsoft.com/en-us/azure/active-directory/active-directory-identityprotection
[owasp]: https://www.owasp.org/index.php/Password_Storage_Cheat_Sheet
[rfc]: https://www.ietf.org/rfc/rfc2119.txt
[troy]: https://www.troyhunt.com/passwords-evolved-authentication-guidance-for-the-modern-era/
[zxcvbn]: https://dl.dropboxusercontent.com/u/209/zxcvbn/test/index.html
