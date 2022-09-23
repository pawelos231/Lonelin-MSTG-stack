<script lang="ts">
	import { onMount } from 'svelte';
	import { POST } from '../../../../constants/FetchDataMethods';
	let profile: any;
	onMount(() => {
		profile = JSON.parse(localStorage.getItem('profile') || '{}');
		console.log(profile);
	});
	const SendComment: () => Promise<void> = async () => {
		await fetch(`http://localhost:8080/comments/CommentOnPostByUser?q=${profile.token}`, {
			method: POST,
			body: JSON.stringify({ name: 'siema', wiadomosc: 'nic', niewiem: 'gowno' })
		});
	};
</script>

<button on:click={SendComment} class="bg-neutral-800 basis-[20%] p-2 rounded hover:scale-95">
	Opublikuj
</button>
