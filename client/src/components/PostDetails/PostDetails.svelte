<script lang="ts">
	import type { PostDetails } from '../../interfaces/PostDetails';
	import Comments from './Comments/Comments.svelte';
	import ModalImage from './Modal/ModalImage.svelte';
	let ClickedModal: boolean = false;
	const HandleOnClickModal: () => void = () => {
		ClickedModal = !ClickedModal;
	};
	export let post: PostDetails | any;
</script>

<h1 class="flex justify-center text-5xl m-8 skew-x-3">{post.title}</h1>
<div class="flex flex-col items-center">
	<div class="w-1/2 pb-7 text-left">Created By: {post.username}</div>
	<div class=" w-5/5 flex justify-center select-none ">
		<div class="w-1/2  rounded overflow-hidden">
			<img
				on:click={HandleOnClickModal}
				class="basis-2/4 object-cover cursor-pointer"
				src={post.image}
				alt={post.title}
			/>
		</div>
	</div>
	<p class="p-10 text-lg w-3/5">{post.message}</p>
</div>
{#if ClickedModal}
	<ModalImage src={post.image} {HandleOnClickModal} />
{/if}

<Comments />
