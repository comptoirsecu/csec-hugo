<script lang="ts">
	import '../app.css';
	import { page } from '$app/state';
	import { Button, ButtonGroup } from 'flowbite-svelte';
	export const prerender = true; // Build static pages
	let { children } = $props();

	let b1Style: string = '';
	let b2Style: string = '';
	if(page.url.pathname == '/') {
		b1Style = "text-slate-400"
	} else if(page.url.pathname == '/sechebdo') {
		b2Style = "text-slate-400"
	}
</script>

<div class="bg-no-repeat bg-cover min-h-48 pt-8 pb-8 text-white" style="background-position: 50%; background-image: url('/hero-blue.jpg');">
	<div class="ml-16">
		<h1 class="font-titres mb-1 text-5xl font-bold leading-relaxed [text-shadow:_1px_1px_2px_rgb(0_0_0_/_100%)]"><a href="/"><i class="inline-block w-12 h-12 mr-2 bg-cover bg-center" style="background-image: url('/favicon.png')"></i>Le Comptoir Sécu</a></h1>
		<h5 class="font-titres text-xl font-normal [text-shadow:_1px_1px_2px_rgb(0_0_0_/_100%)]">Parlions cybersécurité… autour d'un verre&nbsp;!</h5>
		<h5 class="font-titres font-normal [text-shadow:_1px_1px_2px_rgb(0_0_0_/_100%)]">Notre sagesse grandissante est toujours disponible sur <a href="https://discord.comptoirsecu.fr" class="text-yellow-400">Discord</a></h5>
	</div>
</div>
<div class="flex justify-center pt-4">
<ButtonGroup>
	<Button href="/" color="light" class={b1Style}>Épisodes thématiques</Button>
	<Button href="/sechebdo" color="light" class={b2Style}>SECHebdos</Button>
</ButtonGroup>
</div>

<div class="m-8">
{@render children()}
</div>
