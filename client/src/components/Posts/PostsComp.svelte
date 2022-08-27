<script lang="ts">
	import type { PostDetails } from '../../interfaces/PostDetails';
	import type { UserInfo } from '../../interfaces/UserInfoLogin';
	import { PostStore } from '../../store/PostStore';
	import { onMount } from 'svelte';
	import Post from './Post/Post.svelte';
	import { POST } from '../../constants/FetchDataMethods';
	import { ProgressCircular, MaterialApp } from 'svelte-materialify';
	import PostForm from '../Form/PostForm.svelte';
	let OpenModalPostForm: boolean = false;

	const OpenModalPostFormHandler = () => {
		OpenModalPostForm = !OpenModalPostForm;
	};

	let Posts: PostDetails[] | any = [];
	onMount(async function () {
		const response: Response = await fetch('http://localhost:8080/getdata');

		//delay function to check if laoding state works
		PostStore.subscribe(async (data) => {
			data = await response.json();
			console.log(data);
			Posts = data;
		});
	});
	let parsed: any;
	onMount(async function () {
		parsed = JSON.parse(localStorage.getItem('profile') || '{}');
		console.log(parsed.token);
		const res = await fetch(`http://localhost:8080/auth/userId?q=${parsed.token}`);
		let data = await res.json();
		console.log(data);
	});
	let Title: string = '';
	let Message: string = '';
	let image: any;
	let ProfileObj: UserInfo | any = {};
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
		var today: any = new Date();
		let date: string = today.getFullYear() + '-' + (today.getMonth() + 1) + '-' + today.getDate();
		const obj: PostDetails = {
			title: Title,
			createdat: date,
			message: Message,
			image: image
		};
		await fetch(`http://localhost:8080/PostAPost?q=${ProfileObj.token}`, {
			method: POST,
			body: JSON.stringify(obj)
		});
		Posts = [...Posts, obj];
		console.log(Posts);
		ResetValuesOfForm();
		OpenModalPostForm = false;
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
	console.log(Posts);
</script>

<div class="mb-8 bg-black pt-15">
	{#if Object.keys(ProfileObj).length !== 0}
		{#if OpenModalPostForm}
			<PostForm
				{OpenModalPostFormHandler}
				{HandleOnClick}
				bind:Title
				bind:Message
				{onFileSelected}
			/>
		{/if}
		<button class="absolute w-48 bg-slate-400 p-3 m-5" on:click={OpenModalPostFormHandler}
			>Stwórz Post</button
		>
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

	<h1 class="text-5xl font-bold fa text-center bg-black p-9 text-white">Posty</h1>
	{#if Posts == null}
		<div class="m-20" style="display: flex; justify-content: center">
			<ProgressCircular indeterminate color="pink" size={80} />
		</div>
	{/if}
	{#if Posts?.length == 0}
		<div>Nic tu nie ma niestety</div>
	{/if}

	<div class="display flex gap-12 m-4 flex-wrap justify-center font-sans">
		{#if Posts !== null}
			{#each Posts as post}
				{#if post != null}
					<Post {post} {parsed} />
				{/if}
			{/each}
		{/if}
	</div>
</div>
