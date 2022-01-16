<script>
	import { operationStore, query } from '@urql/svelte';
	const posts = operationStore(
		`query {
            ListPage {
				id
                title
                slug
			}
        }`
	);
	query(posts);
	console.log(posts);
</script>

<header>
	<nav>
		<a href="/">Home</a>

		{#if !$posts.fetching && !$posts.error}
			{#each $posts.data.ListPage as post}
				<a href="/{post.slug}_{post.id}">{post.title}</a>
			{/each}
		{/if}

		<div class="corner">
			<!-- TODO put something else here? github link? -->
		</div>
	</nav>
</header>
