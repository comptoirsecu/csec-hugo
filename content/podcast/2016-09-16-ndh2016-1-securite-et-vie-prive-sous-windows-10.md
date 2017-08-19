---

title: "Nuit du Hack 2016 - 1 : Sécurité et vie privé sous Windows 10"
date: 2016-09-16T18:55:43+00:00


aliases: /2016/09/ndh2016-1-securite-et-vie-prive-sous-windows-10/
podcast:
  feed: https://podcast.comptoirsecu.fr/CSEC.HS23.2016-07-02.NDH2k16_Windows10.mp3
views: 923
image: /images/covers/2016-09-ndh2k16.jpg
categories:
  - Podcast
  - Hors-Serie
tags:
  - live
  - ndh2k16
  - vie privée
  - windows 10

authors:
  - lois
  - justin
  - morgan
  - youenn

guests:
  - thomas_aubin
  - paul_hernault

---
Première épisode enregistré en live à la Nuit du Hack 2016, encore tout chaud! Bon, ok, légèrement tiède...

Dans cette épisode, nous nous entretenons avec Thomas Aubin et Paul Hernault sur la sécurité et le respect de la vie privé dans Windows 10, dans la continuité directe de leur conférence.

{{< podigee >}}

Pour ceux qui n'avaient pas la chance d'être là, ci-dessous l'abstract et la rediffusion de leur conférence :

> Nous avons commencé nos recherches en nous intéressant aux droits légaux de Microsoft à propos de nos données. Lire les conditions d'utilisation est souvent vu comme une plaie, les gens l'acceptent sans vraiment y faire attention. Cependant, son contenu peut être intéressant et cacher des informations importantes. Dans cette partie du talk, nous parlerons du type de données qui sont collectées et les dangers de leurs divulgations. Quels sont les pouvoirs de Microsoft sur nos données ?
>
> Si il y a un problème de vie privée dans windows, il y a communication, c'est pourquoi nous avons analysé les flux sur le réseau. Nous parlerons principalement de notre installation et de ce que nous avons découvert en examinant ces flux.
>
> Un rapide tour sur comment fonctionne un man-in-the-middle sur SSL pour ceux qui ne le savent pas. Nous expliquerons également comment nous avons mis en place notre environnement pour étudier la communication de Windows. Nous avons décrypté plusieurs paquets SSL et trouvé que les données n'étaient pas anonymisées du tout. Chaque utilisateur a un identifiant utilisé avec plusieurs outils de Windows 10 (Cortana, recherches en lignes etc.) Nous n'avons pas réussi à décrypter chaque paquet, nous ne savons pas vraiment pourquoi, c'est comme si ils avaient détecté notre MITM, l'ayant contourné ou tout simplement refusant d'accepter un autre certificat.
>
> Après la sortie de Windows 10, plusieurs développeurs voulaient préserver leur vie privée et ont décidé de créer un programme afin de bloquer automatiquement toutes les IP et DNS de Microsoft. La base de données de ces programmes doit être régulièrement mise à jour. Plusieurs programmes disposent de plus d'options comme désinstaller les applications Metro, mises à jour, cherchant plusieurs solutions, nous avons choisi de passer notre temps à analyser le plus utilisé (DWSLite). Etant open-source, nous avons sondé son code trouvant d'étranges modifications appliquées au système. Plutôt que de bloquer tous les services pourquoi ne pas les fausser et berner Microsoft à propos de nos profils? Les utilisateurs peuvent alors continuer à se servir des services sans être enregistrés. L'idée est d'envoyer beaucoup de données afin de rendre indissociable les vraies des fausses. Nous avons mis en place un proof of concept sur

> un service : Cortana (Nous pourrions le faire sur n'importe quel autre service comme celui des diagnostics). Le programme que nous avons développé s'appelle CortaSpoof. Ce programme envoie en continue des requêtes aléatoires au serveur Bing.

{{< video "https://www.youtube.com/embed/dkBn6kKRs6g" >}}
