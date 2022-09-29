<script lang="ts">
	import { onMount } from 'svelte';
	import { POST } from '../../../constants/FetchDataMethods';
	import UserComments from './CommentsByUsers/UserComments.svelte';
	import CancelComment from './PublishComment/CancelComment.svelte';
	import PublishButton from './PublishComment/PublishButton.svelte';
	import type {
		AllDataComments,
		CommentPostDetails
	} from '../../../interfaces/CommentsInterfaces/CommentPostinterface';
	let valueOfComment: string = '';
	export let post: any;

	let CommentDetails: CommentPostDetails[] & AllDataComments[] = [];
	let RenderComps = false;
	const OnFocusComments: () => void = () => {
		RenderComps = true;
	};
	const OutFocusComments: () => void = () => {
		RenderComps = false;
	};
	onMount(async function () {
		//GetAllCommentsOfGivenPosts
		await fetch(`http://localhost:8080/GetAllCommentsOfGivenPosts`, {
			method: POST,
			body: JSON.stringify(post._id)
		})
			.then((res) => res.json())
			.then((data) => (CommentDetails = data));
	});
</script>

<section class="w-full flex justify-center flex-col">
	<div class="w-[90%]">
		<h2 class="text-2xl">Comments</h2>
		<div class="mt-4 w-1/2">
			<div class="border-b border-white w-[70%]">
				<input
					bind:value={valueOfComment}
					on:focus={OnFocusComments}
					class="mt-5 w-full bg-neutral-700 resize-none outline-none pb-1"
					type="text"
					placeholder="Napisz co o tym sądzisz"
				/>
			</div>

			<div class="relative flex gap-8 mt-2 w-[70%] justify-end">
				{#key RenderComps}
					{#if RenderComps}
						{#key valueOfComment}
							<CancelComment {OutFocusComments} />
							<PublishButton postDetailsId={post._id} {valueOfComment} {OutFocusComments} />
						{/key}
					{/if}
				{/key}
			</div>
		</div>
	</div>
	{#key CommentDetails}
		<UserComments {CommentDetails} />
	{/key}
</section>
