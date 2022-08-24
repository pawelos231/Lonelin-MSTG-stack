<script lang="ts">
	import { onMount } from 'svelte';
	import type { UserInfo } from '../../interfaces/UserInfoLogin';
	import { createScene } from '../../scenes/scene';
	onMount(function () {
		const ProfileObj: Object | UserInfo = JSON.parse(localStorage.getItem('profile') || '{}');
		console.log(ProfileObj);
		if (Object.keys(ProfileObj).length == 0) {
			console.log('NIEZALOGOWANY');
			location.href = '/';
		}
	});
	let el: any;
	onMount(() => {
		createScene(el);
	});
</script>

<div class="overflow-hidden flex flex-col justify-center relative mt-0">
	<canvas bind:this={el} />
</div>
