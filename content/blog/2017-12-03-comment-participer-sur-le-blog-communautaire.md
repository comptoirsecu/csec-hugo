---
title: "Comment participer sur le blog communautaire ?"
authors:
  - coink0in
date: 2017-12-03
image: /images/covers/2017-12-03-comment-participer-sur-le-blog-communautaire.jpg
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



# Participer ?



<center>*Oui ! N'hésitez à partager avec le reste de la communautée.*</center>

D'ailleur cet article est lui même communautaire donc sentez-vous libre de le complèter, corriger, modifier avec vos propres expériences. L'idée générale étant de simplifier le processus pour facilité les contributions !

C'est grâce à la publication du code source, ainsi que la modification de la section article en blog communautaire, qu'il est maintenant possible de participer en publiant son propre article sur le ComptoirSécu. Vous pouvez retrouver son principe et ses différentes motivations dans l'annonce de [l'ouverture du blog communautaire](/annonce/ouverture-du-blog-communautaire/).

Pour résumer il est nécessaire pour créer son propre article de respecter quelques règles élémentaires et de faire une copie local du site web.

<center>**Facile !** Mais détaillons un peu plus la théorie avant de passer à la pratique.</center>



# La théorie #



Le site web du [comptoirsecu.fr](/) est un site dit [statique](https://fr.wikipedia.org/wiki/Page_web_statique), codé en [Go](https://fr.wikipedia.org/wiki/Go_(langage)), utilisant la syntaxe [Markdown](https://fr.wikipedia.org/wiki/Markdown) pour la mise en forme des publications. Il utilise le générateur de pages statiques [Hugo](https://gohugo.io/about/what-is-hugo/), ainsi que [Nodejs](https://nodejs.org/fr/about/) pour servir les différentes pages et le [Javascript](https://fr.wikipedia.org/wiki/JavaScript). 

L’arborescence du site respect celle décrite par Hugo [ici](https://gohugo.io/getting-started/directory-structure/), complétée par les fichiers de configuration des outils de développement, telle que le gestionnaire de paquets [Yarn](https://yarnpkg.com/en/docs) ou encore le module [Gulp](https://github.com/gulpjs/gulp/tree/master/docs).

L'ensemble étant disponible sur [le Github du comptoirsecu](https://github.com/comptoirsecu/csec-hugo) pour permettre de contribuer en apportant ses modifications, qui, une fois validées et approuvées, seront publiées.
 
<center>**Et vous voila contributeur ! Merci**</center>



# La pratique #



Dans la pratique il y a quelques paquets à installer et un environnement à récupérer avant de pouvoir réellement contribuer. Nous allons voir comment créer cet environnement virtuel qui sera composé des éléments suivants :

 - **Nodejs** en version 8 ou +,
 - **npm** en version 5 ou +,
 - **Hugo** en version 0.25 ou +,
 - **Yarn** en version 1.2 ou +,
 - **Gulp** en version 3.9.1 ou +,

Commençons donc sur des bonnes bases avec un environnement de developpement virtuel dédié *(VirtualBox, Workstation)* et pré-configuré celons vos goûts. Si quelqu'un souhaite compléter cet article avec une partie *container* ou faire un retour sur les versions fonctionnelles des paquets. il est le bienvenue.



### Step 0: Les dépendances ###



Touts les outils que nous allons déployer et les actions que nous allons devoir faire pour dupliquer le site web, nécessites quelques dépendances que nous installée au préalable.

**Distribution type Debian/Ubuntu**
```shell
apt install git-core curl build-essential openssl libssl-dev
```


### Step 1: Nodejs & npm ###



Pour rappel Nodejs est un environnement d’exécution JavaScript qui à la particularité d'être asynchrone. Il dispose de son gestionnaire de paquet officiel Npm qui permet de télécharger et de partager des librairies.

Ces deux paquets sont disponible dans les dépots communautaires de la plus part des distributions et sont à jour. L'installation sur les différents types de système d'exploitation est donc possible via leurs gestionnaires de paquets. La compilation des sources est aussi envisageable si vous avez tu temps ...

**Distribution type Debian/Ubuntu**
```shell
curl -sL https://deb.nodesource.com/setup_8.x | bash -
apt-get install -y nodejs
```
Vous pouvez retrouver les méthodes d'installation pour les autres type de système [sur cette page](https://nodejs.org/en/download/package-manager/).

__Pour valider__ assurez vous d'avoir les bonnes versions d'installées avec les commandes `npm -v ` et `node -e "console.log(process.versions)"`.



### Step 2: Yarn & Gulp ###



Gulp et Yarn sont deux paquets également disponible dans les dépôts communautaire du gestionnaire de paquet npm. Toutefois d'après certain retour la méthode l'installation via npm ne fonctionne pas toujours correctement. 

**Installation via npm**
```shell
npm install yarn -g 
npm install gulp -g 
```

Si vous rencontrez des difficultés lors de l'installation via npm, d'autres méthodes d'installation sont décrites dans les documentation [ici pour Yarn](https://yarnpkg.com/en/docs/install#linux-tab). Toutefois il ne semble pas exister de méthode d'installation alternative pour Gulp. *(Si quelqu'un souhaite complèter cette partie.)*

__Vérifiez__ les versions que vous avez installée avec les commandes `yarn -v` et `gulp -v`.



### Step 3: Hugo ###



La dernière pièce du puzzle est donc Hugo qui va permetre de générer les pages statiques et qui devra être lancé en arrière plan pour rendre le site accessible localement.

Hugo est disponible dans les dépôts des principales distribution Linux, toutefois les méthodes d'installations détaillées dans la documentation d'Hugo ne permettent pas toujours d'installer une version 0.25 ou supérieur. C'est le cas, notement, pour le dêpot stable de Debian Stretch, qui propose uniquement la version `hugo/stable 0.18.1-1+b2 amd64`.

La version `0.26-1` est disponible à via le gestionnaire de paquet dans les depôts `testing` uniquement.

**Distribution type Debian/Ubuntu**

Ajoutez les depôts Sid de Debian et configurez le gestionnaire de paquet APT pour utiliser par defaut les depôts Stable.

```shell
echo -e "# Unstable (Sid)\ndeb ftp://ftp.fr.debian.org/debian/ testing main\n" >> /etc/apt/sources.list
echo -e "deb-src ftp://ftp.fr.debian.org/debian/ testing main\n" >> /etc/apt/sources.list
echo -e "deb http://security.debian.org/ testing/updates main\n" >> /etc/apt/sources.list
echo -e "APT::Default-Release "stable";" > /etc/apt/apt.conf.d/80default_stable
apt-get -t testing install hugo
```

Des methodes d'installations pour d'autres distributions sont disponible dans [la documentation d'Hugo](https://gohugo.io/getting-started/installing/#linux), mais attention à la version. 

__Vérifiez__ là version avec la commande `hugo -v`.



### Step 4: Clone & Run ! ###



Pour récupérer le code du comptoirsecu disponible sur GitHub, vous pouvez choisir de faire un [fork](https://help.github.com/articles/fork-a-repo/) vers votre compte GitHub ou de cloner directement le projet depuis [comptoirsecu/csec-hugo](https://github.com/comptoirsecu/csec-hugo). Cependant, dans ce dernier cas de figure, vous ne pourrez pas [push](https://help.github.com/articles/pushing-to-a-remote/) les branches distantes.

Placez vous dans le repertoire de votre choix avant de lancer un clone du Git de votre choix.

```shell
cd /opt/
git clone https://github.com/comptoirsecu/csec-hugo.git && cd ./csec-hugo
```

Une fois le clone terminé, la première chose à faire est de récupérer toutes les dependances du projet qui sont décrites dans le fichier `yarn.lock`. Heureusement le gestionnaire de paquet Yarn va le faire automatiquement pour vous, lancez simplement la commande `yarn`. 

<center>Il ne reste plus cas faire, enfin, run le site !</center>

Ouvrez plusieurs shell, ou utilisez des utilitaires comme [screen](https://linux.die.net/man/1/screen) ou [tmux](https://linux.die.net/man/1/tmux) pour conserver un accès à la machine, car par defaut les processus gulp et hugo ne retourne pas le prompt.

```shell
screen -dsm gulp-realtime gulp
screen -dsm comptoirsecu-web hugo server --bind `hostname -I` --port 80 -b http://`hostname -I` —navigateToChanged --verboseLog --log --verbose
```

Pensez à modifier les directives `--bind` et `--port` celon votre environnement de développement.

<center>*Vous pouvez maintenant accèder au site du comptoirsecu.fr via votre navigateur préféré à l'adresse utilisée dans la directive bind d'hugo. Prennez le temps de vérifier que les logs ne retourne pas d'erreur pendant la navigation, en sirrotant une bière, c'est la pause ! ![pause_biere](/images/misc/2017-12-15-pause_biere.jpg)*</center>



# Le processus #



C'est maintenant que les choses sérieuses commences ! Pour les mauvaises éléves qui ont sautés la courte partie théorique, le projet que vous avez cloné respect la stucture décrite par Hugo. Les repertoires qui vont nous interesser pour publier un acticle sont les suivants :

 - **content** qui regroupe l'ensemble des differentes sections, parmis lesquelles `blog` avec tous les articles.
 - **data** qui est utilisé pour stocker les fichiers de configurations.
 - **src** qui permet de placer les fichiers sources avant transformation par Gulp.



### Step 5: Créer votre profile ###



La première chose à faire est d'éditer le fichier de configuration des contributeurs `csec-hugo/data/authors.yml` et d'ajouter votre profile. Les informations de base à complètées sont détaillées en dessous. Vous pouvez vous inspirez des autres pour ajouter votre twitter ou site personnel, mais n'essayez pas le `staff: true` ça ne marche pas :)

```yml
<votre_pseudo>:
  name: <votre_nom>
  description: <votre_description>
  image: <votre_pseudo>.jpg
```

Notez que l'image de votre avatar étant générée et optimisée automatiquement par gulp, la systaxe du champ image __DOIT__ êtreau format .jpg et que votre pseudo __DOIT__ être écrit en minuscule. 



### Step 6: Ajouter des images ###



Toutes les images que vous souhaitez ajouter doivent être placées dans un des sous répertoires de `src/images` en fonction de l'endroit où elles doivent apparaitres.

 - **misc** qui regroupe les images contenues dans les articles.
 - **thumbnails** contient les images des avatars et des bières !
 - **covers** pour les images de couverture des articles.

Pour ajouter l'avatar décrit lors de la création de votre profile, placez l'image dans le répertoire `src/images/thumbnails` avec comme nom de fichier votre pseudo écrit en minuscule *(comme dans le yaml authors)*. \
Puis retournez sur votre terminal lancer la commande `gulp img` pour générer automatiquement les différents formats d'image. L'image est réencodée au format `.jpg`, redimenssionée et déplacée dans le répertoire *static/images/thumbnails/* avant d'être détectée par gulp que vous avez lancé en arrière plan précédément. __Notez__ que l'image est forcement réencodée vous pouvez donc mettre une image d'un autre format *(Et que c'est foutu pour la stéganographie)*.

<center>*Vous avez maintenant un profile dans la section [contributeurs](/authors)*</center>

Celon le même principe, ajouter l'image de couverture de votre article dans le répertoire `/src/images/covers/` avec un nom de fichier en minuscule, sans espace. Une date est alors ajoutée au debut du nom de fichier, sous le format `aaaa-mm-jj-` et l'extension modifée en `.jpg`. L'image est déplacée dans `/static/images/covers/`.

Dans le cas des images contenues dans les articles, une date est également ajoutée, avant qu'elles soient déplacées dans le répertoire `/static/images/misc/`.



### Step 7: Ecrire un article ###



Pour commencer à écrire un article dupliquez l'un des fichiers `.md` existant et en le renommant avec la date et sans espace. C'est ce nom qui apparaitra dans l'URL du site et qui permettra d'identifier vos images misc plus facilement.

L'en-tête du fichier dupliqué, écrite en yaml, contient toutes les informations rélitives à son contenu. Sous la structure suivante, que vous devez modifier.

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

Utilisez la syntaxe Markdown pour mettre en forme votre contenu et le processus d'ajout d'image pour l'enrichir. Une fois vos modifications enregistrées, votre profile de contributeur est automatiquement relié à l'article via le champ `authors:`.

__TIPS :__

 * Lorsque vous faites référence à vos images dans une publication utilisez un lien au format suivant : \
  `![text_si_problem_d'affichage](/images/misc/aaaa-mm-jj-<nom_fichier>.jpg)`.
 * Si vous avec choisi de faire un fork du projet csec-hugo, vous pouvez récupérer les modifications du projet d'origin dans votre fork en [configurant un dêpot distant](https://help.github.com/articles/configuring-a-remote-for-a-fork/#platform-linux) puis en [synchronisant les modifications dans votre fork](https://help.github.com/articles/syncing-a-fork/#platform-linux). Il est également possible de faire un webhook pour synchroniser les deux projets *(des infos?)*



### Step 8: Soumission & publication ###



Une fois satisfait de votre travail, vérifiez une dernière fois que le shell du site ne retourne pas d'erreur et récupérez les dernières modifications du projet d'origin, avant de [créer une pull request](https://help.github.com/articles/creating-a-pull-request/). \
Vous devez au préalable, avoir valider vos modifications avec `git add .`, `git commit -m 'comment'`, 