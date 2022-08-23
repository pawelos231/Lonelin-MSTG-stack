<script context="module" lang="ts">
</script>

<script lang="ts">
	import Input from '../components/addons/Input.svelte';

	let email: string = '';
	let password: string = '';
	import { fly } from 'svelte/transition';
	import { POST } from '../constants/FetchDataMethods';
	import { onMount } from 'svelte';
	import { createScene } from '../scene';

	let el: any;

	onMount(() => {
		createScene(el);
	});
	export let SessionData: any;
	export async function Login(e: any) {
		localStorage.clear();
		e.preventDefault();
		if (email != '' && password != '') {
			let ObjectLogin: object = {
				email,
				password
			};
			await fetch('http://localhost:8080/loginUser', {
				method: POST,
				body: JSON.stringify(ObjectLogin)
			})
				.then((response) => response.json())
				.then((data) => console.log((SessionData = data)));
		}
		localStorage.setItem('profile', JSON.stringify(SessionData.UserInfo));
		console.log(SessionData.UserInfo);
		let ProfileObj: any;
		ProfileObj = JSON.parse(localStorage.getItem('profile') || '{}');
		console.log(ProfileObj);
		if (SessionData.status == 1) {
			location.href = '/';
		}
	}
</script>

<div class="flex flex-col items-center justify-center h-screen overflow-hidden">
	<canvas bind:this={el} />

	{#if SessionData}
		{#if SessionData.status == 0}
			<p
				in:fly={{ y: -800, duration: 130, delay: 50 }}
				out:fly={{ duration: 100, delay: 50 }}
				class=" text-2xl absolute -translate-x-1/2 top-24 left-1/2 text-red-900"
			>
				{SessionData.text}
			</p>
		{/if}
		{#if SessionData.status == 1}
			<p
				in:fly={{ y: -800, duration: 130, delay: 50 }}
				out:fly={{ duration: 100, delay: 50 }}
				class="absolute -translate-x-1/2 top-24 left-1/2 text-green-900 text-2xl"
			>
				{SessionData.text}
			</p>
		{/if}
	{/if}

	<form
		class="bg-slate-300 p-5 m-5 rounded flex flex-col gap-4 w-1/5 absolute"
		on:submit={(e) => Login(e)}
	>
		<Input label="wprowadz maila" bind:value={email} />
		<Input label="wprowadz hasło" bind:value={password} />

		<button
			type="submit"
			class="p-4 bg-black text-white rounded transition ease-linear duration-75 hover:scale-105"
		>
			Login
		</button>
	</form>
	<div class="absolute mt-96 text-white">
		<a href="/register"> I want to register </a>
	</div>
</div>
