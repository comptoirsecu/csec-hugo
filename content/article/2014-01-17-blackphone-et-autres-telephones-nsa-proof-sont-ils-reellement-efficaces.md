---

title: "Blackphone et autres téléphones 'NSA-proof' : Sont-ils réellement efficaces ?"
date: 2014-01-17T09:30:19+00:00
publisher: morgan


aliases: /2014/01/blackphone-et-autres-telephones-nsa-proof-sont-ils-reellement-efficaces/
views: 2209
image: /images/covers/2014-01-blackphone.jpg
categories:
  - Article
tags:
  - Baseband
  - Blackphone
  - Cyanogen
  - NSA
  - SilentCircle
  - vie privée
---
Nous sommes depuis avant hier envahis par des articles sur le fameux [Blackphone](https://www.blackphone.ch), que ce soit sur [Korben](http://korben.info/le-blackphone-vraie-promesse-ou-future-deception.html), [01Net](http://www.01net.com/editorial/612046/blackphone-un-smartphone-anti-nsa-base-sur-android/), [Numerama](http://www.numerama.com/magazine/28074-blackphone-un-smartphone-open-source-dedie-a-la-vie-privee.html), [Techcrunch](http://techcrunch.com/2014/01/15/blackphone/)... Vous l’avez compris, la vie privée a le vent en poupe ces derniers temps grâce à Snowden. Un produit qui n’aurait intéressé que certains paranoïaques, les professionnels du renseignement et potentiellement le senior management est désormais vendu auprès du grand public.

Je ne vais pas trop m’attarder sur le Blackphone en lui-même, mais parler de ces téléphones sécurisés de façon générale. J'aimerai vous montrer à quel point il faut prendre des pincettes sur leur aspect « NSA-proof ».

Vous remarquerez que j’ai dit « ces téléphones ». Blackphone se vend comme le premier téléphone à mettre la protection de la vie privée au-dessus de tout. Disons que c’est une interprétation très singulière. Les smartphones sécurisés ne sont pas chose nouvelle. On vous a déjà parlé lors d'un Comptoirs Sécu du [Hoox](http://www.bull.fr/hoox/)  ou encore l'hilarant [Quasar IV](http://qsalpha.com/en/quasar-iv/) : premier smartphone adoptant les principes du Ninjutsu ! Certes ces téléphones parlent de confidentialité et non de vie privée, vous admettrez que l’un va rarement sans l’autre dans le cadre d'utilisation d'un particulier.

Je ne pourrai m’étendre sur les détails techniques du Blackphone. Tout simplement parce que nous n’avons aucune information concrète disponible pour le moment à ce sujet.

{{< toc >}}

# Vie privée dans les communications (appels, SMS)

Ici il ne faut pas s’attendre à beaucoup de surprises. Blackphone va probablement réutiliser le travail effectué par [SilentCircle ](https://silentcircle.com/web/how-it-works/). C’est tout naturel étant donné qu’ils participent à l’élaboration du produit.

SilentCircle est similaire à Whatsapp/Viber/GoogleVoice/Skype et consorts. En d’autres termes, vous n’utilisez pas le réseau cellulaire, mais votre forfait data pour les appels et messages textes. La valeur ajoutée des produits SilentCircle est qu’ils assurent un chiffrement de vos communications de bout en bout entre vous et le destinataire.

J’y vois déjà deux problèmes dans une utilisation courante voir exclusive :

  1. Il faut que l’interlocuteur soit aussi un utilisateur de l’application pour que cela fonctionne
  2. La qualité de l’appel dépend de votre connexion Internet et affecte votre quota mensuel

Bien entendu, il vous sera possible de contacter des personnes qui n’ont pas un Blackphone ou une application SilentCircle. Par contre le chiffrement s’arrêtera aux serveurs de SilentCircle, le reste de la connexion sera transmise via le réseau cellulaire traditionnel.

En clair. Il n’y aura pas de miracles sur ce point, pour ne pas être écoutable, il ne faut pas utiliser le réseau GSM. Cela limite très fortement le nombre de vos conversations qui seront effectivement protégées de la méchante NSA.

Parlons maintenant de ces fameux échanges sécurisés. Si le chiffrement de bout en bout permet de s’assurer que SilentCircle ne peut écouter les conversations qui transitent par ses serveurs, car elle n’a pas la clé, ils ont au moins des métadonnées.

Par métadonnées j’entends :

  1. qui vous appelez
  2. quand et à quelle heure
  3. pendant combien de temps.

Ce genre d’informations à elles seules peuvent révéler énormément d’informations sur vous. Je vous laisse regarder [quelques croustillants exemples](https://twitter.com/pinkflawd/status/417649515131596801/photo/1). Peut-être effacent-ils les métadonnées, je n’ai pas trouvé l’information sur le site de SilentCircle.

# Vie privée et applications

Je pars du principe que l’OS du Blackphone, PrivatOS, sera compatible avec le Play Store, sans cela c’est l’échec commercial assuré auprès du grand public.

Votre téléphone n’est plus un Nokia 3310, ils font bien plus que des appels et des SMS. Vous avez maintenant une multitude d’applications, souvent gratuites, qui adorent récupérer toutes les informations possibles et imaginables vous concernant. Je vous parle de Facebook, Google+, Instagram, Whatsapp, LinkedIn… à peu près tout ce que vous voulez.

Souvenez-vous de cet adage :

> **Si un produit est gratuit, VOUS êtes le produit.**

Ces applications raffolent de vos données, car elles les revendent aux publicitaires. Vous avez beau avoir un téléphone sécurisé, tout ce que vous mettrez, volontairement ou non, sur les serveurs de vos applications favorites, sera récoltable par les agences de renseignement.

Vous êtes prêt à vous séparer de toutes les applications susceptibles de collecter des données personnelles ? Parce que moi non.

La seule avancée que j’envisage de la part de Blackphone est de vous laisser le choix des permissions systèmes accordées aux applications. Lorsque vous téléchargez une application sur le Play Store, on vous impose toute une liste de permissions demandées par l’application. On pense par exemple à « Accéder au carnet d’adresses », « Voir l’identifiant de l’appelant » ou « Voir votre position GPS ».

Vous voulez utiliser l’application X, mais ne comprenez pas pourquoi elle veut voir les numéros de téléphone de tous vos amis et vous voulez lui refuser cette autorisation ? Impossible sur un Android non modifié, c’est tout ou rien. La possibilité de revoir les droits des applications est disponible sur des versions « modifiées » d’Android telles que [Cyanogen ](http://www.cyanogenmod.org/), je n’en attends pas moins de PrivatOS.

Pour ce qui est du surf, ils parlent d’un VPN, qui effectivement peut cacher votre IP, c’est l’IP du VPN qui se connecte au site Internet. On vous a déjà montré précédemment qu’il était possible d’identifier quelqu’un [par un tas d’autres données](https://panopticlick.eff.org/), et si vous vous connectez sur un compte nominatif quelconque, la confidentialité est à coup sûr perdue, même sur les autres sites qui ne manqueront pas de mettre un mouchard Facebook ou Google+ sur leur page…

Excusez-moi, on me dit dans mon oreillette qu’officiellement ça s’appelle un bouton « like » ou « +1 ».

# Vie privée et géolocalisation

Même si PrivOS a retiré tout mouchard éventuel de Google, que vous n’avez pas installé Google Maps ou que vous avez explicitement coupé la fonction GPS de votre téléphone, on peut encore vous suivre à la trace.

Votre téléphone, pour être joignable, cherche en permanence les antennes GSM à proximité et communique avec celles-ci en s’identifiant. Il en est de même pour votre connexion Edge/3G/4G. Ainsi, par le principe de [triangulation](http://fr.wikipedia.org/wiki/Triangulation), on peut facilement vous géolocaliser à quelques dizaines de mètres prêts.

Vous ne voulez pas que l’on sache ou vous êtes en permanence ? Ne prenez pas de smartphone.

# Vie privée et système d’exploitation

Sur ce point, il y a beaucoup de potentiel. Nous ne maîtrisons pas le système d’exploitation fourni dans nos téléphones, que ce soit Windows Phone, iOS, Android, BlackBerry ou autre. Vous n’êtes généralement pas administrateur de ce système et ne savez pas tout ce qu’il contient.Par conséquent, vous avez beau utiliser des applications sécurisées, si l’environnement d’exécution vous trahit, ça ne change pas grand-chose.

Laissez-moi m’essayer à une analogie :

> Alice veut envoyer un message à Bob, mais elle veut être sûre que personne d’autre ne puisse le lire, surtout pas, Mr NSA. Elle oublie donc d’emblée l’idée d’envoyer son message dans une simple lettre par la poste.
>
> Elle a défini un protocole ingénieux avec Bob. Ils ont chacun un lot de boites noires qui se ferment à clé. Chacun possède un exemplaire de la clé. Dès qu’Alice veut écrire à Bob, elle met son message dans la boite noire, la ferme à clé et envoie la boite à Bob sans la clé. Bob, qui possède un exemplaire de la clé, ouvre le coffre chez lui et lit le message.

Si l’on part du principe que la serrure est suffisamment forte pour être considéré inviolable, on est bon niveau confidentialité, n’est-ce pas ?

… Sauf si Mr NSA peut rentrer dans la maison d’Alice et Bob comme bon lui semble ! Il pourra mettre une caméra de surveillance qui filmera Alice en train d’écrire le message et Mr NSA pourra donc lire le contenu par-dessus l’épaule d’Alice qui le rédige. Il pourra aussi venir discrètement faire un double des clés, ouvrir la boite noire quand elle transite par la poste, lire le message, refermer la boite et la laisser arriver chez Bob. Les possibilités sont nombreuses.

Remplacez le protocole de la boite noire par le logiciel de SilentCircle et la maison par le système d’exploitation : même cas de figure.

Si Blackphone fait correctement son travail, audite scrupuleusement le code source d’Android, trouve toutes les backdoors potentielles et les retire du code, on pourra être davantage serein sur l’efficacité des communications sécurisées.

# Vie privée et matérielle, processeur baseband et carte SIM

Je ne vais pas me perdre dans les détails pour cette dernière partie, car je la maîtrise très peu. Sachez juste que votre téléphone ne se résume pas à son OS et ses applications. Android/iOS/etc.. [ne sont pas les seuls systèmes d’exploitation](http://www.extremetech.com/computing/170874-the-secret-second-operating-system-that-could-make-every-mobile-phone-insecure) de votre téléphone, il y a à minima :

  * Le firmware du processeur « baseband » qui se charge de toute la partie communication avec les antennes téléphoniques (appels, SMS, data)
  * Le processeur de votre carte SIM, celui-ci contient un micro-kernel pour accéder aux données chiffrées dans la puce.

Chaque système d’exploitation s’envoie des instructions en se faisant une confiance aveugle. Problème : ces micrologiciels sont fournis par les fabricants du téléphone. Ils sont livrés compilés et ne donnent pour ainsi dire jamais les sources.

Pour pallier à ce problème, Blackphone, plus précisément Geeksphone qui semble s’occuper de la partie matérielle, devrait utiliser un baseband open source. Cela permettrait de l’auditer et donc de vérifier l’absence de backdoors exploitables.

Apparemment il existe [au moins un baseband open source](http://bb.osmocom.org/trac/). Je doute fortement qu’il soit compatible avec le matériel de dernière génération qui sera utilisé dans le Blackphone.

D’ailleurs, dans la liste des téléphones compatible, il y a le fameux [OpenMoko](http://wiki.openmoko.org/wiki/Main_Page). Si vous avez de vrais geeks dans votre entourage, peut-être avez-vous déjà vu ~~cette bouse~~ ce bijou ? La dernière fois que j’en ai vu un, dès que son utilisateur mettait le téléphone en veille, le driver gérant le son ne fonctionnait plus et il fallait redémarrer l’appareil… merveilleux.

Imaginons une situation idyllique ou ce baseband serait open source et fonctionnel… reste le matériel. Certains experts soupçonnent la NSA d’avoir fait [implanter des backdoors dans les processeurs](http://wccftech.com/intel-possibly-amd-chips-permanent-backdoors-planted-nsa-updated-1/) Intel et AMD. Avec toutes les révélations de ces derniers mois, plus grand-chose ne nous semble impossible… Honnêtement sur cette partie je ne vois aucune façon d’obtenir le moindre réconfort. Il est sûr que Blackphone ne maîtrisera pas la fabrication du matériel et la sous-traitera en Asie comme tout le monde.

Un code source, ça s’audite et l’on peut le recompiler nous même pour être certains que la version binaire est bien celle que nous avons auditée. Pour auditer le matériel, il va falloir jouer avec l’oscilloscope et le microscope électronique, ce n’est déjà pas la même tisane. Pour ce qui est de refondre les pièces selon les plans, en partant du principe qu’ils soient fournis, je vous souhaite bien du courage.

# En fait le Blackphone ne sert à rien…

Je ne dis pas ça, et j’espère de tout cœur que nous aurons des avancées dans le domaine de la vie privée grâce à ce matériel. Même si tout n’est pas revu et 100 % fiable, une grosse partie du travail sera fait. Le téléphone restera peut-être vulnérable à une attaque ciblée, mais passera certainement entre les mailles du filet d’une grande partie de la surveillance de masse.

Bien sûr, tout cela ne dépendra pas que de l’appareil, mais de l’utilisation que vous en ferez.
