---

title: "Achats en ligne : Risques associées et comment s’en prémunir"
date: 2014-01-25T09:30:14+00:00
publisher: morgan


aliases: /2014/01/achats-en-ligne-risques-associees-et-comment-sen-premunir/
views: 1090
image: /images/covers/2014-01-Credit-Card-Fraud.jpg
categories:
  - Article
tags:
  - ecarte bleue
  - nagraID
  - otp
  - Target
---
<ins datetime="2014-01-23T21:35:25+00:00">Cet article est un import depuis l'ancienne plateforme de publication Medium, depuis, on sait que la [les caisses enregistreuses](http://krebsonsecurity.com/2014/01/a-first-look-at-the-target-intrusion-malware/" >fuite </a>de [PoS ](http://krebsonsecurity.com/2014/01/a-closer-look-at-the-target-malware-part-ii/" >Target </a>provient des <a href="http://techcrunch.com/2014/01/16/meet-the-malware-that-took-down-target/)(Point of Sale), autrement dit le réseau où se situent <a href="http://www.infosecisland.com/blogview/23561-How-Targets-Point-of-Sale-System-May-Have-Been-Hacked.html).</ins>

<ins datetime="2014-01-23T22:09:19+00:00">Les mesures de sécurité qui étaient données dans cet article sont efficaces pour des achats en ligne. Dans le cas de l'affaire Target. Il n'y a rien à faire pour l'utilisateur final, mis à part payer systématiquement en cash...ce qui vous expose à d'autres risques comme le racket, la perte ou le retrait sur un distributeur piraté!<ins>



    <ins datetime="2014-01-23T22:09:19+00:00">Cela me fait penser que la plupart des DAB tournent sous Windows XP, [un système qui ne sera bientôt plus supporté par Microsoft](http://mobile.businessweek.com/articles/2014-01-16/atms-face-deadline-to-upgrade-from-windows-xp)...</ins>





    L’achat en ligne se démocratise, [les résultats record d’Amazon](http://www.boursier.com/actualites/economie/noel-record-pour-amazon-22495.html) pour Noël 2014 le confirment. Bien naturellement, le nombre de fuites de cartes bancaires auprès de ces différents e-commerces augmente également.





    ![TargetLogo](/images/2014/01/TargetLogo.jpg)





    Le dernier en date est [Sony](http://krebsonsecurity.com/2013/12/cards-stolen-in-target-breach-flood-underground-markets/" >Target</a>, 5e distributeur en volume aux États-Unis. En octobre, nous avions [Adobe](http://krebsonsecurity.com/2013/10/adobe-breach-impacted-at-least-38-million-users/), et quelques années auparavant nous avions le géant <a href="http://krebsonsecurity.com/2011/04/millions-of-passwords-credit-card-numbers-at-risk-in-breach-of-sony-playstation-network/). Ceux-là ne sont que les plus gros, bien d’autres se font attaquer régulièrement, et il y a de grandes chances pour que la majorité des attaques ne soit jamais connue du public. Elles ne sont parfois même jamais connues du magasin…





    La plupart des vols de donnée concernent uniquement le numéro de carte de crédit, parfois accompagné de leur CVV. Les dernières informations concernant Target sont préoccupantes, car elles contiendraient également [les codes PIN](http://money.cnn.com/2013/12/27/technology/target-pin/).





    Target affirme que les PIN étaient stockés de façon chiffrée, que le chiffrement était réalisé par leur prestataire de paiement et que seul celui-ci possède la clé. En d’autres termes, les voleurs ont un gros coffre sans la clé. Restons prudents, ce n’est pas la première fois qu’un vendeur nous rassure sur la robustesse de son dispositif de sécurité, pour voir annoncées quelques semaines plus tard que celui-ci a été cassé ou contourné.




  <h2>
    Qu'est ce qu'un prestataire de paiement
  </h2>




    ![4-etapes-pour-une-transaction-3d-secure](/images/2014/01/4-etapes-pour-une-transaction-3d-secure.gif)





    Il est très rare que les commerces en ligne réalisent les paiements eux-mêmes. Lorsque vous validez votre panier, vous êtes transféré sur le site internet d’une autre entreprise, le prestataire de paiement, qui prendra vos informations bancaires, vérifiera la légitimité de la transaction, et avertira le commerçant que la transaction s’est passée dans les règles et qu’il peut valider la commande.





    Cela permet au commerçant de totalement se décharger de cette action très encadrée par l’industrie des cartes de paiement, qui impose de montrer patte blanche avec une conformité à leur standard PCI DSS.





    ![pci-dss](/images/2014/01/pci.gif)




  <h2>
    Target n’a jamais vu la clé ou ne l’a juste jamais stocké ?
  </h2>




    C’est toujours le même problème lorsque l’on souhaite chiffrer une donnée qui vit, j’entends par là qui est régulièrement lue et modifiée, il faut que la clé de déchiffrement ne soit pas bien loin. Les bonnes pratiques concernant les bases de données consistent donc généralement à chiffrer la base de données et à stocker la clé de chiffrement sur un autre serveur.





    <a href="/images/2014/01/transdata.gif">![transdata](/images/2014/01/transdata.gif)





    Lorsque la base de données démarre, elle récupère la clé de chiffrement et la met en mémoire vive. La plupart des attaques sur les bases de données sont des fuites de type injection SQL. Dans ce genre d’attaques, ce n’est pas tout le serveur qui est compromis, juste l’accès en lecture ou écriture sur la base de données. Par conséquent, les attaquants peuvent potentiellement siphonner l’intégralité des données en base, mais n’ont pas accès à la mémoire vive du serveur ou au HSM et donc pas accès à la clé de déchiffrement.





    Cette méthode à ses limites. S’il existe une fonction dans l’application, ou une procédure stockée dans la base, qui permet d’accéder à cette clé pour chiffrer et déchiffrer la donnée, elle est peut-être également accessible par l’attaquant. Si l’intégralité du serveur est compromise, la recherche en mémoire vive d’une clé de chiffrement n’est pas triviale, mais est tout à fait possible.





    Enfin, quand bien même la clé n’aurait jamais été stockée ni même accédée par les serveurs de Target, il est fort probable que Target dispose d’un accès privilégié au système d’information de son prestataire de paiement. Une attaque par rebond depuis les serveurs de Target vers celui de leur prestataire reste envisageable.
  




    ![attaque-rebond](/images/2014/01/rebond.gif)





    J’avoue ne pas comprendre pourquoi le code PIN aurait été stocké dans un premier lieu. Pour un achat en ligne, seuls le numéro de carte, le nom du porteur, la date de validité et le CVV (code au dos de la carte) sont nécessaires pour valider la transaction. Cela voudrait dire que les codes PIN saisis sur les terminaux de paiement dans les magasins sont enregistrés et stockés ? À ma connaissance, le code PIN n’est pas transmis au système de paiement, il sert uniquement à « déverrouiller » la carte. Il est transmis à la carte et celle-ci informes le terminal de paiement que le code fourni est correct et qu’il peut continuer la transaction.





    C’est d’ailleurs ce qu’exploitent les [Yes Cards](http://fr.wikipedia.org/wiki/YesCard), cartes à puce modifiées, répondant systématiquement « Oui, le code est bon » peu importe le PIN fourni par l’utilisateur.





    Quoi qu’il en soit, le mal est fait pour les clients de Target. Ce dernier leur a d’ailleurs conseillé d’aller faire changer leur code PIN, et leur carte bancaire dans la foulée « par mesure de précaution ».




  <h2>
    Comment se protéger ?
  </h2>



  <h3>
    Ne pas enregistrer ses données bancaires
  </h3>




    Plus facile à dire qu’à faire, le problème étant que, même si le site vous dit qu’il ne stocke pas votre numéro, vous n’avez aucune garantie sur la véracité de cette affirmation ! Il est vrai qu’enregistrer sa carte peut être tentant, payer sur des sites comme Amazon avec une carte enregistrée est un confort indéniable, cependant aucun commerce en ligne, même un géant comme Amazon, n’est à l’abri d’une attaque.





    ![dashlane-2-0-credit-cards](/images/2014/01/326280-dashlane-2-0-credit-cards.jpg)





    Si saisir votre numéro de carte vous donne des boutons, confier le plus tôt à un coffre-fort numérique de type [Keepass](https://www.dashlane.com/" >Dashlane</a>, [1Password ](https://lastpass.com/" >Lastpass</a>, <a href="https://agilebits.com/onepassword)ou <a href="http://keepass.info/). Dans tous les cas, évitez le stockage dans le navigateur, ou mettez à minima un mot de passe maître sur celui-ci, sinon ils ne bénéficieront d’aucune protection. Ils seront stockés en clair ou bien chiffrés avec une clé stockée juste à côté, ce qui est aussi sécurisé que de mettre la clé de chez soi sous le paillasson ou dans le pot de fleurs.




  <h3>
    3D Secure
  </h3>




    <a href="/images/2014/01/schema-3D-secure.jpg">![schema-3D-secure](/images/2014/01/schema-3D-secure.jpg)





    Le système 3D-Secure n’est pas une protection pour le client, c’est une protection pour le commerçant ! En effet, ce système est optionnel et activé à la demande du commerçant. Il lui permet de valider, pour tous les achats ou ceux dépassant un certain montant, que le client ait bien le possesseur de la carte. Pour cela, un SMS est envoyé au numéro de téléphone associé au porteur de la carte, celui-ci doit le renseigner sur le site internet pour valider sa commande. Ce système ne vous protège en rien, en effet, de très nombreux commerces en ligne ne vérifient pas ce système, l’attaquant n’aura qu’à utiliser votre carte de crédit dans l’une de ces boutiques.




  <h3>
    e-Carte bleue
  </h3>




    ![ecarte-bleu](/images/2014/01/au_quotidien.Par_.12044.Image_.256-1.png)





    Le principe est simple, vous pouvez créer autant de « cartes virtuelles » que vous le souhaiter. Chaque carte à un numéro unique, une capacité d’achat et une durée de validité choisie par vos soins. Ce service est proposé par la quasi-totalité des banques émettrice de cartes bleues [Visa ](http://www.visa.fr/les-innovations-visa/paiement-sur-internet/e-carte-bleue.aspx)et MasterCard. Selon les banques le service va de gratuit à une dizaine d’euros par an, et il y a souvent une application smartphone mis à disposition pour la création rapide de nouveaux numéros.





    Le système à ses défauts, il faut faire l’effort de créer une nouvelle carte pour chaque transaction, avec comme capacité le prix d’achat, si l’on souhaite une sécurité maximale. Cela demande un effort non négligeable de la part du client. Enfin, certains services se basant sur la détention du numéro de carte pour s’assurer du savoir-vivre de leur client, j’entends par là tous les services d’hôtellerie et de location, vous refuseront en général ce type de dispositif.




  <h3>
    Cartes de crédit avec OTP
  </h3>




    ![displaycard_sc](/images/2014/01/displaycard_sc.jpg)





    Variante au 3D-Secure, le principe est ici d’accompagner la carte bancaire d’un token, ou, plus confortable, de l’intégrer directement dans la carte de crédit comme le fait [NagraID](http://www.nidsecurity.com/microsite/mastercard/products/index.html). Étrangement, ce système existe depuis au moins 2008 et pourtant je n’en ai jamais vu en circulation. De toute façon, ce système ne vaut pas mieux que 3DSecure si la demande de l’OTP est optionnelle pour valider la transaction en ligne. J’ai beau chercher, je n’arrive pas à savoir si c’est le cas.




  <h2>
    Restons calme
  </h2>




    Quoi qu’il en soit, l’objectif de cet article n’est pas de vous faire arrêter l’utilisation du commerce en ligne. Il est important de montrer que le système n’est pas infaillible, mais il ne faut pas céder à la panique pour autant. Nous sommes jusqu’à présent relativement bien protégés en France, la grande majorité des pertes par fraude à la carte bancaire sont couvertes non par le client ou par la banque, mais par les commerçants, et [la note peut vite être très salée](http://www.hubinternational.com/data-breach-cost-calculator/).
