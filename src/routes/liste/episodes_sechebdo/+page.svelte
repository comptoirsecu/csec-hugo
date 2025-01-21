<script lang="ts">
    import all_podcasts from '$lib/all_sechebdo.json';

    const all_urls = all_podcasts.map(podcast => podcast.metadata.podcast.feed).join("\n");
</script>
<svelte:head>
    <title>Liens de téléchargement des archives des épisodes thématiques du Comptoir Sécu</title>
    <meta name="description" content="Le Comptoir Sécu était un podcast à destination des professionnels de l'informatique, traitant de la cybersécurité à travers des dossiers thématiques et des revues de l'actualité spécialisée. Cette page contient les liens vers ses épisodes.">
</svelte:head>
<pre>
<code class="text-white">
{all_urls}
</code>
</pre>