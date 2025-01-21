<script lang="ts">
    import Podcast from '$lib/Podcast.svelte';
    import all_podcasts from '$lib/all_sechebdo.json';
</script>
<svelte:head>
    <title>Archives du Comptoir Sécu</title>
    <meta name="description" content="Le Comptoir Sécu était un podcast à destination des professionnels de l'informatique, traitant de la cybersécurité à travers des dossiers thématiques et des revues de l'actualité spécialisée. Cette page contient les liens vers ses épisodes.">
</svelte:head>

<h1 class="text-titres text-white text-3xl pl-4 pt-4">Archives des revues d'actualité</h1>
<p class="text-titres text-slate-200 italic pl-4 pb-4 m-0">Ctrl+F est ton ami 😆. Si tu es gourmand, voici la <a class="underline" href="liste/episodes_sechebdo">liste des liens</a> pour wget them all!</p>

<div class="lg:grid lg:grid-cols-3 lg:items-stretch justify-center">
    {#each all_podcasts as podcast}
        <Podcast data={podcast} />
    {/each}
</div>
