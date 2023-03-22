<script>
	import Header from '$lib/header/Header.svelte';
	import Footer from '$lib/footer/Footer.svelte';
	import '../output.css';

	export let data;

</script>

<div class="drawer">
  <input id="my-drawer-3" type="checkbox" class="drawer-toggle" /> 
  <div class="drawer-content flex flex-col">

<Header siteName={data.siteName} pages={data.pages}/>

<main>
	<div class="container mx-auto p-6">
		<slot />
	</div>
</main>

<Footer siteTagline={data.siteTagline} />

</div> 
  <div class="drawer-side">
    <label for="my-drawer-3" class="drawer-overlay"></label> 
    <ul class="menu p-4 w-80 bg-base-100">
		{#each data.pages as page}
  			<li><a href="/{page.slug}">{page.title}</a></li>
		{/each}
    </ul>
    
  </div>
</div>

<style>
</style>
