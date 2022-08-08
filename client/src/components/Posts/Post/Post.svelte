<script lang="ts">
	import type { PostDetails } from '../../../interfaces/PostDetails';
	import { Card, CardText } from 'svelte-materialify';
	export let post: PostDetails | any;
	console.log(post._id);
</script>

<Card hover outlined style="max-width:300px;">
	<a href="/posts/{post._id}">
		<div class="basis-1/4 bg-slate-300 rounded p-3">
			<h2 class="font-bold text-3xl text-center mb-3 p-6">{post.title}</h2>
			<h2 class="text-overline">Dodane przez: <span class="font-bold">{post.username}</span></h2>
			<p class="text-overline">Stworzone w: <span>{post.createdat}</span></p>
			<CardText>
				Lorem ipsum dolor sit amet consectetur adipisicing elit. Reprehenderit, qui quaerat rerum
				incidunt nisi ducimus?
			</CardText>
		</div>
	</a>
</Card>
