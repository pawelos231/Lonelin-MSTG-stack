<script lang="ts">
	import type { PostDetails } from '../../interfaces/PostDetails';
	import type { UserInfo } from '../../interfaces/UserInfoLogin';
	import { onMount } from 'svelte';
	import Post from './Post/Post.svelte';
	import { POST } from '../../constants/FetchDataMethods';
	import { ProgressCircular } from 'svelte-materialify';
	import { PostsStoreHandler } from '../../store/PostStore';
	import PostForm from '../Form/PostForm.svelte';
	import CreatePostButtonMain from './Post/CreatePostButton/CreatePostButtonMain.svelte';
	import LoginHandleButton from './Post/LoginHandleButtons/LoginHandleButton.svelte';
	import type { HandlePostFormInterface } from '../../interfaces/PostInterfaces/PostsHandleInterfaces';

	let OpenModalPostForm: boolean = false;
	const OpenModalPostFormHandler = () => {
		OpenModalPostForm = !OpenModalPostForm;
	};

	const { PostsFetched, loading, error, get } = PostsStoreHandler(
		'http://localhost:8080/posts/getdata'
	);

	let Posts: PostDetails[] | any = [];
	let ParsedUserObject: UserInfo | any = {};

	//data about post from the user
	let Title: string = '';
	let Message: string = '';
	let image: any;

	//for better code reading
	class HandlePostsForm implements HandlePostFormInterface {
		[x: string]: any;
		constructor(Title: string, Message: string, image: any) {
			this.Title = Title;
			this.Message = Message;
			this.image = image;
		}

		//reset form values
		ResetValuesOfForm() {
			this.Title = '';
			this.Message = '';
			this.image = null;
		}

		//logout user
		LogOut() {
			localStorage.clear();
			ParsedUserObject = {};
		}

		onFileSelected(e: any) {
			let ImageFromSelect: any = e.target.files[0];
			let reader: FileReader = new FileReader();
			reader.readAsDataURL(ImageFromSelect);
			reader.onload = (e: any) => {
				image = e.target?.result;
			};
		}

		GenerateDateString() {
			let today: Date = new Date();
			let date: string =
				today.getFullYear() +
				'-' +
				(today.getMonth() + 1 <= 9 ? '0' + Number(today.getMonth() + 1) : today.getMonth() + 1) +
				'-' +
				(today.getDate() <= 9 ? '0' + Number(today.getDate()) : today.getDate());
			return date;
		}
		//handler for submitting post forms
		async HandleSubmitPostForm(e: any) {
			e.preventDefault();
			let date: string = this.GenerateDateString();
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
			let data: any = await response.json();
			PostsFetched.update((PrevState: any) => data);
			PostsFetched.subscribe((value: any) => {
				Posts = value;
			});

			this.ResetValuesOfForm();
			OpenModalPostForm = false;
		}
	}

	const HandlePostFormObject: HandlePostsForm = new HandlePostsForm(Title, Message, image);

	const LogOutUser = () => HandlePostFormObject.LogOut();
	const HandleSubmitPostForm = (e: any) => HandlePostFormObject.HandleSubmitPostForm(e);
	const HandleFileSelect = () => (e: any) => HandlePostFormObject.onFileSelected(e);

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
</script>

<div class="mb-8 bg-black pt-15">
	{#if Object.keys(ParsedUserObject).length !== 0}
		{#if OpenModalPostForm}
			<PostForm
				{OpenModalPostFormHandler}
				{HandleSubmitPostForm}
				bind:Title
				bind:Message
				{HandleFileSelect}
			/>
		{/if}

		<CreatePostButtonMain {OpenModalPostFormHandler} trimmedName={ParsedUserObject.name.trim()} />
	{/if}

	{#if Object.keys(ParsedUserObject).length === 0}
		<p class="text-white">Zaloguj się by móc tworzyć quality posty</p>
	{/if}

	{#key ParsedUserObject}
		<LoginHandleButton {ParsedUserObject} {LogOutUser} />
	{/key}
	<!---
	<div class="text-white bg-slate-500 w-48 text-center p-2 rounded">Stwórz posta</div>
	-->

	<h1 class="text-5xl font-bold fa text-center bg-black mt-20 p-9 text-white">Posty</h1>

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
