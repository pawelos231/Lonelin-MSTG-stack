<script lang="ts">
	import MainPostComponent from '../components/Posts/MainPostsComponent.svelte';
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
	});
</script>

<div>
	<div>
		<MainPostComponent />
	</div>
</div>
