<script lang="ts">
	import type { PostDetails } from '../../../interfaces/PostDetails';
	import { Card, CardText } from 'svelte-materialify';
	import { POST } from '../../../constants/FetchDataMethods';
	import { PostsFetched } from '../../../store/PostStore';
	import EditPost from './EditPostModal/EditPost.svelte';
	import type { modifyUserPostsInterface } from '../../../interfaces/PostInterfaces/PostsHandleInterfaces';

	export let post: PostDetails | any;
	export let ParsedUserObject: any;
	export let index: any;

	let OpenModalEdit: boolean = false;
	let Temp: string = '';

	class ModifyUserPostsClass implements modifyUserPostsInterface {
		[x: string]: any;
		constructor(OpenModalEdit: boolean, Temp: string) {
			this.OpenModalEdit = OpenModalEdit;
			this.Temp = Temp;
		}

		//Check if the post needs to be stripped
		CheckIfPostDesIsToLong() {
			if (post.message.length > 150) {
				for (let i = 0; i <= 150; i++) {
					Temp += post.message[i];
				}
				Temp += '...';
			}
		}

		//Open modal to edit Post
		OpenModalEditHandler() {
			OpenModalEdit = !OpenModalEdit;
		}

		//remove given post Handler
		async removePostHandler() {
			await fetch(`http://localhost:8080/posts/DeletePost?q=${ParsedUserObject.token}`, {
				method: POST,
				credentials: 'include',
				body: JSON.stringify(post._id)
			})
				.then((response) => response.json())
				.then((data) => {
					console.log(data);
				});
			PostsFetched.update((data) => {
				data.splice(index, 1);
				return data;
			});
		}
	}

	const ModifyUserPostsClassHandler = new ModifyUserPostsClass(OpenModalEdit, Temp);

	const removePostHandlerFunction = () => ModifyUserPostsClassHandler.removePostHandler();

	const OpenModalEditHandler = () => ModifyUserPostsClassHandler.OpenModalEditHandler();

	ModifyUserPostsClassHandler.CheckIfPostDesIsToLong();
</script>

<Card hover outlined style="width:450px; height:70vh;">
	<div class="basis-1/4 bg-slate-300 h-full rounded p-3  relative">
		<a sveltekit:prefetch href="/posts/{post._id}">
			<div class="h-2/4 overflow-hidden">
				<img class=" top-0 left-0 h-full object-cover w-full" src={post.image} alt="" />
			</div>
			<h2 class=" text-3xl text-left mb-3 pb-6 font-light">{post.title}</h2>
			<h2 class="text-overline text-gray-700">
				Dodane przez: <span class="font-bold text-gray-700">{post.username}</span>
			</h2>
			<p class="text-overline text-gray-700">Stworzone w: <span>{post.createdat}</span></p>
		</a>
		<CardText>
			<p class="text-gray-600">
				{#if post.message.length > 150}
					{Temp}

					<div class=" pt-5 absolute bottom-2 left-5 text-slate-900">Czytaj więcej</div>
				{/if}
				{#if post.message.length < 150}
					{post.message}

					<div class="absolute bottom-3 left-5 text-slate-900">Czytaj więcej</div>
				{/if}
			</p>
			{#if post.email == ParsedUserObject.email}
				<div class="flex flex-col absolute right-10 bottom-0 bg-slate-300 p-2 ">
					<p
						on:click={OpenModalEditHandler}
						class="p-2 text-green-800 hover:text-yellow-200 hover:bg-slate-500 rounded"
					>
						Edytuj Posta
					</p>
					<p
						on:click={removePostHandlerFunction}
						class="text-red-600 hover:text-yellow-200 hover:bg-slate-500 p-2 rounded"
					>
						Usuń posta
					</p>
				</div>
			{/if}
		</CardText>
	</div>
</Card>

{#if OpenModalEdit}
	<EditPost {OpenModalEditHandler} />
{/if}
