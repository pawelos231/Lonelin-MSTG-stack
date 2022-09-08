<script lang="ts">
	import PostsComp from '../components/Posts/PostsComp.svelte';
	import { onMount } from 'svelte';
	onMount(async function () {
		let respons: any = JSON.parse(localStorage.getItem('profile') || '{}');
		await fetch('http://localhost:8080/refreshToken', {
			method: 'POST',
			credentials: 'include',
			body: JSON.stringify(respons)
		})
			.then((res: Response) => res.json())
			.then((data: any) => {
				localStorage.setItem('profile', JSON.stringify(data.UserInfo));
			});
		respons = JSON.parse(localStorage.getItem('profile') || '{}');
		console.log(respons.token);
	});
</script>

<div>
	<div>
		<PostsComp />
	</div>
</div>
