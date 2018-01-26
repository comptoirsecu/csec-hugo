---
title: "RGPD : le chiffrement ne vous absoudra pas"
authors:
  - jil
date: 2018-01-26
image:  /images/covers/2018-01-26-banniere-rgpd-cnil.jpg
categories:
  - Article
tags:
  - GDPR
  - RGPD
  - chiffrement
---

# Chiffrer, c’est la clé !

Comme la plupart des professionnels de la sécurité de l’information, j’en ai ras la casquette des RGPD ceci, RGPD cela… Il y en a tout de même un qui revient fréquemment et qui mérite de s’y arrêter : « chiffrez, vous serez sauf ! »

Que ce soit un intégrateur peu au fait de la sécurité, ou un commercial qui s’exerce à la fricassée de *buzzwords*, on vous laisse entendre que chiffrer les bases de données vous libérera d’une grande partie des contraintes du règlement. D’après eux, si vous achetez l’option (onéreuse en argent, et en performance sans le matériel adéquat), vous pourrez dormir sur vos deux oreilles.

Il s’agit de chiffrer les fichiers qui contiennent les données de la base. Cela vous prémunit d’un arrachement de disque dans le cas où le secret du chiffrement n’est pas sur ledit disque, ou d’une copie des fichiers de la base de données effectuée par un utilisateur ayant accès à la machine et qui n’aurait pas réussi à élever ses privilèges pour récupérer le secret en mémoire (ou dans la configuration). Je vous invite à prendre la liste de vos menaces et regarder où celles-ci se trouvent. Il y a fort à parier que ce soit en bas, loin en bas. Surtout si vous durcissez vos serveurs et qu’un utilisateur ne franchit pas l’ACL du dossier où sont stockées les données.

# RSSI : le rabat-joie

En tête à tête, cet argument est facile à balayer et vous pouvez raccrocher. En présence des métiers, il me semble préférable de démonter l’approche sans dire qu’une mesure de sécurité est inutile. C’est toujours utile, une mesure de sécurité. Si si. Pour vos utilisateurs, ça doit toujours être utile. 

Il s’agit donc d’expliquer que la menace première réside dans la compromission de la couche applicative (ou « métier », appelez-la comme bon vous semble), par une injection SQL ou par une exécution de code arbitraire (oui, bon, c’est la même chose, mais si vous avez quelqu’un de la DSI, il devrait pouvoir se raccrocher à « injection SQL » et se croire compétent). Or, quand le filou qui en a après vos données prend le contrôle de la couche applicative, il détient des accès valides et, pour les IDS, légitimes, aux données, lesquelles sont servies par le moteur de la base de données en clair. Dans ce scénario d’attaque, le chiffrement des données sur le disque n’est pas suffisant.

Normalement, votre interlocuteur vous regarde avec les yeux de merlan frit, mais c’est surtout vos partenaires qu’il faut observer pour s’assurer qu’ils ont compris qu’il y a autre chose à faire que payer une licence et activer une option. Vous pouvez alors glisser une ou deux mesures appropriées au contexte. Et n’oubliez pas d’ajouter qu’un audit de sécurité préalable à la mise en production est désormais un impératif de [la réglementation][cnil].

Si d’aventure l’on vous parle d’une option de chiffrement payante pour une solution « cloud », demandez-leur, si c’est une obligation réglementaire, pourquoi leur solution n’est pas en conformité dans le prix de base. Effet garanti :D

# Pourquoi diantre cette ruée vers le chiffrement ?

Je me suis longtemps demandé d’où sortait cette idée saugrenue que le chiffrement des données au repos était salutaire. 

Le [considérant 83][lex] (le bla-bla à peu près compréhensible qui précède les articles) nous indique : 

