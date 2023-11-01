<script>
	export let siteTagline;

	let now = new Date(),
		year;
	year = now.getFullYear();
</script>

<footer class="footer footer-center p-10 bg-neutral text-neutral-content">
	<div>
		<p>
			{siteTagline}
		</p>
		<p class="font-thin">Copyright © {year} - All right reserved</p>
	</div>
</footer>
