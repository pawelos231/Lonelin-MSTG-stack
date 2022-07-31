<script lang="ts">
	import { onMount } from 'svelte';
	import type { PostDetails } from '../interfaces/PostDetails';
	let Posts: PostDetails[] = [];

	onMount(async function () {
		const response: Response = await fetch('http://localhost:8080/getdata');
		Posts = await response.json();
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
	<div>
		{#each Posts as post}
			<h2>{post.title}</h2>
			<h3>uzytkownik: {post.username}</h3>
			<p>Stworzone w: {post.createdat}</p>
			<p>Wiadomość: {post.message}</p>
		{/each}
	</div>
</ul>
