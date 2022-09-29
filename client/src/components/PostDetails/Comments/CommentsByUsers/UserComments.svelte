<script lang="ts">
	import type {
		AllDataComments,
		CommentPostDetails
	} from '../../../../interfaces/CommentsInterfaces/CommentPostinterface';

	export let CommentDetails: AllDataComments[] & CommentPostDetails[];
</script>

<div class="w-[40%]  mt-36 ml-6">
	{#each CommentDetails as { username, createdat, comment }}
		<div class="flex mt-10">
			<div class=" bg-stone-100 basis-[1.5px] mr-5" />
			<div class="flex flex-col w-full">
				<div class="flex  ">
					<div class="basis-1/2  font-semibold">{username}</div>
					<div class="text-xs text-stone-300">{createdat}</div>
				</div>
				<div class="mt-2 text-sm">{comment}</div>
			</div>
		</div>
	{/each}
</div>
