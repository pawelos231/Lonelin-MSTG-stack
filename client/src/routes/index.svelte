<script lang="ts">
	import type { JSONObject } from '@sveltejs/kit/types/private';

	import { onMount } from 'svelte';

	let test: string[] = [];

	onMount(async function () {
		const response: Response = await fetch('http://localhost:8080/test');
		test = await response.json();
	});
	let name: string = '';
	let Title: string = '';
	let Message: string = '';
	const HandleOnClick = async () => {
		const obj = {
			Title: Title,
			CreatedAt: '22002',
			Message: Message,
			UserName: name
		};
		await fetch('http://localhost:8080/homeland', {
			method: 'POST',
			body: JSON.stringify(obj)
		})
			.then((response) => response.json())
			.then((data) => console.log(data));
	};
</script>

<h1>Combining SvelteKit and Golang server</h1>

<ul>
	<form>
		<input bind:value={Title} type="text" placeholder="wprowadz tytuł" />
		<input bind:value={Message} type="text" placeholder="wprowadz message" />
		<input bind:value={name} type="text" placeholder="wprowadź nazwe" />
	</form>
	<button on:click={HandleOnClick}> Kliknij na mnie </button>
</ul>
