<script lang="ts">
	import type { PostDetails } from '../../../interfaces/PostDetails';
	import { Card, CardText } from 'svelte-materialify';
	export let post: PostDetails | any;
	console.log(post._id);
</script>

<Card hover outlined style="max-width:300px;">
	<a href="/posts/{post._id}">
		<div class="basis-1/4 bg-slate-300 rounded p-3  relative">
			<img class=" top-0 left-0" src={post.image} alt="" />
			<h2 class=" text-3xl text-left mb-3 pb-6 font-light">{post.title}</h2>
			<h2 class="text-overline text-gray-700">
				Dodane przez: <span class="font-bold text-gray-700">{post.username}</span>
			</h2>
			<p class="text-overline text-gray-700">Stworzone w: <span>{post.createdat}</span></p>
			<CardText>
				<p class="text-gray-600">
					{post.message}
				</p>
			</CardText>
		</div>
	</a>
</Card>
