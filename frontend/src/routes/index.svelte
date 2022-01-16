<script>
	import { operationStore, query } from '@urql/svelte';

	const limitInit = 2;

	const posts = operationStore(
		`query ($limit: Int!, $page: Int!) {
            ListPost(isPublished:true, limit: $limit, page: $page) {
		        id
                title
                slug
		        publishedAt
                author {
                    id
                    firstName
                    lastName
                    }
                }
            }`,
		{ limit: limitInit, page: 1 }
	);

	query(posts);

	function more() {
		$posts.variables = {
			...$posts.variables,
			limit: $posts.variables.limit + limitInit
		};
		$posts.reexecute();
	}
</script>

{#if $posts.fetching}
	<p>Loading...</p>
{:else if $posts.error}
	<p>Oh no... {$posts.error.message}</p>
{:else}
	<ul>
		{#each $posts.data.ListPost as post}
			<li>
				<a href="/{post.slug}/{post.id}">
					<b>{post.title}</b> -
					<i>{new Date(post.publishedAt).toDateString()}</i>
				</a>
			</li>
		{/each}
	</ul>

	<button on:click={more}>More</button>
{/if}