> Afin de garantir la sécurité et de prévenir
> tout traitement effectué en violation du présent règlement, il importe
> que le responsable du traitement ou le sous-traitant évalue les risques
> inhérents au traitement et mette en œuvre **des mesures pour les atténuer**,
> *telles que le chiffrement*. Ces mesures devraient assurer un niveau de
> sécurité approprié, y compris la confidentialité, compte tenu de l'état
> des connaissances et des coûts de mise en œuvre par rapport aux risques
> et à la nature des données à caractère personnel à protéger. Dans le
> cadre de l'évaluation des risques pour la sécurité des données, **il
> convient de prendre en compte les risques que présente** le traitement de
> données à caractère personnel, tels que la destruction, la perte ou
> l'altération, **la divulgation non autorisée de données à caractère
> personnel transmises, conservées ou traitées d'une autre manière ou
> l'accès non autorisé à de telles données, de manière accidentelle ou
> illicite, qui sont susceptibles d'entraîner des dommages physiques,
> matériels ou un préjudice moral.**

Le Législateur nous dit qu’il convient de limiter les risques de divulgation non autorisée, et pour ce faire, le chiffrement peut être une bonne idée. Vous noterez qu’il ne précise pas à quel niveau intervient le chiffrement, et que seul lui importe la réduction du risque en phase à la fois avec l’état de l’art, votre portefeuille et la sensibilité des données.

On va donc retrouver cette notion dans l’article 32 relatif à la sécurité du traitement :

> (1) Compte tenu de l'état des connaissances, des coûts de mise en œuvre et de la nature, de la portée, du contexte et des finalités du traitement ainsi que des risques, dont le degré de probabilité et de gravité varie, pour les droits et libertés des personnes physiques, le responsable du traitement et le sous-traitant **mettent en œuvre les mesures techniques et organisationnelles appropriées afin de garantir un niveau de sécurité adapté au risque, y compris entre autres, selon les besoins**:
>
> a) la pseudonymisation et *le chiffrement* des données à caractère personnel;
>
> b) des moyens permettant de garantir la confidentialité, l'intégrité, la disponibilité et la résilience constantes des systèmes et des services de traitement;
>
> c) des moyens permettant de rétablir la disponibilité des données à caractère personnel et l'accès à celles-ci dans des délais appropriés en cas d'incident physique ou technique;
>
> d) une procédure visant à tester, à analyser et à évaluer régulièrement l'efficacité des mesures techniques et organisationnelles pour assurer la sécurité du traitement.

Le Législateur aurait difficilement pu prendre plus de gants à l’égard du chiffrement. 

Toutefois, je pense que c’est surtout l’article 34, relatif à la notification des personnes, qui fait pousser les ailes de nos interlocuteurs :

> (3) La communication à la personne concernée visée au paragraphe 1 n'est pas nécessaire si l'une ou l'autre des conditions suivantes est remplie:
> 
> a) le responsable du traitement a mis en œuvre les mesures de protection techniques et organisationnelles appropriées et ces mesures ont été appliquées aux données à caractère personnel affectées par ladite violation, en particulier les mesures qui rendent les données à caractère personnel incompréhensibles pour toute personne qui n'est pas autorisée à y avoir accès, *telles que le chiffrement*;

Là ! On le tient, le raccourci : « tu as chiffré, t’es sauvé, tu n’auras pas à communiquer ! » Que nenni ! Si les données volées sont inexploitables, et uniquement dans ce cas-là, tu es dispensé. Or, avec le chiffrement des fichiers sur le disque et une attaque par injection SQL, le chiffrement est ineffectif et communiquer, il faudra.

# TL;DR

Le RGPD n’impose aucunement le chiffrement ; il impose toute mesure proportionnée au risque. Le chiffrement des données sur disque est probablement inefficace contre l’une des menaces les plus importantes : celle de la compromission de la couche applicative pour exécuter un code arbitraire (pensez à la fréquence des vulnérabilités liées aux injections SQL). Enfin, si le filou s’en va avec des données inexploitables, vous n’aurez pas à communiquer l’incident. En revanche, si votre base est chiffrée, mais que l’attaquant a pu obtenir les données en clair, ou, dans la compromission, ait été susceptible d’atteindre les secrets qui participent au chiffrement, vous devrez communiquer.


[lex]: http://eur-lex.europa.eu/legal-content/FR/TXT/?uri=CELEX:32016R0679
[cnil]: https://www.comptoirsecu.fr/blog/2018-01-10-cnil-obligation-securite-precisee/

