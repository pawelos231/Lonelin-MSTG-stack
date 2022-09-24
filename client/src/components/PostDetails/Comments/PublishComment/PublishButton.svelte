<script lang="ts">
	import { onMount } from 'svelte';
	import { POST } from '../../../../constants/FetchDataMethods';
	let profile: any;
	export let postDetailsId;
	export let valueOfComment: string;

	onMount(() => {
		profile = JSON.parse(localStorage.getItem('profile') || '{}');
	});
	const SendComment: (e: any) => Promise<void> = async (e: any) => {
		var today: any = new Date();
		let date: string = today.getFullYear() + '-' + (today.getMonth() + 1) + '-' + today.getDate();

		//to change
		let Likes: number = 0;
		let ParentId: string = 'fdsfdsfds';
		let NestedLevel: number = 100;
		let Repondsto: string = 'fdsfdsfds';
		let UpdatedAt: string = 'fdsfdsfds';

		//make interface for comments
		const PassObject: Object = {
			likes: Likes,
			parentId: ParentId,
			nestedlevel: NestedLevel,
			respondto: Repondsto,
			updatedat: UpdatedAt,
			comment: valueOfComment,
			createdat: date
		};

		await fetch(`http://localhost:8080/comments/CommentOnPostByUser?q=${profile.token}`, {
			method: POST,
			body: JSON.stringify(PassObject)
		})
			.then((res) => res.json())
			.then((data) => console.log(data));
	};
</script>

<button on:click={SendComment} class="bg-neutral-800 basis-[20%] p-2 rounded hover:scale-95">
	Opublikuj
</button>
