<script lang="ts">
	let email: string = '';
	let password: string = '';
	let name: string = '';
	import { fly } from 'svelte/transition';
	import { POST } from '../constants/FetchDataMethods';
	import type { ErrorMessage } from '../interfaces/UserInfoLogin';
	let ErrorMess: ErrorMessage;
	const Login = async (e: any) => {
		e.preventDefault();
		if (email != '' && password != '') {
			let ObjectLogin: object = {
				name,
				email,
				password
			};
			await fetch('http://localhost:8080/loginUser', {
				method: POST,
				body: JSON.stringify(ObjectLogin)
			})
				.then((response) => response.json())
				.then((data) => console.log((ErrorMess = data)));
		}
	};
</script>

<div class="flex flex-col items-center justify-center h-screen">
	{#if ErrorMess}
		<p
			in:fly={{ y: -800, duration: 130, delay: 50 }}
			out:fly={{ duration: 100, delay: 50 }}
			class="absolute -translate-x-1/2 top-24 left-1/2 text-red-900"
		>
			{ErrorMess.text}
		</p>
	{/if}
	<form class="bg-slate-300 p-5 m-5 rounded flex flex-col gap-4 w-1/5" on:submit={(e) => Login(e)}>
		<input
			bind:value={email}
			type="text"
			placeholder="wprowadz Email"
			class="bg-white p-2 rounded-lg transition duration-150 ease-linear hover:bg-black hover:text-white"
		/>
		<input
			bind:value={password}
			type="text"
			placeholder="wprowadz hasło"
			class="bg-white p-2 transition duration-150 ease-linear rounded-lg hover:bg-black hover:text-white"
		/>
		<button
			type="submit"
			class="p-4 bg-black text-white rounded transition ease-linear duration-75 hover:scale-105"
			>Login</button
		>
	</form>
	<div>
		<a href="/register"> I want to register ! </a>
	</div>
</div>
