<script lang="ts">
	import { onMount } from 'svelte';
	import type { UserInfo } from '../../interfaces/UserInfoLogin';
	import { createScene } from '../../scenes/scene';
	import ProileInfoMainSite from '../../components/addons/UserProfile/ProfileInfoMainSite.svelte';
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
	<section class=" absolute w-screen h-screen flex flex-col justify-center items-center ">
		<div
			class="bg-white text-black w-1/3 flex justify-center flex-col items-center rounded-lg overflow-hidden p-5"
		>
			<h2 class="text-4xl">Profile Info</h2>
			<div class="flex  w-full mt-6">
				<div class="flex flex-col basis-1/2 gap-3">
					<ProileInfoMainSite title="Ustawienia" />
					<ProileInfoMainSite title="Dane konta" />
				</div>
				<div class="flex flex-col basis-1/2 gap-3">
					<ProileInfoMainSite title="Polityka Prywatności" />
					<ProileInfoMainSite title="Ustawienia" />
				</div>
			</div>
		</div>
	</section>
</div>
