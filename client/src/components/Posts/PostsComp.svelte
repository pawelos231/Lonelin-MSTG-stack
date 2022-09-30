<script lang="ts">
	import type { PostDetails } from '../../interfaces/PostDetails';
	import type { UserInfo } from '../../interfaces/UserInfoLogin';
	import { onMount } from 'svelte';
	import Post from './Post/Post.svelte';
	import { POST } from '../../constants/FetchDataMethods';
	import { ProgressCircular } from 'svelte-materialify';
	import { PostsStoreHandler } from '../../store/PostStore';
	import PostForm from '../Form/PostForm.svelte';
	let OpenModalPostForm: boolean = false;

	const OpenModalPostFormHandler = () => {
		OpenModalPostForm = !OpenModalPostForm;
	};

	const { PostsFetched, loading, error, get } = PostsStoreHandler(
		'http://localhost:8080/posts/getdata'
	);
	let Posts: PostDetails[] | any = [];

	let ParsedUserObject: UserInfo | any = {};
	onMount(async function () {
		ParsedUserObject = JSON.parse(localStorage.getItem('profile') || '{}');
	});
	onMount(async function () {
		PostsFetched.subscribe(async (data: any) => {
			if (Object.keys(data).length !== 0) {
				Posts = data;
			}
		});
	});

	//get user Profile info

	let Title: string = '';
	let Message: string = '';
	let image: any;

	//dodac tagi, sprawić by kliknięcie na stwórz post to był taki popup,

	const ResetValuesOfForm = (): void => {
		Title = '';
		Message = '';
		image = null;
	};

	const LogOut = (): void => {
		localStorage.clear();
		ParsedUserObject = {};
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

		await fetch(`http://localhost:8080/posts/PostAPost?q=${ParsedUserObject.token}`, {
			method: POST,
			credentials: 'include',
			body: JSON.stringify(obj)
		});

		const response: Response = await fetch('http://localhost:8080/posts/getdata');
		let data = await response.json();
		PostsFetched.update((PrevState) => data);
		PostsFetched.subscribe((value) => {
			Posts = value;
		});

		ResetValuesOfForm();
		OpenModalPostForm = false;
	};

	//select photo for your potst
	const onFileSelected: (e: any) => void = (e: any) => {
		//later change it, and add possibilty to add up to 10 images
		let ImageFromSelect: any = e.target.files[0];
		let reader: FileReader = new FileReader();
		console.log(ImageFromSelect);
		reader.readAsDataURL(ImageFromSelect);
		reader.onload = (e: any) => {
			image = e.target?.result;
		};
	};
</script>

<div class="mb-8 bg-black pt-15">
	{#if Object.keys(ParsedUserObject).length !== 0}
		{#if OpenModalPostForm}
			<PostForm
				{OpenModalPostFormHandler}
				{HandleOnClick}
				bind:Title
				bind:Message
				{onFileSelected}
			/>
		{/if}
		<button class="absolute w-48 bg-slate-400 p-3 m-5 z-10" on:click={OpenModalPostFormHandler}
			>Stwórz Post</button
		>
	{/if}

	{#if Object.keys(ParsedUserObject).length === 0}
		<p class="text-white">Zaloguj się by móc tworzyć quality posty</p>
	{/if}

	{#key ParsedUserObject}
		{#if Object.keys(ParsedUserObject).length === 0}
			<div
				class=" transition ease-in duration-200 cursor-pointer absolute top-20 right-20 bg-slate-300 p-4 rounded-md hover:bg-black hover:text-white"
			>
				<a href="/register"> Zaloguj się </a>
			</div>
		{/if}
		{#if Object.keys(ParsedUserObject).length !== 0}
			<div
				class=" transition ease-in duration-200 cursor-pointer absolute top-20 right-20 bg-slate-300 p-4 rounded-md hover:bg-black hover:text-white z-20"
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
	{#if $loading}
		<div class="m-20" style="display: flex; justify-content: center">
			<ProgressCircular indeterminate color="pink" size={80} />
		</div>
	{/if}
	{#key Posts}
		<div class="display flex gap-12 m-4 flex-wrap justify-center font-sans">
			{#each Posts as post, index}
				{#if post != null}
					<Post {post} {ParsedUserObject} {index} />
				{/if}
			{/each}
		</div>
	{/key}
</div>
