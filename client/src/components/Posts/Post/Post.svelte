<script lang="ts">
	import type { PostDetails } from '../../../interfaces/PostDetails';
	import { Card, CardText } from 'svelte-materialify';
	export let post: PostDetails | any;
	export let parsed: any;
	console.log(parsed.email);
	console.log(post.email);
	let Temp: string = '';
	if (post.message.length > 150) {
		for (let i = 0; i <= 150; i++) {
			Temp += post.message[i];
		}
		Temp += '...';
	}
</script>

<Card hover outlined style="width:450px; height:70vh;">
	<a sveltekit:prefetch href="/posts/{post._id}">
		<div class="basis-1/4 bg-slate-300 h-full rounded p-3  relative">
			<div class="h-2/4 overflow-hidden">
				<img class=" top-0 left-0 h-full object-cover w-full" src={post.image} alt="" />
			</div>
			<h2 class=" text-3xl text-left mb-3 pb-6 font-light">{post.title}</h2>
			<h2 class="text-overline text-gray-700">
				Dodane przez: <span class="font-bold text-gray-700">{post.username}</span>
			</h2>
			<p class="text-overline text-gray-700">Stworzone w: <span>{post.createdat}</span></p>
			<CardText>
				<p class="text-gray-600">
					{#if post.message.length > 150}
						{Temp}
						<div class="absolute bottom-3 left-5 text-slate-900">Czytaj więcej</div>
					{/if}
					{#if post.message.length < 150}
						{post.message}
						<div class="absolute bottom-3 left-5 text-slate-900">Czytaj więcej</div>
					{/if}
				</p>
				{#if post.email == parsed.email}
					<div class="flex flex-col absolute right-10 bottom-0 ">
						<p class="p-2 text-green-800 hover:text-yellow-200 hover:bg-slate-500 rounded">
							Edytuj Posta
						</p>
						<p class="text-red-600 hover:text-yellow-200 hover:bg-slate-500 p-2 rounded">
							Usuń posta
						</p>
					</div>
				{/if}
			</CardText>
		</div>
	</a>
</Card>
