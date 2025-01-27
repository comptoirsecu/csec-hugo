<script lang="ts">
    import { Button, AccordionItem, Accordion } from 'flowbite-svelte';
    import { ClapperboardPlaySolid, DownloadSolid } from 'flowbite-svelte-icons';
    import SvelteMarkdown from 'svelte-markdown'

    let readerHidden = $state(true);
    let props = $props();
    const title = props.data.metadata.title;
    const subtitle = props.data.metadata.subtitle ?? '';
    const url = props.data.metadata.podcast.feed;
    const description = props.data.content.description;
    const links = props.data.content.links;
    const songs = props.data.metadata.songs ?? null;
    const tags = props.data.metadata.tags.map(tag => `#${tag}`).join(', ');

    function formatDate (dateString: string) {
        const options: Intl.DateTimeFormatOptions = { year: 'numeric', month: 'long', day: 'numeric' };
        return new Date(dateString).toLocaleDateString(undefined, options);
    };
    const date = formatDate(props.data.metadata.date);

    let guests: string[] = [];
    props.data.metadata.authors?.forEach(author => {
        guests.push(author.name);
    });
    props.data.metadata.guests?.forEach(guest => {
        guests.push(guest.name);
    });
</script>   

<div class="p-2">
    <div class="rounded-lg border basis-1/3">
        <div class="bg-black rounded-lg p-4">
            <h2 class="font-titres text-slate-100 text-xl">{title}</h2>
            <h3 class="text-slate-400 text-lg">{subtitle}</h3>
            <div class="text-slate-500 italic align-right text-sm">{date}</div>
            <div class="text-slate-500 align-right text-sm">{tags}</div>
        </div>
        <div class="bg-gray-300 p-4">
            <div class="flex mb-2">
                <Button color="light" size="xs" class="mr-2" on:click={() => readerHidden=!readerHidden}><ClapperboardPlaySolid class="w-4 h-4 me-2" /> Écouter</Button>
                <Button href={url} color="light" size="xs"><DownloadSolid class="w-4 h-4 mr-2" /> Télécharger</Button>
                {#if !readerHidden}
                <audio controls class="p-2">
                    <source src="{url}" type="{url.endsWith('.mp3') ? 'audio/mpeg' : 'audio/aac'}">
                    Ton navigateur ne sait pas lire les audios, mais tu peux le <a href="{url}">télécharger ici</a>.
                </audio>
                {/if}
            </div>
            <Accordion defaultClass="text-black">
            <AccordionItem open paddingDefault="p-2" paddingFlush="py-2">
                <span slot="header">Contenu</span>
                <SvelteMarkdown source={description} />
            </AccordionItem>
            
            <AccordionItem paddingDefault="p-2" paddingFlush="py-2">
                <span slot="header">Sources</span>
                <ul class="list-disc list-inside">
                    {#each links as link}
                    <li><a href="{link.url}">{link.title}</a></li>
                    {/each}
                </ul>
            </AccordionItem>
            <AccordionItem paddingDefault="p-2" paddingFlush="py-2">
                <span slot="header">Participants</span>
                <p>Avec la participation de 
                    {guests.join(', ')}
                </p>
            </AccordionItem>
            {#if songs}
            <AccordionItem paddingDefault="p-2" paddingFlush="py-2">
                <span slot="header">Musiques</span>
                <ul class="list-disc list-inside">
                    {#each songs as song}
                    <li><a href="{song.url}">{song.song} ({song.artist})</a></li>
                    {/each}
                </ul>
            </AccordionItem>
            {/if}
            </Accordion>
        </div>
    </div>
</div>