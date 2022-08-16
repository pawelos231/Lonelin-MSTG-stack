<script lang="ts">
	import type { PostDetails } from '../../interfaces/PostDetails';
	import { PostStore } from '../../store/PostStore';
	import { onMount } from 'svelte';
	import Post from './Post/Post.svelte';
	import { POST } from '../../constants/FetchDataMethods';
	import { ProgressCircular, MaterialApp } from 'svelte-materialify';
	import Form from '../Form/Form.svelte';
	let Posts: PostDetails[] | any = [];
	onMount(async function () {
		const response: Response = await fetch('http://localhost:8080/getdata');
		function delay(time: number) {
			return new Promise((resolve) => setTimeout(resolve, time));
		}
		//delay function to check if laoding state works
		PostStore.subscribe(async (data) => {
			data = await response.json();
			console.log(data);
			Posts = data;
		});
	});
	let name: string = '';
	let Title: string = '';
	let Message: string = '';
	let image: any;
	let ProfileObj: any = {};
	//dodac tagi, sprawić by kliknięcie na stwórz post to był taki popup,
	if (typeof localStorage !== 'undefined') {
		ProfileObj = JSON.parse(localStorage.getItem('profile') || '{}');
	}
	const ResetValuesOfForm = (): void => {
		Title = '';
		Message = '';
		image = null;
	};

	const LogOut = (): void => {
		localStorage.clear();
		ProfileObj = {};
	};

	const HandleOnClick = async (e: any): Promise<void> => {
		e.preventDefault();

		const obj: PostDetails = {
			title: Title,
			createdat: '22002',
			message: Message,
			image: image,
			username: ProfileObj.name
		};
		await fetch('http://localhost:8080/PostAPost', {
			method: POST,
			body: JSON.stringify(obj)
		});
		Posts = [...Posts, obj];
		console.log(Posts);
		ResetValuesOfForm();
	};
	const onFileSelected = (e: any) => {
		let ImageFromSelect: any = e.target.files[0];
		let reader: FileReader = new FileReader();
		console.log(ImageFromSelect);
		reader.readAsDataURL(ImageFromSelect);
		reader.onload = (e: any) => {
			image = e.target?.result;
		};
	};
</script>

<div class="mb-8 bg-black">
	{#if Object.keys(ProfileObj).length !== 0}
		<Form {HandleOnClick} bind:Title bind:Message {onFileSelected} />
	{/if}
	{#if Object.keys(ProfileObj).length === 0}
		<p class="text-white">Zaloguj się by móc tworzyć quality posty</p>
	{/if}

	{#key ProfileObj}
		{#if Object.keys(ProfileObj).length === 0}
			<div
				class=" transition ease-in duration-200 cursor-pointer absolute top-20 right-20 bg-slate-300 p-4 rounded-md hover:bg-black hover:text-white"
			>
				<a href="/register"> Zaloguj się </a>
			</div>
		{/if}
		{#if Object.keys(ProfileObj).length !== 0}
			<div
				class=" transition ease-in duration-200 cursor-pointer absolute top-20 right-20 bg-slate-300 p-4 rounded-md hover:bg-black hover:text-white"
				on:click={LogOut}
			>
				Wyloguj się
			</div>
		{/if}
	{/key}
	<!---
	<div class="text-white bg-slate-500 w-48 text-center p-2 rounded">Stwórz posta</div>
	-->

	<h1 class="text-5xl font-bold fa text-center m-9 text-white">Posty</h1>

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
