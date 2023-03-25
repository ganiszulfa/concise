<script>
	export let data;
</script>

<svelte:head>
	<title>{data.siteName} - {data.siteTagline} - Home</title>
	<meta
		name="description"
		content="{data.siteTagline}"
	/>
</svelte:head>

<section>
	{#each data.posts as post}
	<div class="w-auto shadow-xl my-4">
		<div class="card-body">
		  <h2 class="card-title">
			<a href="/{post.slug}">{post.title}</a>
		  </h2>
		  <p>{(new Date(post.publishedAt)).toDateString()}</p>
		</div>
	</div>
	{:else}
	<div class="text-center my-10">
		<h1>No More Posts...</h1>
	</div>
	{/each}

	<div class="btn-group my-8 flex justify-center">
		{#if data.page <= 1}
			<span class="btn btn-disabled">Prev</span>
		{/if}
		{#if data.page > 1}
			<a href="/?p={(data.page-1)}" class="btn">Prev</a>
		{/if}

		{#if data.posts.length > 0}
			<a href="/?p={(data.page+1)}" class="btn">Next</a>
		{/if}
		{#if data.posts.length <= 0}
			<span class="btn btn-disabled">Next</span>
		{/if}
	</div>
</section>

<style>
</style>
