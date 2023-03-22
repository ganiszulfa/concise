<script>
	export let data, siteName, siteTagline;

</script>

<svelte:head>
	<title>{siteName} - {siteTagline}</title>
	<meta
		name="description"
		content="{siteTagline}"
	/>
</svelte:head>

<section>
	{#each data.posts as post}
	<div class="card w-auto bg-neutral text-primary-content shadow-xl my-4">
		<div class="card-body">
		  <h2 class="card-title">{post.title}</h2>
		  <p>If a dog chews shoes whose shoes does he choose?</p>
		  <div class="card-actions justify-end">
			<a class="btn btn-primary" href="/{post.slug}">Read More</a>
		  </div>
		</div>
	</div>
	{:else}
		<h1>No posts</h1>
	{/each}
	<ul>
	</ul>
</section>

<style>
</style>
