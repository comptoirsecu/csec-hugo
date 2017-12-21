---
title: "Comment participer sur le blog communautaire ?"
authors:
  - coink0in
date: 2017-12-20
image: /images/covers/2017-12-20-comment-participer-sur-le-blog-communautaire.jpg
categories:
  - Article
tags:
  - Debian
  - Nodejs
  - Git
  - Gulp
  - Yarn
  - Hugo
---



# Participer ? #



<center>*Oui ! N'hésitez pas à partager avec le reste de la communauté.*</center>

D'ailleurs, cet article est lui-même communautaire, vous êtes libre de le compléter, corriger, modifier avec vos propres expériences. L'idée générale étant de simplifier le processus, pour faciliter les contributions !

C'est grâce à la publication du code source, ainsi que la modification de la section article en blog communautaire, qu'il est maintenant possible de participer en publiant son propre article sur le Comptoir Sécu. Vous pouvez retrouver son principe et ses différentes motivations dans l'annonce de [l'ouverture du blog communautaire](/annonce/ouverture-du-blog-communautaire/).

Pour résumer, il est nécessaire, pour créer son propre article, de respecter quelques règles élémentaires et de faire une copie locale du site web.

<center>**Facile !** Mais parlons un peu de théorie, avant de détailler la pratique. </center>



# La théorie #



