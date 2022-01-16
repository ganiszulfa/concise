<script>
	import Header from '$lib/header/Header.svelte';
	import { createClient, setClient } from '@urql/svelte';

	const client = createClient({
		url: 'http://localhost:8080/gql'
	});

	setClient(client);
</script>

<link rel="stylesheet" href="https://cdn.simplecss.org/simple.min.css" />

<Header />

<main>
	<slot />
</main>

<footer>
	<p>visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to learn SvelteKit</p>
</footer>
