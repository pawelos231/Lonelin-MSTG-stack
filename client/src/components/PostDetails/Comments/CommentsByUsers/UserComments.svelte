<script lang="ts">
	import { onMount } from 'svelte';
	import type {
		AllDataComments,
		CommentPostDetails
	} from '../../../../interfaces/CommentsInterfaces/CommentPostinterface';
	export let CommentDetails: AllDataComments[] & CommentPostDetails[] = [];
	console.log(CommentDetails);
</script>

<div class="w-[40%]  mt-36 ml-6">
	{#if CommentDetails == null}
		<div class="flex mt-10">
			<div class=" bg-stone-700 basis-[1.5px] mr-5">Ładuje</div>
		</div>
	{/if}
	{#if CommentDetails != null}
		{#if CommentDetails.length == 0}
			<div class="flex mt-10  justify-center text-xl absolute w-screen ml-0">
				<div class=" mr-5">Nie ma tu narazie zadnych komentarzy</div>
			</div>
		{/if}
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
	{/if}
</div>
