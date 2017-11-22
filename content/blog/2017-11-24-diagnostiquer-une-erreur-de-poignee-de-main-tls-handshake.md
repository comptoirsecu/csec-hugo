---
title: "Diagnostiquer une erreur lors de la poignée de main TLS"
authors:
  - jil
date: 2017-11-24
image: /images/covers/2017-08-26-password-policy-nist.jpg
categories:
  - Article
tags:
  - TLS
  - SSL
  - handshake
---

# Il était une fois…

… dans les couloirs d’une DSI somme toute classique. On y évoquait les fantômes, pour ne pas dire les légendes, de ces rares créatures qui figurent au registre du personnel et qu’on ne croise que rarement, tant elles courent de désastre en désastre pour sauver des projets dévastés par un retard devenu insupportable pour la haute direction. Faiseurs de miracles, ils aimeraient surtout que les pratiquants de l’informatique se renseignent et se forment sur la réalité technique induite par leurs formules magiques.

Cet article aura donc pour objet de résoudre une problématique précise, celle de la fameuse poignée de main TLS qui échoue. Pas de quoi casser trois pattes à un canard, mais qui pourrait sauver des heures de recherche au malheureux développeur.

# Le cas tordu

Nous sommes en présence d’une application FaisDesGlaçons (F), hébergée chez un mastodonte (A), avec une assistance technique aussi utile qu’un pot de yaourt. 

Cette application F doit appeler un service web chez nous (Z), PrendsUnVerreDeLiquideBrun (V). La chaîne de transmission est donc classique :

	[A::F] ↣ [A::fw] ↣ {internet} ↣ [Z::fw] ↣ [Z::lb] ↣ [Z::V]

Il y a donc deux pare-feu et un répartiteur de charge. C’est ce dernier qui assure la terminaison TLS, après quoi le flux est en clair vers V.

Ajoutez à cela un intégrateur censé paramétrer F, et qui va lire ce précieux article du comptoir sécu :D

# Jusqu’au problème de la poignée de main
	
En premier lieu, l’intégrateur vous aura demandé le port de votre service web, que vous lui aurez communiqué selon vos normes, et il vous aura informé que cela ne fonctionne pas. 

C’est bien évidemment de votre faute, vos équipes ne peuvent rivaliser avec la perfection des infrastructures de A. Vous obtenez donc de votre équipe réseau une capture d’écran des journaux du pare-feu qui démontre qu’aucune tentative de connexion n’est visible sur le port indiqué.

> Les gens du réseau et de la sécu ont, dans leur grande sagesse,
> la possibilité de laisser à l’architecte (responsable du chantier
> d’intégration d’un nouveau bousin) l’accès à une console qui lui
> permet de consulter les journaux adéquats du pare-feu.

Après avoir laissé chercher l’intégrateur, il aura trouvé, peut-être à l’aide de l’assistance, que le pare-feu de A ne laisse sortir que sur le range TCP/2340-2349.

Qu’à cela ne tienne, vous demandez à vos équipes réseau d’exposer V sur TCP/2345, ce qui est le travail du répartiteur et du pare-feu. La sécu comme le réseau a bien pesté, mais le port est ouvert.

Nouvel essai : pas mieux ! Cette fois-ci, vous voyez la connexion dans les journaux du pare-feu, et comme les gens du réseau travaillent bien, la connexion est autorisée. On vous fait aussi savoir que, oui, le répartiteur est bien configuré et que le service V est up.

En relisant le courriel de détresse de l’intégrateur, vous lui faites ajouter le « s » à « http », parce qu’on traverse internet, donc on prend ses précautions.

Hourra ! On arrive au cœur du problème, ce qui a déjà pris 10 jours : *TLS handshake error*

# Solutions

## L’application bien programmée

Elle remonte l’erreur exacte, et votre intégrateur vous la colle, ce qui vous permet de l’interpréter. Tous les langages indiquent la routine qui a planté. Par exemple, vous avez un problème de suite cryptographique si vous échouez à l’étape `GET_SERVER_HELLO` :

	error:14077410:SSL routines:SSL23_GET_SERVER_HELLO:sslv3 alert handshake failure

Mais bon, si c’était bien codé, ou si le message d’erreur était accessible, ça ne serait pas réaliste.

## Si le répartiteur le veut bien

