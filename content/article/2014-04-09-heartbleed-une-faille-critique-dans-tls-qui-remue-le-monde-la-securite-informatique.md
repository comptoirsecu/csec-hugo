---

title: "Heartbleed: une faille critique dans OpenSSL qui remue le monde de la sécurité informatique"
date: 2014-04-09T13:44:55+00:00
publisher: justin


aliases: /2014/04/heartbleed-une-faille-critique-dans-tls-qui-remue-le-monde-la-securite-informatique/
views: 1749
image: /images/covers/2014-04-heartbleed_by_gamershy-d6w2bnr.jpg
categories:
  - Article
---
> A missing bounds check in the handling of the TLS heartbeat extension can be
used to reveal up to 64k of memory to a connected client or server.

Voici ce qu’ont pu lire ceux qui suivent l'actualité sécu sur le site d'[OpenSSL.org](http://www.openssl.org) hier matin.

La faille joliment appelée "[heartbleed](http://heartbleed.com/)" à cause de la fonctionnalité de heartbeat d'OpenSSL est présente dans les bibliothèques d'OpenSSL 1.0.1 jusqu'à 1.0.1f et permet de déchiffrer les échanges entre le client et le serveur utilisant SSL/TLS. Elle permet de compromettre la clé privée du serveur utilisée pour chiffrer le trafic. Ainsi, il est possible, par exemple, de récupérer le login et le mot de passe de services ou d'utilisateurs, afin d’usurper leurs identités et d'infiltrer le serveur sans laisser de traces.

Cette vulnérabilité a été découverte par un groupe de chercheurs de [Codenomicon](http://www.codenomicon.com/), avec l’aide d’une personne de Google Security. D'après [Netcraft](http://news.netcraft.com/archives/2014/04/08/half-a-million-widely-trusted-websites-vulnerable-to-heartbleed-bug.html), plus de 500 000 serveurs seraient impactés.

{{<video "https://player.vimeo.com/video/91425662?title=0&byline=0&portrait=0" >}}


Comme souvent dans les failles liées au chiffrement, ce n'est pas l'algorithme en lui-même qui est faillible mais l’implémentation qui en est faite. Une erreur de programmation dans la bibliothèque permet de révéler la clé privée du certificat X509 du serveur. Bien heureusement, un patch est déjà disponible. Cependant, il faudra régénérer l'intégralité des certificats au cas où une fuite a déjà eu lieu (rappelez-vous, l'attaque ne laisse pas de traces). De plus, des comptes techniques ou utilisateurs ont pu être dévoilé. Il faudra ainsi changer les mots de passe ainsi que les éventuels cookies des utilisateurs. On peut d'ailleurs s’étonner sur le fait que très peu de sites Web ont réinitialisé les passwords de leurs utilisateurs...

Cette vulnérabilité fait beaucoup bouger le monde de la sécurité informatique car non seulement elle est très critique mais en plus demande un coût humain important pour la combler (régénération des mots de passes et des certificats).

Comment les administrateurs vont devoir combler cette faille ?

  * Tout d'abord, il faut mettre à jour la bibliothèque soit en passant en version en 1.0.1g ou bien en recompilant votre version d’OpenSSL avec le flag OPENSSL_NO_HEARTBEATS (Attention toutefois aux binaires qui disposent d'OpenSSL compilé en natif qui seront donc toujours vulnérable).
  * Ensuite, régénérer les certificats car il est impossible de savoir si ils ont été effectivement compromis.
  * Révoquer les cookies des utilisateurs
  * Encourager, voir forcer les utilisateurs à réinitialiser leurs mots de passe
  * Enfin, pour compléter le tout, faites un tour du côté du [Perfect Forward Secrecy](http://www.slashroot.in/what-perfect-forward-secrecy-and-its-impact-ssl-https) qui permet de garder l'intégrité de la clé de session même si la clé privée du serveur est compromise

Pour les utilisateurs, une procédure est envisageable :

  * Il faut tout d'abord tester si les sites sur lesquels vous avez un compte sont vulnérables. Pour tester la faille sur les serveurs, il existe un [site web](http://filippo.io/Heartbleed/#openssl.org) vous permettant de le vérifier en direct.
  * Si le serveur n'est plus vulnérable, il faut vérifier que les certificats ont été régénérés. Pour cela, regardez la date d’émission, si elle est antérieure au 8 avril, c'est mauvais signe.
  * Si et seulement si ces deux conditions sont validées, vous pouvez changer votre mot de passe, sinon cela ne sert à rien.
  * Enfin regardez du côté de la communication du site web pour savoir s’ils ont été gravement touchés. Par exemple, [Lastpass](http://blog.lastpass.com/2014/04/lastpass-and-heartbleed-bug.html) a publié un post sur son blog pour annoncer qu’aucun mot de passe situé sur leur plateforme ne pouvait avoir été compromis via cette faille
  * Faites aussi attention aux éventuels mails de phishing qui risquent d’arriver dans vos boites aux lettre ces jours prochains. C'est d'ailleurs l'occasion, si cela n'est déjà fait, de suivre [les bonnes pratiques de sécurité pour vos mots de passe](https://www.comptoirsecu.fr/2014/01/have-i-been-pwned-cachez-moi-ce-leak-que-je-ne-saurais-voir/), et pourquoi pas passer sur un gestionnaire de mots de passe! Il facilitera grandement votre processus de réinitialisation, en mettant des mots de passe robuste dont vous n'aurez pas à vous souvenir.

![heartbleed](/images/misc/2014-04-heartbleed.jpg)

Dernier point, comme le souligne [TroyHunt](http://www.troyhunt.com/2014/04/everything-you-need-to-know-about.html), assurez vous que votre navigateur vérifie si le certificat du site visité est encore actif. Si la révocation n'est pas vérifié, un certificat compromis par heartbleed pourra encore vous être servi lors d'une attaque en man in the middle dans plusieurs semaines sans que Chrome n'avertisse de quoi que ce soit...

![chrome_ssl](/images/misc/2014-04-chrome_ssl.jpg)

Évidemment, sur ce genre de vulnérabilité à grande échelle, des exploits sont déjà vendu sur le marché noir. C'est pour cela qu'il est très important de corriger cette faille au plus vite.

P.S: Pour rappel, nous avions parlé des différentes failles potentielles de TLS et de SSL dans notre [épisode pilote](https://www.comptoirsecu.fr/2013/05/podcast-episode-1-ssl-historique-faiblesses-et-avenir-2/).
