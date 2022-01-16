<script>
	import { createClient, setClient } from '@urql/svelte';

	const client = createClient({
		url: 'http://localhost:8080/gql'
	});

	setClient(client);
</script>
