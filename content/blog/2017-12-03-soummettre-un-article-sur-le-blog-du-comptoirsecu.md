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

# Comment participer ?

Grâce à la publication du code source, ainsi que de la modification de la section article en blog communautaire, il est maintenant possible de participer en publiant son propre article sur le ComptoirSécu. Retrouvez son principe et ses différentes motivations dans l'annonce de [ouverture du blog communautaire](https://www.comptoirsecu.fr/annonce/ouverture-du-blog-communautaire/).

Il donc nécessaire pour créer son propre article de respecter quelques règles élémentaires et de faire une copie local du site web. **Facile !** Mais détaillons un peu plus la théorie avant de passer à la pratique.


## La théorie

Le site web du [comptoirsecu.fr](https://www.comptoirsecu.fr/) est un site dit *statique*, codé en [Go](https://fr.wikipedia.org/wiki/Go_(langage)), utilisant la syntaxe [Markdown](https://fr.wikipedia.org/wiki/Markdown) pour la mise en forme des publications. Il utilise le générateur de page statiques [Hugo](https://gohugo.io/about/what-is-hugo/), ainsi que [Nodejs](https://nodejs.org/fr/about/) pour servir les différentes pages et le [Javascript](https://fr.wikipedia.org/wiki/JavaScript). 

L’arborescence du site respect celle décrite par Hugo [ici](https://gohugo.io/getting-started/directory-structure/), complétée par les fichiers de configuration des outils de développement, telle que le gestionnaire de paquets [Yarn](https://yarnpkg.com/en/docs) ou encore le module [Gulp](https://github.com/gulpjs/gulp/tree/master/docs).

L'ensemble étant disponible sur [le Github du comptoirsecu](https://github.com/comptoirsecu/csec-hugo) pour permettre de contribuer en apportant ses modifications, qui, une fois validées et approuvées, seront publiées.
 
**Et vous voila contributeur !**


## La pratique *(L'installation)*

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
```bash
apt install git-core curl build-essential openssl libssl-dev
```

### Step 1: Nodejs & npm ###

Pour rappel Nodejs est un environnement d’exécution JavaScript qui à la particularité d'être asynchrone. Il dispose de son gestionnaire de paquet officiel NPM qui permet de télécharger et de partager des librairies.

Ces deux paquets sont disponible dans les dépots communautaires et sont à jour. L'installation sur les différents types de système d'exploitation est donc possible via les gestionnaires de paquets. Il est également possible pour les plus barbus d'entre vous, de compiler depuis les sources.

**Distribution type Debian/Ubuntu**
```bash
curl -sL https://deb.nodesource.com/setup_8.x | bash -
 apt-get install -y nodejs
```
Vous pouvez retrouver les méthodes d'installation pour les autres type de système [sur cette page](https://nodejs.org/en/download/package-manager/).

__Pour valider__ assurez vous d'avoir les bonnes versions d'installées avec les commandes `npm -v ` et `node -e "console.log(process.versions)"`

### Step 2: Yarn & Gulp ###

Gulp et Yarn sont deux paquets également disponible dans les dépôt communauté du gestionnaire de paquet npm. Toutefois d'après certain retour la méthode l'installation via npm ne fonctionne pas toujours correctement. 

**Installation via npm**
```bash
npm install yarn -g 
npm install gulp -g 
```

Si vous rencontrez des difficultés lors de l'installation via npm, d'autres méthodes d'installation sont décrites dans les documentation [ici pour Yarn](https://yarnpkg.com/en/docs/install#linux-tab). Toutefois il ne semble pas exister de méthode d'installation alternative pour Gulp. *(Si quelqu'un souhaite compèter cette partie.)*

__Vérifiez__ là version que vous avez installée avec les commandes `yarn -v` et `gulp -v`.

### Step 3: Hugo ###

La dernière pièce du puzzle est donc Hugo qui va permetre de générer les pages statiques et qui devra être lancé en arrière plan pour rendre le site accessible localement.

Hugo est disponible dans les dépôts des principales distribution Linux, toutefois les méthodes d'installations détaillées dans la documentation d'Hugo ne permettent pas d'installer une version 0.25 ou supérieur. C'est le cas, notement, pour le dêpot stable de Debian Stretch, qui propose uniquement la version `hugo/stable 0.18.1-1+b2 amd64`.

Les methodes d'installations des autres distributions pour Hugo sont disponible dans [la documentation](https://gohugo.io/getting-started/installing/#linux). Pour Debian Stretch la version `0.26-1` est disponible à via le gestionnaire de paquets dans les depôts `testing`.

**Distribution type Debian/Ubuntu**

Ajoutez les depôts Sid de Debian et configurez le gestionnaire de paquet APT pour utiliser par defaut les depôts Stable.

```bash
echo -e "# Unstable (Sid)\ndeb ftp://ftp.fr.debian.org/debian/ testing main\n" >> /etc/apt/sources.list
echo -e "deb-src ftp://ftp.fr.debian.org/debian/ testing main\n" >> /etc/apt/sources.list
echo -e "deb http://security.debian.org/ testing/updates main\n" >> /etc/apt/sources.list
echo -e "APT::Default-Release "stable";" > /etc/apt/apt.conf.d/80default_stable
apt-get -t testing install hugo
```

__Vérifiez__ là version avec la commande `hugo -v`.

### Step 4: Clone & Run ! ###

Enfin rècupèrez le code du comptoirsecu sur le projet [comptoirsecu/csec-hugo](https://github.com/comptoirsecu/csec-hugo) disponible sur GitHub. Placez vous dans le repertoire de votre choix avant de lancer un clone du répertoire.

```bash
cd /opt/
git clone https://github.com/comptoirsecu/csec-hugo.git && cd ./csec-hugo
```

Une fois le clone terminé, la première chose à faire est de récupérer toutes les dependances du projet qui sont décrites dans le fichier `yarn.lock`. Heureusement le gestionnaire de paquet Yarn va le faire automatiquement pour vous, lancez simplement `yarn`. 

Il ne reste plus cas faire, enfin, run le site !

Ouvrez plusieurs shell, ou utilisez des utilitaires comme screen ou tmux pour conserver un accès à la machine car par defaut les processus gulp et hugo ne retourne pas le prompt.

```bash
screen -dsm gulp-realtime gulp
screen -dsm comptoirsecu-web hugo server --bind `hostname -I` --port 80 -b http://`hostname -I` —navigateToChanged --verboseLog --log --verbose
```

Pensez à modifier les directives `--bind` et `--port` celon votre environnement de développement.

*Vous pouvez maintenant accèder au site du comptoirsecu.fr via votre navigateur préféré à l'adresse utilisée dans la directive bind d'hugo. Prennez le temps de vérifier que les logs ne retourne pas d'erreur pendant la navigation, en sirrotant une bière, c'est la pause !*

## La suite *(Le processus)* ##

C'est maintenant que les choses sérieuses commences ! Pour les mauvaises éléves qui ont sautés la courte partie théorique, le projet que vous avez cloné respect la stucture décrite par Hugo. Les repertoires qui vont nous interesser pour publier un acticle sont les suivants :

 - **content** qui regroupe l'ensemble des differentes sections, parmis lesquelles `blog` avec tous les articles.
 - **data** qui est utilisé pour stocker les fichiers de configurations.
 - **src** qui permet de placer les fichiers sources avant transformation par Gulp.

### Step 5: C'est vous ! ###

La premère chose à faire est d'éditer le fichier de configuration des contributeurs `csec-hugo/data/authors.yml`et d'ajouter votre profile. Les informations de base à complèter sont détaillées en dessous. Vous pouvez vous inspirez des autres pour ajouter votre twitter ou site personnel, mais n'essayez pas le `staff: true` ça ne marche pas :)

```yml
<votre_pseudo>:
  name: <votre_nom>
  description: <votre_description>
  image: <votre_pseudo>.jpg
```