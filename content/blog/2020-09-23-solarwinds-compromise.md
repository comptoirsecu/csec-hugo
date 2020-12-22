---
title: "Compromission par chaîne d’approvisionnement - SolarWinds"
authors:
  - swithak
  - jil
date: 2020-12-23
image:  /images/covers/2020-07-17-sans-dfir-2020.jpg
categories:
  - Article
tags:
  - forensic
  - sunburst
  - teardrop
  - supernova
  - soloriate
  - solarstorm
  - darkallow
  - unc2452
---

# Contexte 

## Mise en bouche

Le père fouettard est passé tôt cette année. Le 8 décembre 2020, FireEye, après avoir [découvert][fireeye-faq] des enregistrements secondaires sur leur authentification à deux facteurs ainsi que des écarts comportementaux, annonce avoir été partiellement compromis par un attaquant avancé, très probablement à la solde d’un État. 

Le 13 décembre, le vecteur d’intrusion initial est rendu public : SolarWinds Orion, un composant du socle applicatif de SolarWinds présent dans une dizaine de produits. L’attaquant a introduit une porte dérobée, nommée [Sunburst][fireeye-sunburst] par FireEye, dans le code de l'application sous la forme d’une DLL malveillante. Une première altération chez SolarWinds aurait eu lieu entre [septembre et octobre 2019][SOURCEMANQUANTE]. La charge malveillante aurait été introduite en mars, pour une campagne s’étalant [de mars à juin][ACONFIRMER].

SolarWinds est une société États-Unienne qui développe des logiciels permettant la gestion centralisée des réseaux, des systèmes et de l'infrastructure informatique. Ses produits disposent d’accès aux infrastructures de la victime et communique des métriques d’utilisation à des fins d’amélioration du produit.

## Et alors ? On en voit tous les jours, des logiciels pourris upstream, non ?

Cette campagne est une orfèvrerie. Le code intégré à Orion respecte totalement les pratiques de développement du logiciel. Le fonctionnement de la porte dérobée utilise des fichiers de configuration. Elle communique avec son C2 en se fondant dans le protocole Orion Improvement Program. Elle tente subrepticement de désactiver certains produits de sécurité sans laisser de trace. Teardrop, le deuxième étage de l’attaque, serait un dropper Cobalt-Strike résidant uniquement en mémoire. Il démarre un service, lit sa configuration depuis un fichier JPEG et vérifie la présence d’une clé de registre. Une fois Cobalt-Strike opérationnel, il communique vers un autre C2, et nous sommes presque de retour en terrain connu. Presque, car cet attaquant aurait extrait le certificat de signature des jetons ADFS pour pouvoir forger ses propres jetons valides et se connecter aux ressources dont l’accès est contrôlé par ce service de fédération d’identités.

## Cibles et portée de la campagne

En se basant sur la télémétrie à sa disposition, Microsoft annonce le 17 décembre qu’environ vingt pourcent des compromissions détectées le sont en dehors des États-Unis : au Canada, au Mexique, en Belgique, en Espagne, au Royaume-Uni, en Israël ainsi qu’aux Émirats Arabes-Unis.

![ms-map-picture](https://1gew6o3qn6vx9kp3s42ge0y1-wpengine.netdna-ssl.com/wp-content/uploads/prod/sites/5/2020/12/cyber1.jpg)

## Répercussions diplomatiques et stratégiques

Il n’y a pas eu, à cette heure, de mise en accusation. Il ne fait aucun doute que le dossier est en train d’être monté et que les services de renseignement vont passer des fêtes épuisantes.

Alors, s’agit-il d’une prémisse de ce qui nous guette avec la 5G bâtie sur des équippements de l’étranger ?

## Attribution

Ni FireEye ni Microsoft n’accusent la Russie en tant qu’État. 

# Contrôles minimaux

## Savoir si je suis touché

Votre antivirus a dû hurler. Oui, il faut regarder la console AV, c’est la première étape.

Présence de la DLL `SolarWinds.Orion.Core.BusinessLayer.dll` et du fichier `SolarWinds.Orion.Core.BusinessLayer.dll.config` associé.

La porte dérobée ne s’exécute pas si un produit de sécurité est détecté et qu’elle ne peut le désactiver. Pour le désactiver, elle va prendre le contrôle de la clé de registre qui définit le service en question, modifier son mode de démarrage à Désactivé, puis restaurer l’ACL d’origine sur la clé. Elle attendra ensuite que le service s’arrête de lui même, probablement au prochain reboot.

La liste des produits détectés, qui sont répartis entre les produits qui font cesser toute opération à Sunburst, et ceux dits malléables, est disponible [ici][URLÀRETROUVER].

Si vous avez les deux clés suivantes dans le fichier `.config` : 

	METTRE CODE 

Si la valeur de la première clé est 3 et la deuxième est LISTE, Teardrop (le deuxième étage) n’a pas été déployé.

Si la valeur de la première clé est 3 et la deuxième est LISTE, Teardrop (le deuxième étage) n’a pas été déployé, mais vous devriez confirmer que votre antivirus n’a pas été désactivé en vérifant les dates de modifications de la clé de registre du produit concerné.

## Mettre à jour vers quelle version ?

Versions minimales d’Orion saines, selon la branche :

- 2020.2.1 hotfix 2 
- 2019.4 hotfix 6

# Plongeon dans la beauté de la bête



[fireeye-faq]:https://www.fireeye.com/current-threats/sunburst-malware.html
[fireeye-sunburst]:https://www.fireeye.com/blog/threat-research/2020/12/evasive-attacker-leverages-solarwinds-supply-chain-compromises-with-sunburst-backdoor.html
[ms-map]:https://blogs.microsoft.com/on-the-issues/2020/12/17/cyberattacks-cybersecurity-solarwinds-fireeye/
