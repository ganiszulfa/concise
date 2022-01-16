<script>
	import { page } from '$app/stores';

	const pageSlug = $page.params.pageSlug;
	const pageId = $page.params.pageId;
	console.log(pageSlug, pageId);

	import { operationStore, query } from '@urql/svelte';

	const post = operationStore(
		`query ($id: Int!) {
            GetPage(id: $id) {
		        id
                title
				content
                slug
                author {
                    id
                    firstName
                    lastName
                    }
                }
            }`,
		{ id: pageId }
	);

	query(post);
	console.log(post);
</script>

{#if $post.fetching}
	<p>Loading...</p>
{:else if $post.error}
	<p>Oh no... {$post.error.message}</p>
{:else}
	<h1>{$post.data.GetPage.title}</h1>
	<div>{$post.data.GetPage.content}</div>
{/if}