Le site web [comptoirsecu.fr](/) est un site dit [statique](https://fr.wikipedia.org/wiki/Page_web_statique), codé en [Go](https://fr.wikipedia.org/wiki/Go_(langage)), utilisant la syntaxe [Markdown](https://fr.wikipedia.org/wiki/Markdown) pour la mise en forme des publications. Il utilise le générateur de pages statiques [Hugo](https://gohugo.io/about/what-is-hugo/), ainsi que [Nodejs](https://nodejs.org/fr/about/) pour générer, concatener, optimiser  différents contenus comme les images, le  [JavaScript](https://fr.wikipedia.org/wiki/JavaScript) ou le [CSS](https://fr.wikipedia.org/wiki/Feuilles_de_style_en_cascade).

L’arborescence du site respecte celle décrite par Hugo [ici](https://gohugo.io/getting-started/directory-structure/), complétée par les fichiers de configuration des outils de développement, tels que le gestionnaire de paquets [Yarn](https://yarnpkg.com/en/docs) ou encore le module [Gulp](https://github.com/gulpjs/gulp/tree/master/docs).

L'ensemble étant disponible sur [le GitHub du Comptoir Sécu](https://github.com/comptoirsecu/csec-hugo) pour permettre de contribuer en apportant ses modifications, qui, une fois validées et approuvées, seront publiées.

<center>**Et vous voilà contributeur ! Merci**</center>



# La pratique #



Dans la pratique, il y a quelques paquets à installer et un environnement à récupérer avant de pouvoir réellement contribuer. Nous allons voir comment créer cet environnement virtuel, qui sera composé des éléments suivants :

 - **Nodejs** en version 8 ou +,
 - **Npm** en version 5 ou +,
 - **Hugo** en version 0.30 ou +,
 - **Yarn** en version 1.2 ou +,
 - **Gulp** en version 3.9.1 ou +,


Il est vivement conseillé de faire tout ça dans un environnement Unix/Linux ou sur MacOS. Pour cela vous pouvez
vous créer une machine virtuelle dédié *(VirtualBox, Workstation)* et préconfiguré, selon vos goûts. Sous windows il est aussi maintenant possible d'utiliser "Ubuntu" avec [quelques manipulations simples](http://www.eugenetoons.fr/utiliser-le-shell-bash-de-linux-sous-windows-10/). Si quelqu'un souhaite compléter cet article avec une partie *container* ou faire un retour sur les versions fonctionnelles des paquets, il est le bienvenu.



### Les dépendances ###



Pour fonctionner, les outils que nous allons utiliser nécessitent quelques dépendances, qu'il faut installer au préalable.

**Paquets à installer pour une distribution de type Debian/Ubuntu**
```Shell
apt install git-core curl build-essential openssl libssl-dev
```


### Nodejs & Npm ###



Pour rappel, Nodejs est un environnement d’exécution JavaScript qui a la particularité d'être asynchrone. \
Il dispose de son gestionnaire de paquet officiel Npm qui permet de télécharger et de partager des librairies.

Ces deux paquets sont disponibles dans les dépôts communautaires de la plupart des distributions et sont à jour. L'installation sur les différents types de systèmes d'exploitation est donc possible via leurs gestionnaires de paquets respectifs. La compilation depuis les sources est envisageable si vous avez du temps, beaucoup de temps ...

**Sur une distribution type Debian/Ubuntu**
```Shell
curl -sL https://deb.nodesource.com/setup_8.x | bash -
apt-get install -y nodejs
```
Vous pouvez retrouver les méthodes d'installation pour les autres types de systèmes [sur cette page](https://nodejs.org/en/download/package-manager/).

__Pour valider__ assurez-vous d'avoir les bonnes versions installées avec les commandes `npm -v ` et `node -e "console.log(process.versions)"`.



### Yarn & Gulp ###

Gulp et Yarn sont deux paquets également disponibles dans les dépôts communautaires du gestionnaire de paquet Npm. Toutefois selon certains retours, la méthode d'installation via Npm ne fonctionne pas toujours correctement *(L'informatique déterministe.. :karma:)*.

**Installation via Npm**
```Shell
npm install yarn -g
npm install gulp -g
```

Si vous rencontrez des difficultés lors de l'installation via Npm, d'autres méthodes d'installation sont décrites dans la documentation [ici pour Yarn](https://yarnpkg.com/en/docs/install#linux-tab). Il ne semble pas exister de méthode d'installation alternative pour Gulp. *(Si quelqu'un souhaite compléter cette partie.)*

__Vérifiez__ les versions que vous avez installées avec les commandes `yarn -v` et `gulp -v`.


### Hugo serveur ###

La dernière pièce du puzzle est donc Hugo qui va permettre de générer les pages statiques et qui devra être lancé en arrière-plan pour rendre le site accessible localement.

Hugo est disponible dans les dépôts, toutefois les méthodes d'installations détaillées dans la [documentation d'Hugo](https://gohugo.io/getting-started/installing/) ne permettent pas toujours d'installer une version 0.25 ou supérieur. C'est le cas, notamment, pour le dépôt stable de Debian Stretch, qui propose uniquement la version `hugo/stable 0.18.1-1+b2 amd64`.

La version `0.26-1` est disponible sur le gestionnaire de paquet dans les dépôts `testing` uniquement.

Hugo est un outil encore très jeune avec des mises à jour régulières qui pour la plupart ajoutent de nouvelles fonctionalités qui ne sont pas rétro-compatibles. À l'heure ou j'écris cet article, le site du Comptoir Sécu utilise Hugo 0.30. Le plus simple reste parfois de récupérer le binaire/paquet directement sur le [dépot GitHub des auteurs](https://github.com/gohugoio/hugo/releases).

Si vous tenez à utiliser votre gestionnaire de paquet, cela se passe comme suit :

**Sur une distribution type Debian/Ubuntu**

Ajoutez les dépôts Sid de Debian et configurez le gestionnaire de paquet APT, pour utiliser par défaut le dépôt Stable.

```Shell
echo -e "# Unstable (Sid)\ndeb ftp://ftp.fr.debian.org/debian/ testing main\n" >> /etc/apt/sources.list
echo -e "deb-src ftp://ftp.fr.debian.org/debian/ testing main\n" >> /etc/apt/sources.list
echo -e "deb http://security.debian.org/ testing/updates main\n" >> /etc/apt/sources.list
echo -e "APT::Default-Release "stable";" > /etc/apt/apt.conf.d/80default_stable
apt-get -t testing install hugo
```

Des méthodes d'installations pour d'autres distributions sont disponibles dans [la documentation d'Hugo](https://gohugo.io/getting-started/installing/#linux), mais attention à la version.

__Vérifiez__ la version avec la commande `hugo -v`.



### Clone & Run ! ###



Pour récupérer le code du Comptoir Sécu disponible sur GitHub, vous pouvez choisir de faire un [fork](https://help.github.com/articles/fork-a-repo/) vers votre compte GitHub, puis cloner depuis votre projet ou de cloner directement le projet depuis [comptoirsecu/csec-hugo](https://github.com/comptoirsecu/csec-hugo). Cependant, dans ce dernier cas de figure, vous ne pourrez pas [push](https://help.github.com/articles/pushing-to-a-remote/) vers les branches distantes.

Pour les allergiques des commandes Git, tout peut se faire en clicodrome avec leur application [Github Desktop](https://desktop.github.com/).

Placez-vous dans le répertoire avant de lancer un clone du Git de votre choix.

```Shell
cd /opt/
git clone https://github.com/comptoirsecu/csec-hugo.git && cd ./csec-hugo
```

Une fois le clone terminé, la première chose à faire est de récupérer toutes les dépendances du projet qui sont décrites dans le fichier `yarn.lock`. Heureusement, le gestionnaire de paquet Yarn va le faire automatiquement pour vous, lancez simplement la commande `yarn`.

<center>*C'est le moment de voir le résultat !*</center>

Ouvrez plusieurs terminaux, ou utilisez des utilitaires comme [screen](https://linux.die.net/man/1/screen) ou [tmux](https://linux.die.net/man/1/tmux) pour conserver un accès à la machine, car par défaut les processus Gulp et Hugo ne retournent pas l'invite de commande .

Dans un terminal :
```Shell
gulp
```

Dans un second terminal :

```Shell
hugo server --navigateToChanged --verbose
```

Si vous voulez personaliser le port, activer le logging des requêtes, etc. Je vous invite à regarder la sortie de la commande `hugo help`;).

<center>*Vous pouvez maintenant accéder au site du comptoirsecu.fr via votre navigateur préféré à l'adresse http://localhost:1313. Prenez le temps de vérifier que les journaux d'événements ne retournent pas d'erreur pendant la navigation, en sirotant une bière. C'est la pause ! ![pause_biere](/images/misc/2017-12-20-pause_biere.jpg)*</center>



# Le processus #



C'est maintenant que les choses sérieuses commencent ! Pour les mauvais élèves qui ont sauté la courte partie théorique, le projet que vous avez cloné, respecte la structure décrite par Hugo. Les répertoires qui vont nous intéresser pour publier un article sont les suivants :

 - **content** qui regroupe l'ensemble des sections, parmi lesquelles, `blog` avec tous les articles.
 - **data** pour les fichiers de configurations, notamment le fichier `authors` contenant les contributeurs.
 - **src** qui permet de placer les fichiers sources telles que les images avant transformation par Gulp.



### Créer votre profil ###



La première chose à faire est d'éditer le fichier de configuration `csec-hugo/data/authors.yml`, décrivant les contributeurs et d'ajouter votre profil. Les informations de base à compléter sont détaillées en dessous. Vous pouvez vous inspirer des autres profils, pour ajouter votre twitter ou site personnel, mais n'essayez pas le `staff: true` ça ne marche pas. :)

```yml
<votre_pseudo>:
  name: <votre_nom>
  description: <votre_description>
  image: <votre_pseudo>.jpg
```

__Notez__ que l'image de votre avatar étant générée et optimisée automatiquement par Gulp, il sera automatiquement redimensionné/retaillé et converti en JPG. La syntaxe du champ image __DOIT__ donc être au format .jpg car c'est le format final que verra le site.

 Le champs `votre_pseudo` quand à lui __DOIT__ être écrit en minuscule. Dites vous que `votre_pseudo`n'est qu'un alias vers votre fiche profile, il n'apparaitra nulle part, à part dans l'URL, donc tout en minuscule, pas d'espace ni de caractères spéciaux et tout se passera bien! Ce qui sera affiché sur le site dans votre profile sera le contenu du champ `<votre_nom>`, qui peut utiliser des majuscules, espaces, etc.




### Ajouter des images ###



Toutes les images que vous souhaitez ajouter doivent être placées dans un des sous-répertoires de `src/images` en fonction de l'endroit où elles doivent apparaitre.

 - **misc** qui regroupe les images contenues dans les articles.
 - **thumbnails** contient les images des avatars, des chansons et des bières !
 - **covers** pour les images de couverture des articles.

Par exemple, pour ajouter l'avatar décrit lors de la création de votre profil, placez l'image dans le répertoire `src/images/thumbnails` avec comme nom de fichier votre pseudo écrit en __minuscule__. \
Puis lancez la commande `gulp img` pour générer automatiquement les différents formats d'image. L'image est réencodée au format `.jpg`, redimensionnée et déplacée dans le répertoire *static/images/thumbnails/*. \
__Notez__ que l'image est réencodée, donc vous pouvez utiliser une image de base dans un autre format, comme svg, png, jpeg, etc ..

Si jamais vous n'êtes pas content du résultat, par exemple si le retaillage automatique est mal fait, et que vous souhaitez réessayer avec une nouvelle image, aucun souci, remetter une image avec le même nom dans `src/images/thumbnails` et re-faites `gulp img`, le résultat sera remplacé.

<center>*Vous avez maintenant un profil dans la section [contributeurs](/authors) !*</center>

Selon le même principe, vous pouvez ajouter l'image de couverture de votre article dans le répertoire `/src/images/covers/` avec un nom de fichier en minuscule, sans espace. Lors de la génération des images, une date est ajoutée au début du nom de fichier, sous le format `aaaa-mm-jj-` et l'extension modifiée en `.jpg`. L'image est déplacée dans `/static/images/covers/`.

Par conséquent, si vous effectuez la commande `gulp img` le 22 mars 2018 avec un fichier `/src/images/covers/partage.png`, vous aurez à la fin une image dans `/static/images/covers/2018-03-22-partage.jpg`

Dans le cas des images contenues dans les articles, une date est également ajoutée, avant qu'elles soient déplacées dans le répertoire `/static/images/misc/`.

Pourquoi rajouter des dates me direz-vous? Tout simplement pour que vous évitiez d'écraser les images des autres contributions qui porteraient le même nom!

La commande `gulp img` convertira toutes les nouvelles images d'un coup, mais si vous rajoutez des images par la suite, vous __devez__ relancer la commande `gulp img`.



### Écrire un article ###



Pour commencer à écrire un article, tapez la commande suivante, en remplaçant yyyy-mm-dd par la date de rédaction et mon-super-titre par le titre du billet en minuscule et avec des tirets à la place des espaces:

```Shell
hugo new blog/yyyy-mm-dd-mon-super-titre.md
```

Par exemple pour un billet intitulé "Mon premier article!" rédigé le 22 mars 2018 ça donnerait:

```Shell
$ hugo new blog/2018-03-22-mon-premier-article.md
/repos/private/csec/content/blog/2018-03-22-mon-premier-article.md created
$ cat content/blog/2018-03-22-mon-super-article.md
---
title: "Titre principal de l'article"
authors:
  - <pseudo_minuscule>
date: 2017-12-21
image:  /images/covers/2017-12-21-<nom_de_votre_image_covers>.jpg
categories:
  - Article
tags:
  - <TAG1>
  - <TAG2>
---

Contenu
```

Vous pouvez aussi  dupliquer l'un des fichiers `.md` existant dans le répertoire `content/blog/` en le renommant __avec__ la date et sans espace. C'est ce nom qui apparaitra dans l'URL du site et qui permettra d'identifier vos images plus facilement.

Vous devez modifier l'en-tête yaml du fichier dupliqué, qui contient toutes les informations relatives au contenu et qui a la structure suivante :

```yml
---
title: "Titre principal de l'article"
authors:
  - <pseudo_minuscule>
date: <aaaa-mm-jj>
image: /images/covers/aaaa-mm-jj-<nom_de_votre_image_covers>.jpg
categories:
  - Article
tags:
  - <tag1>
  - <tag2>
---
```

Vous pouvez commencer la rédaction de votre article à la suite. Après chaque modification du fichier sur le système de fichier, Hugo actualisera votre version locale du site.

Utilisez-la [syntaxe Markdown](https://www.markdownguide.org/basic-syntax) pour mettre en forme votre contenu et le processus d'ajout d'image pour l'enrichir. Une fois vos modifications enregistrées, votre profil de contributeur est automatiquement relié à l'article par le champ `authors:`.

__TIPS :__

 * Lorsque vous faites référence à vos images dans une publication utilisez un lien au format suivant : \
 `![Texte illustrant l'image](/images/misc/aaaa-mm-jj-<nom_fichier>.jpg)`.
 * Si vous avez choisi de faire un fork du projet csec-hugo, vous pouvez récupérer les modifications du projet d'origine dans votre fork en [configurant un dépôt distant](https://help.github.com/articles/configuring-a-remote-for-a-fork/#platform-linux) puis en [le synchronisant avec votre fork](https://help.github.com/articles/syncing-a-fork/#platform-linux). \
 Il est également possible de faire un webhook pour synchroniser les deux projets *(des infos?)*
 * Pensez à faire `gulp img` et à mettre les images dans `src/images/` pour qu’elles soient valides.



### Soumission & publication ###


Une fois que vous êtes satisfait de votre travail, vérifiez une dernière fois que les logs d'Hugo ne retournent pas d'erreurs et récupérez les dernières modifications du projet d'origine, avant de [créer une pull request](https://help.github.com/articles/creating-a-pull-request/). \
Vous devez au préalable, avoir validé vos modifications avec `git add .`, `git commit -m 'comment'` et push dans votre fork GitHub avec `git push origin master`.

Encore une fois, pour les allergiques des commandes Git, tout peut se faire en clicodrome avec leur application [Github Desktop](https://desktop.github.com/).

N'hésitez pas à passer sur [Discord](annonce/venez-échanger-avec-nous-sur-discord/) pour échanger avec les fondateurs sur le processus de validation.

<center>__À vous de jouer !__</center>
