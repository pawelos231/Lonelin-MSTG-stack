<script lang="ts">
	import LiElement from '../../components/addons/UserProfile/liElement.svelte';
	interface LinksToUserDetailsPages {
		text: string;
		link: string;
	}
	const dataOfNavbarTextElements: LinksToUserDetailsPages[] = [
		{ text: 'Profile Info', link: '/user' },
		{ text: 'Posty', link: '/user/CreatedPosts' },
		{ text: 'Komentarze', link: '/user' },
		{ text: 'Polubione posty', link: '/user' },
		{ text: 'Customize account', link: '/user' },
		{ text: 'Darkmode', link: '/user' }
	];
</script>

<main>
	<nav class="h-[calc(100vh-40px)] absolute opacity-90 z-10 top-10">
		<ul class="flex flex-col gap-5 bg-gray-800 w-52 h-full pt-5">
			{#each dataOfNavbarTextElements as { text, link }}
				<LiElement value={text} {link} />
			{/each}
		</ul>
	</nav>
	<slot />
</main>
