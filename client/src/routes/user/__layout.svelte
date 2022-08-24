<script lang="ts">
	import LiElement from '../../components/addons/UserProfile/liElement.svelte';
	const dataOfNavbarTextElements: Array<string> = [
		'Główna',
		'Posty',
		'Komentarze',
		'Polubione Posty',
		'Customize Account',
		'DarkMode'
	];
</script>

<main>
	<nav class="h-screen absolute opacity-90 z-10">
		<ul class="flex flex-col gap-5 bg-gray-800 w-52 h-full pt-5">
			{#each dataOfNavbarTextElements as element}
				<LiElement value={element} />
			{/each}
		</ul>
	</nav>
	<slot />
</main>
