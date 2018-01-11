---
title: "La CNIL précise l’obligation de sécurité"
authors:
  - jil
date: 2018-01-09
image:  /images/covers/2018-01-10-orage-justice.jpg
categories:
  - Article
tags:
  - CNIL
  - GDPR
  - RGPD
---

# Il vient, il vient, c’est le RGPD…

Il y a l’effet d’annonce pour attiser la curiosité : voilà le premier responsable d’un traitement sanctionné à hauteur de cent mille euros d’amende par la CNIL. Nous sommes en janvier 2018, le RGPD sera appliqué en mai. Cela ressemble à un coup de semonce, mais en y regardant de plus près, c’est très, très intéressant. Oublions les particularités du cas d’espèce pour nous intéresser aux règles générales que l’on peut raisonnablement déduire de la [délibération n°SAN-2018-001 du 8 janvier 2018][san] (le [communiqué][presse] de presse est moins complet).

La Commission dit elle-même qu’il y a lieu d’en tirer des enseignements :

> La formation restreinte considère qu’au regard des éléments précités, du
> contexte actuel dans lequel se multiplient les incidents de sécurité et
> de la nécessité de sensibiliser les internautes quant au risque
> pesant sur la sécurité de leurs données, il y a lieu de rendre publique
> sa décision.

# C’est pour ta pomme

Peu importe que le sous-traitant ait rendu accessible un formulaire qui n’apparaît pas dans le cahier des charges. Peu lui chaut que le responsable du traitement ignore son existence. Point de transfert de responsabilité tant que le sous-traitant ne traite pas les données ainsi collectées pour ses propres fins (ce point changera avec le RGPD pour une coresponsabilité). Ce sont tes données, ton traitement ; c’est insécurisé, tu trinques.

# Un audit avant la mise en production, tu feras…

> La formation restreinte considère qu’en retenant un logiciel standard
> dit sur étagère proposé par son prestataire, il incombait à la société
> de procéder aux vérifications des caractéristiques de ce produit qui
> auraient permis d’identifier le risque résultant de l’existence d’un
> accès aux données des clients contenues dans l’outil de gestion des
> demandes de service après-vente et d’empêcher celui-ci.

> La vérification préalable notamment des règles de filtrage des URL
> fait partie des tests élémentaires qui doivent être réalisés par une
> société en matière de sécurité des systèmes informatiques.

En l’espèce, il semble qu’il suffisait de décrémenter l’ID de la fiche pour trouver celle d’un tiers. Conclusion logique, faites un audit avec parcours automatisé du site pour vous assurer de ce qui est disponible, et faites-le auditer avant la mise en production (c’est toujours plus facile de changer des paramètres avant le lancement de l’application).

# …et régulièrement, tu recommenceras.

Les audits doivent ensuite être réguliers, nous dit la CNIL :

> il appartenait à la société de procéder de façon régulière à la revue
> des formulaires de demande de service après-vente accessible et
> permettant d’alimenter l’outil de gestion des demandes de service 
> après-vente.

Il faut donc contrôler les points d’entrée des données, qui sont les points les plus susceptibles d’être attaqués. 

# Point ne suffit de promptement réagir

Même si la vulnérabilité a été corrigée dans un délai que la commission restreinte considère comme raisonnable, il est reproché au gestionnaire de traitement de ne pas avoir demandé *quotidiennement* l’avancée de la correction à son prestataire. Gardez les traces, les amis.

Et collaborer avec la CNIL ne suffit pas à limiter le montant de l’amende. Un mythe s’effondre-t-il ainsi ? Ah ! Que vous vois-je frissonner !

# Qui c’est qui va demander une rallonge de budget ? C’est bibi !

Déjà qu’on a parfois du mal à recenser tous les sites qui collectent des données, il faut à présent les auditer de manière systématique et régulière. C’est l’occasion de préciser que si vous pouvez être tentés de vous appuyer sur des outils automatisés, il convient de comparer leur résultat à celui d’un pentester de chair et d’os, car je pense que la défense qui consisterait à s’abriter derrière les rapports des solutions automatisées s’avèrera insuffisante.

Mes pauvres pentesters, déjà que la plupart des sites vous font pleurer tant ils sont friables, en voilà une pelletée de plus qui s’en vient !

Mes pauvres RSSI, voyez-vous l’orage qui s’annonce ? Bonne année !

À noter que cette vulnérabilité a été signalée par un tiers à la CNIL en précisant l’ampleur de la divulgation. En voilà un qui devrait se faire discret.

[san]:https://www.legifrance.gouv.fr/affichCnil.do?oldAction=rechExpCnil&id=CNILTEXT000036403140&fastReqId=306045536&fastPos=1
[presse]: https://www.cnil.fr/fr/darty-sanction-pecuniaire-pour-une-atteinte-la-securite-des-donnees-clients


