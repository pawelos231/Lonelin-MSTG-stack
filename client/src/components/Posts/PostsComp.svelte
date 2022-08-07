<script lang="ts">
	import type { PostDetails } from '../../interfaces/PostDetails';
	import { PostStore } from '../../store/PostStore';
	import { onMount } from 'svelte';
	import Post from './Post/Post.svelte';
	import { POST } from '../../constants/FetchDataMethods';
	import { ProgressCircular, MaterialApp } from 'svelte-materialify';
	let Posts: PostDetails[] = [];
	onMount(async function () {
		const response: Response = await fetch('http://localhost:8080/getdata');
		function delay(time: number) {
			return new Promise((resolve) => setTimeout(resolve, time));
		}

		await delay(10000).then(() => console.log('ran after 1 second1 passed'));
		PostStore.subscribe(async (data) => {
			data = await response.json();
			console.log(data);
			Posts = data;
		});
	});
	let name: string = '';
	let Title: string = '';
	let Message: string = '';

	const ResetValuesOfForm = () => {
		name = '';
		Title = '';
		Message = '';
	};
	const HandleOnClick = async () => {
		if (name == '' || Title == '' || Message == '') {
			return;
		}
		const obj: PostDetails = {
			title: Title,
			createdat: '22002',
			message: Message,
			username: name
		};
		await fetch('http://localhost:8080/PostAPost', {
			method: POST,
			body: JSON.stringify(obj)
		});
		Posts = [...Posts, obj];
		ResetValuesOfForm();
	};
</script>

<div class="mb-8">
	<div>
		<form class="mb-10 flex flex-col w-1/5 m-4 gap-3">
			<input class="border-4" bind:value={Title} type="text" placeholder="wprowadz tytuł" />
			<input class="border-4" bind:value={Message} type="text" placeholder="wprowadz message" />
			<input class="border-4" bind:value={name} type="text" placeholder="wprowadź nazwe" />
		</form>
		<button
			class="bg-slate-300 rounded p-3 cursor-pointer m-3 transition-all ease-in-out duration-300 hover:bg-black hover:text-white"
			on:click={HandleOnClick}
		>
			Kliknij na mnie
		</button>
		<h1 class="text-5xl font-bold fa text-center m-9">Posty</h1>

		{#if Posts.length === 0}
			<div class="m-20" style="display: flex; justify-content: center">
				<ProgressCircular indeterminate color="pink" size={80} />
			</div>
		{/if}
		<div class="display flex gap-12 m-4 flex-wrap justify-center font-sans">
			{#each Posts as post}
				<Post {post} />
			{/each}
		</div>
	</div>
</div>
