<script>
	import { page } from '$app/stores';

	const postSlug = $page.params.postSlug;
	const postId = $page.params.postId;
	console.log(postSlug, postId);

	import { operationStore, query } from '@urql/svelte';

	const post = operationStore(
		`query ($id: Int!) {
            GetPost(id: $id) {
		        id
                title
				content
                slug
		        publishedAt
                author {
                    id
                    firstName
                    lastName
                    }
                }
            }`,
		{ id: postId }
	);

	query(post);
	console.log(post);
</script>

{#if $post.fetching}
	<p>Loading...</p>
{:else if $post.error}
	<p>Oh no... {$post.error.message}</p>
{:else}
	<h1>{post.data.GetPost.title}</h1>
	<h6><i>{new Date(post.data.GetPost.publishedAt).toDateString()}</i></h6>
	<div>{post.data.GetPost.content}</div>
{/if}
