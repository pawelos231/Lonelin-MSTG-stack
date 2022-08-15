<script lang="ts">
	let name: string = '';
	let email: string = '';
	let password: string = '';
	import { fly } from 'svelte/transition';
	import { POST } from '../constants/FetchDataMethods';
	import type { StatusLogin, UserLoginInfo } from '../interfaces/UserInfoLogin';
	import { onMount } from 'svelte';
	import { createScene } from '../scene';
	let el: any;

	onMount(() => {
		createScene(el);
	});
	let Info: StatusLogin;
	const SendUserData = async (e: any) => {
		e.preventDefault();
		if (name != '' && email != '' && password != '') {
			localStorage.clear();
			const obj: UserLoginInfo = {
				name,
				email,
				password
			};
			await fetch('http://localhost:8080/createUser', {
				method: POST,
				body: JSON.stringify(obj)
			})
				.then((response) => response.json())
				.then((data) => {
					Info = data;
					if (Info.status == 1) {
						localStorage.setItem('profile', JSON.stringify(Info.UserInfo));
					}
				});
			const respons: string = JSON.parse(localStorage.getItem('profile') || '{}');
			console.log(respons);
		}
	};
</script>

{#if Info}
	{#if Info.status === 1}
		<p
			in:fly={{ y: -800, duration: 130, delay: 50 }}
			out:fly={{ duration: 100, delay: 50 }}
			class="absolute -translate-x-1/2 top-24 left-1/2 text-green-900"
		>
			{Info.text}
		</p>
	{/if}
	{#if Info.status === 0}
		<p
			in:fly={{ y: -800, duration: 130, delay: 50 }}
			out:fly={{ duration: 100, delay: 50 }}
			class="absolute -translate-x-1/2 top-24 left-1/2 text-red-900"
		>
			{Info.text}
		</p>
	{/if}
{/if}
<div class="flex flex-col items-center justify-center h-screen overflow-hidden">
	<canvas class="siema" bind:this={el} />
	<form
		class="bg-white opacity-90 h-1/3 p-5 m-5 rounded flex flex-col gap-4 w-1/5 items-center absolute"
		on:submit={(e) => SendUserData(e)}
	>
		<input
			bind:value={name}
			type="text"
			required
			placeholder="wprowadz nazwę"
			class=" w-4/5 bg-white p-2  transition duration-150 ease-linear hover:bg-black hover:text-white border-2 rounded-none"
		/>
		<input
			bind:value={email}
			type="email"
			required
			placeholder="wprowadz maila"
			class=" w-4/5 bg-white p-2 transition duration-150 ease-linear  hover:bg-black hover:text-white border-2 rounded-none"
		/>
		<input
			bind:value={password}
			type="password"
			required
			placeholder="wprowadz hasło"
			class=" w-4/5 bg-white p-2 transition duration-150 ease-linear  hover:bg-black hover:text-white border-2 rounded-none"
		/>
		<button
			type="submit"
			class="p-4 bg-black text-white rounded transition ease-linear duration-75 hover:scale-105 w-3/5"
			>Register</button
		>
	</form>
	<div class="absolute mt-96 text-white">
		<a href="/login"> I want to login </a>
	</div>
</div>
