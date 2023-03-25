<script>
	export let siteName, pages;
</script>

<header>
	<div class="w-full navbar bg-base-300">
      <div class="flex-none lg:hidden">
        <label for="my-drawer-3" class="btn btn-square btn-ghost">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="inline-block w-6 h-6 stroke-current"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path></svg>
        </label>
      </div> 
      <div class="flex-1 px-2 mx-2 font-bold text-2xl"><a href="/">{siteName}</a></div>
      <div class="flex-none hidden lg:block">
        <ul class="menu menu-horizontal">
		  {#each pages as page}
          	<li><a href="/{page.slug}">{page.title}</a></li>
		  {/each}
        </ul>
      </div>
    </div>
</header>

<style>
</style>