Dans les journaux du répartiteur de charge, il devrait indique la raison de cette erreur. 

Mais voilà, comme c’est un système qui ronronne comme le moteur d’une 2CV, son admin est parti pêcher le gardon dans la Creuse.

## Reproduire le problème

Puisque A est un géant, et que vous n’êtes qu’un microbe, vous n’aurez pas droit à une conférence avec un vrai technicien qui va ouvrir un shell et lancer les commandes. La première étape consiste donc à faire ouvrir le flux pour une machine sous votre contrôle le temps de déboguer (parce que non, on n’ouvre pas un service au web à `any` quand on peut s’en affranchir). 

Soit dit en passant, vous risquez de ne pas pouvoir utiliser les machines internes sans vous faire des nœuds au cerveau en cas de problème de routage. En effet, quand le paquet par à droite et revient par la gauche, les firewalls jouent les Clint Eastwood.

Bref. Vous avez enfin une machine sous la main, le pare-feu Z vous autorise l’accès, ce que vous confirmez ainsi :

	nmap -sT prends-un-verre.moi -p 2345

	Starting Nmap 7.40 ( https://nmap.org ) at 2017-11-21 20:33 CET
	Nmap scan report for prends-un-verre.moi ()
	Host is up (0.000097s latency).
	Other addresses for localhost (not scanned): ::1
	PORT   STATE SERVICE
	2345/tcp open  unknown

	Nmap done: 1 IP address (1 host up) scanned in 0.06 seconds

De là, vous extorquez au développeur de V (ah, pardon ! c’est dans la spécification du service, bien entendu !) une requête de test qui fonctionne. Dans notre cas, on prend un payload JSON.

Pour des questions de praticité, je feinte l’output :

	curl -vv -d '{"liquide":["whisky"]}' https://prends-un-verre.moi/serve
	* Rebuilt URL to: https://www.google.fr/
	*   Trying 2a00:1450:4007:809::2003...
	* TCP_NODELAY set
	* Connected to www.google.fr (2a00:1450:4007:809::2003) port 443 (#0)
	* ALPN, offering http/1.1
	* Cipher selection: ALL:!EXPORT:!EXPORT40:!EXPORT56:!aNULL:!LOW:!RC4:@STRENGTH
	* successfully set certificate verify locations:
	*   CAfile: /etc/ssl/certs/ca-certificates.crt
	  CApath: /etc/ssl/certs
	* TLSv1.2 (OUT), TLS header, Certificate Status (22):
	* TLSv1.2 (OUT), TLS handshake, Client hello (1):
	* TLSv1.2 (IN), TLS handshake, Server hello (2):
	* TLSv1.2 (IN), TLS handshake, Certificate (11):
	* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
	* TLSv1.2 (IN), TLS handshake, Server finished (14):
	* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
	* TLSv1.2 (OUT), TLS change cipher, Client hello (1):
	* TLSv1.2 (OUT), TLS handshake, Finished (20):
	* TLSv1.2 (IN), TLS change cipher, Client hello (1):
	* TLSv1.2 (IN), TLS handshake, Finished (20):
	* SSL connection using TLSv1.2 / ECDHE-ECDSA-AES128-GCM-SHA256
	* ALPN, server accepted to use http/1.1
	* Server certificate:
	*  subject: C=US; ST=California; L=Mountain View; O=Google Inc; CN=*.google.com
	*  start date: Nov  1 13:42:45 2017 GMT
	*  expire date: Jan 24 13:30:00 2018 GMT
	*  subjectAltName: host "www.google.fr" matched cert's "*.google.fr"
	*  issuer: C=US; O=Google Inc; CN=Google Internet Authority G2
	*  SSL certificate verify ok.
	> GET / HTTP/1.1
	> Host: www.google.fr
	> User-Agent: curl/7.56.1
	> Accept: */*
	> 
	< HTTP/1.1 200 OK

Pour vous, ça fonctionne ! Vous avez, en gros :

* réussi à ouvrir la session TCP vers votre service
* annoncé les suites cryptographiques que votre client supporte (`Client Hello`)
* le serveur a retenu l’une des suites (parmi celles qu’il supporte) (`Server Hello`)
* le serveur vous a partagé sont certificat (qui contient sa clé publique et l’autorité de certification pour que vous puissiez vérifier à qui vous parlez) (`Certificate`)
* procédé à la génération des clés de session et à leur échange
* conclut positivement l’échange de clés (les opérations de `change cipher`)

À présent rasséréné par le fonctionnement avec un client moderne, vous pouvez rejouer le cas de G. Compte tenu du protocole, les problèmes possibles sont :

1. l’absence de confiance dans le certificat émis par le serveur
2. l’absence de sous-ensemble concordant entre les suites cryptographiques du client et du serveur
3. une incompatibilité dans l’implémentation du protocole qui fait échouer les contrôles lors de l’échange des clés
4. une interception qui foire au milieu

Pour le premier cas, soit vous avez un certificat émis par une autorité de certification publique reconnue par tous les navigateurs, et donc probablement considérée comme de confiance par votre système d’exploitation, soit vous voyez avec le développeur ou l’admin système pour faire confiance à votre certificat autosigné. Pour un échange entre tiers, cette solution est malavisée, car lors du remplacement du certificat, la synchronisation des équipes ne se fera pas, et vous aurez une indisponibilité le temps d’aligner les deux systèmes.

Le deuxième cas est le plus probable dans le sens où votre répartiteur de charge est configuré selon les préconisations de la sécu, qui dit : TLS 1.2 avec Perfect Forward Secrecy ou va crever !

L’intégrateur va obtenir de l’assistance la liste des suites cryptographiques supportées et vous ne la trouvez pas dans celles autorisées sur le répartiteur. Ainsi, si vous [prenez une suite de TLS 1.0][suites] :

	curl -vv --cipher DHE-RSA-AES128-SHA  https://www.google.fr
	* Rebuilt URL to: https://www.google.fr/
	*   Trying 2a00:1450:4007:817::2003...
	* TCP_NODELAY set
	* Connected to www.google.fr (2a00:1450:4007:817::2003) port 443 (#0)
	* ALPN, offering http/1.1
	* Cipher selection: DHE-RSA-AES128-SHA
	* successfully set certificate verify locations:
	*   CAfile: /etc/ssl/certs/ca-certificates.crt
	  CApath: /etc/ssl/certs
	* TLSv1.2 (OUT), TLS header, Certificate Status (22):
	* TLSv1.2 (OUT), TLS handshake, Client hello (1):
	* TLSv1.2 (IN), TLS header, Unknown (21):
	* TLSv1.2 (IN), TLS alert, Server hello (2):
	* error:14077410:SSL routines:SSL23_GET_SERVER_HELLO:sslv3 alert handshake failure
	* Closing connection 0

Boom ! Cipher non supporté, puisque l’erreur survient lors du message qui est censé indiquer quelle est la suite cryptographique retenue par le serveur (d’où le drapeau d’alerte dans le message TLS).

> — On fait quoi, mon capitaine ?

Puisque vous avez affaire à un client dépassé au point que même Google n’accepte plus la suite cryptographique, vous envoyez la liste des suites supportées à la sécu pour qu’on vous dise s’il est possible de faire une exception.

À bon entendeur, une telle exception n’est donnée qu’en contrepartie d’une feuille de route de mise à jour du système obsolète, laquelle sera suivie par la sécurité pour que vous ne traîniez pas.

Il reste donc à la sécu à se pencher sur l’[excellente recommandation de l’ANSSI sur TLS][anssi] (disponible en français et en anglais !) pour voir si un échange :

* DHE : Diffie-Hellman pour l’échange de clés
* RSA : pour la signature des échanges
* AES128 : AES 128 bits pour la confidentialité des échanges après la poignée de mains
* SHA : SHA-1 pour l’intégrité des échanges après la poignée de main
* *note: n’étant pas un féru du milieu, j’espère ne pas m’être trompé dans la transcription. Pour les détails, c’est [la page 19 du guide v1.1-fr de l’ANSSI][anssi].*

est acceptable dans le contexte de cette application.


La liste officielle des suites est éditée par l’[IANA][iana].


[anssi]: https://www.ssi.gouv.fr/guide/recommandations-de-securite-relatives-a-tls/
[suites]: https://curl.haxx.se/docs/ssl-ciphers.html
[iana]: https://www.iana.org/assignments/tls-parameters/tls-parameters.xhtml
