<script lang="ts">
	import UserComments from './CommentsByUsers/UserComments.svelte';
	import CancelComment from './PublishComment/CancelComment.svelte';
	import PublishButton from './PublishComment/PublishButton.svelte';

	let RenderComps: boolean = false;
	const OnFocusComments: () => void = () => {
		RenderComps = true;
	};
	const OutFocusComments: () => void = () => {
		RenderComps = false;
	};
</script>

<section class="w-full flex justify-center flex-col">
	<div class="w-[90%]">
		<h2 class="text-2xl">Comments</h2>
		<div class="mt-4 w-1/2">
			<div class="border-b border-white w-[70%]">
				<input
					on:focus={OnFocusComments}
					on:focusout={OutFocusComments}
					class="mt-5 w-full bg-neutral-700 resize-none outline-none pb-1"
					type="text"
					placeholder="Napisz co o tym sądzisz"
				/>
			</div>

			<div class="relative flex gap-8 mt-2 w-[70%] justify-end">
				<CancelComment {OutFocusComments} />
				<PublishButton />
			</div>
		</div>
	</div>
	<UserComments />
</section>
