<script>
	export let data;
</script>

<svelte:head>
	<title>{data.post.title} | {data.siteName}</title>
	<meta
		name="description"
		content="{data.post.title} | {data.siteTagline}"
	/>
</svelte:head>

<div class="my-4 p-1">
	<h1 class="text-2xl font-bold">{data.post.title}</h1>

	{#if !data.post.isPage}
	<h3 class="text-m font-light">{(new Date(data.post.publishedAt)).toDateString()}</h3>
	{/if}

	<article class="my-4 prose max-w-none">
		{@html data.post.content}
	</article>
</div>