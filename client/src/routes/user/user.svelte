<script lang="ts">
	import { onMount } from 'svelte';
	onMount(function () {
		let ProfileObj = JSON.parse(localStorage.getItem('profile') || '{}');
		console.log(ProfileObj);
		if (Object.keys(ProfileObj).length == 0) {
			console.log('NIEZALOGOWANY');
			location.href = '/';
		}
	});
</script>

<nav>
	<ul>
		<li>Główna</li>
		<li>Posty</li>
		<li>Komentarze</li>
		<li>Polubione posty</li>
		<li>Customize account</li>
		<li>Darkmode</li>
	</ul>
</nav>
