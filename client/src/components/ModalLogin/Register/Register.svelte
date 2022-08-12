<script lang="ts">
	let name: string = '';
	let email: string = '';
	let password: string = '';
	import { POST } from '../../../constants/FetchDataMethods';
	import type { StatusLogin } from '../../../interfaces/PostDetails';
	let Info: StatusLogin;
	const SendUserData = async () => {
		const obj: object = {
			name,
			email,
			password
		};
		await fetch('http://localhost:8080/createUser', {
			method: POST,
			body: JSON.stringify(obj)
		})
			.then((response) => response.json())
			.then((data) => console.log((Info = data)));
	};
</script>

{#if Info}
	{#if Info.status === 1}
		<p class="text-green-900">{Info.text}</p>
	{/if}
	{#if Info.status === 0}
		<p class="text-red-900">{Info.text}</p>
	{/if}
{/if}
<form class="bg-slate-300 p-5 m-5 rounded flex flex-col gap-4">
	<input
		bind:value={name}
		type="text"
		placeholder="wprowadz nazwę"
		class="bg-white p-2 rounded-lg transition duration-150 ease-linear hover:bg-black hover:text-white"
	/>
	<input
		bind:value={email}
		type="text"
		placeholder="wprowadz maila"
		class="bg-white p-2 transition duration-150 ease-linear rounded-lg hover:bg-black hover:text-white"
	/>
	<input
		bind:value={password}
		type="text"
		placeholder="wprowadz hasło"
		class="bg-white p-2 transition duration-150 ease-linear rounded-lg hover:bg-black hover:text-white"
	/>
</form>
<button
	on:click={SendUserData}
	class="p-4 bg-black text-white rounded transition ease-linear duration-75 hover:scale-105"
	>Register</button
>
