<script lang="ts">
	import { onMount } from 'svelte';
	import type { UserInfo } from '../../interfaces/UserInfoLogin';
	import { Card, CardText } from 'svelte-materialify';
	import { createScene } from '../../scenes/scene';
	import type { PostDetails } from '../../interfaces/PostDetails';
	import MessageDeleteAllPostsOfUser from '../../components/addons/UserProfile/TemporaryComponents/MessageDeleteAllPostsOfUser.svelte';

	let Posts: PostDetails | any = [];
	let ProfileObj: UserInfo | any;
	let messageFromDeleteAllPosts: string = '';
	let showComp: boolean = true;

	let el: any;
	const deleteAllPostsOfUser: () => Promise<void> = async () => {
		showComp = true;
		await fetch(`http://localhost:8080/DeleteAllPostsOfUser?q=${ProfileObj.token}`, {
			method: 'POST',
			credentials: 'include',
			body: JSON.stringify(ProfileObj.email)
		})
			.then((res) => res.json())
			.then((data) => {
				messageFromDeleteAllPosts = data;
			});
		Posts = [];
	};

	onMount(async function () {
		ProfileObj = JSON.parse(localStorage.getItem('profile') || '{}');

		if (Object.keys(ProfileObj).length == 0) {
			console.log('NIEZALOGOWANY');
			location.href = '/';
		}

		const FetchSpecificUserPosts: () => Promise<void> = async () => {
			await fetch(`http://localhost:8080/FetchSpecificUserPosts?q=${ProfileObj.token}`, {
				method: 'POST',
				credentials: 'include',
				body: JSON.stringify(ProfileObj.email)
			})
				.then((res) => res.json())
				.then((data) => {
					Posts = data;
				});
		};
		FetchSpecificUserPosts();
	});

	onMount(() => {
		createScene(el);
	});
</script>

<div class="overflow-hidden flex flex-col justify-center relative mt-0 items-center top-0">
	<canvas bind:this={el} />
	<h1 class=" absolute text-5xl  text-white top-24">Stworzone Posty</h1>
	<div class="absolute text-white z-40 top-1/4 flex flex-wrap w-4/6 justify-center gap-2 ">
		{#each Posts as post}
			<Card hover outlined style="width:180px; height:30vh; overflow:hidden">
				<div class="basis-1/4 bg-slate-300 h-full relative rounded-xl">
					<a sveltekit:prefetch href="/posts/{post._id}">
						<div class="h-2/4 overflow-hidden">
							<img class=" top-0 left-0 h-full object-cover w-full" src={post.image} alt="" />
						</div>
						<h2 class=" text-left mb-3 pb-3 font-light text-sm">{post.title}</h2>
						<p class="text-overline text-gray-700 text-xs">
							Stworzone w: <span class="text-sm">{post.createdat}</span>
						</p>
					</a>
				</div>
			</Card>
		{/each}
	</div>
</div>
<button
	on:click={deleteAllPostsOfUser}
	class=" absolute top-40 left-96 z-30 text-red-500 cursor-pointer transition-colors duration-75 hover:text-red-800"
>
	Usuń Wszystkie Swoje Posty
</button>
{#if messageFromDeleteAllPosts != undefined && showComp}
	<MessageDeleteAllPostsOfUser message={messageFromDeleteAllPosts} bind:showComp />
{/if}
