<script lang="ts">
	import type { PostDetails } from '../../../interfaces/PostDetails';
	export let post: PostDetails;
</script>

<div class="basis-1/4 bg-slate-300 rounded">
	<h2 class="font-bold text-3xl text-center">{post.title}</h2>
	<h3>Dodane przez: <span class="font-bold">{post.username}</span></h3>
	<p>Stworzone w: <span class="font-bold">{post.createdat}</span></p>
	<p>Wiadomość: <span class="font-bold">{post.message}</span></p>
</div>
