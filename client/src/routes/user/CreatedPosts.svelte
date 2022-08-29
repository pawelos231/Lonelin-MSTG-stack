<script lang="ts">
	import { onMount } from 'svelte';
	import type { UserInfo } from '../../interfaces/UserInfoLogin';
	import { createScene } from '../../scenes/scene';

	onMount(async function () {
		let ProfileObj: UserInfo | any = JSON.parse(localStorage.getItem('profile') || '{}');
		console.log(ProfileObj);
		if (Object.keys(ProfileObj).length == 0) {
			console.log('NIEZALOGOWANY');
			location.href = '/';
		}
		console.log(ProfileObj.email);
		const FetchUselessData = async () => {
			await fetch(`http://localhost:8080/FetchSpecificUserPosts?q=${ProfileObj.token}`, {
				method: 'POST',
				body: JSON.stringify(ProfileObj.email)
			})
				.then((res) => res.json())
				.then((data) => console.log(data));
		};
		FetchUselessData();
	});
	let el: any;

	onMount(() => {
		createScene(el);
	});
</script>

<div class="overflow-hidden flex flex-col justify-center relative mt-0">
	<canvas bind:this={el} />
	<button class="text-5xl absolute top-52 z-50 left-52 text-white">click</button>
</div>
