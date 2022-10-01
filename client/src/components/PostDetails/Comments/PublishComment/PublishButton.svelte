<script lang="ts">
	import { onMount } from 'svelte';
	import { POST } from '../../../../constants/FetchDataMethods';
	import type { CommentPostDetails } from '../../../../interfaces/CommentsInterfaces/CommentPostinterface';
	let profile: any;

	export let postDetailsId: any;
	export let valueOfComment: string;
	export let OutFocusComments: any;

	onMount(() => {
		profile = JSON.parse(localStorage.getItem('profile') || '{}');
	});

	const SendComment: (e: any) => Promise<void> = async (e: any) => {
		var today: Date = new Date();
		let date: string = today.getFullYear() + '-' + (today.getMonth() + 1) + '-' + today.getDate();

		//to change
		let Likes: number = 0;
		let ParentId: string = 'Nothing here';
		let NestedLevel: number = 0;
		let Repondsto: string = 'Nothing here, test value';
		let UpdatedAt: string = 'Nothing here, test value';

		//make interface for comments
		const CommentObject: CommentPostDetails = {
			likes: Likes,
			parentId: ParentId,
			nestedlevel: NestedLevel,
			respondto: Repondsto,
			updatedat: UpdatedAt,
			comment: valueOfComment,
			createdat: date,
			postid: postDetailsId
		};

		OutFocusComments();
		if (valueOfComment != '') {
			await fetch(`http://localhost:8080/comments/CommentOnPostByUser?q=${profile.token}`, {
				method: POST,
				credentials: 'include',
				body: JSON.stringify(CommentObject)
			})
				.then((res) => res.json())
				.then((data) => console.log(data));
		}
	};
</script>

<button on:click={SendComment} class="bg-neutral-800 basis-[20%] p-2 rounded hover:scale-95">
	Opublikuj
</button>
