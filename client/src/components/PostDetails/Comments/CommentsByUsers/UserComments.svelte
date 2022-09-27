<script lang="ts">
	import type {
		AllDataComments,
		CommentPostDetails
	} from '../../../../interfaces/CommentsInterfaces/CommentPostinterface';

	export let CommentDetails: AllDataComments[] & CommentPostDetails[];
</script>

<div class="w-[35%] bg-slate-900 mt-36 ml-6">
	{#each CommentDetails as value, index}
		<div>{value.username}</div>
	{/each}
</div>
