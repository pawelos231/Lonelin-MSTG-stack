<script lang="ts">
	import type { JSONObject } from '@sveltejs/kit/types/private';

	import { onMount } from 'svelte';

	let test: string[] = [];

	onMount(async function () {
		const response: Response = await fetch('http://localhost:8080/test');
		test = await response.json();
	});
	const HandleOnClick = async () => {
		const obj = {
			Title: 'siema',
			CreatedAt: '22002',
			Message: 'nie wiem',
			UserName: 'XD'
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
	<button on:click={HandleOnClick}> Kliknij na mnie </button>
</ul>
