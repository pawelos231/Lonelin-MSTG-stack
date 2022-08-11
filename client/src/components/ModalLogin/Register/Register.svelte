<script lang="ts">
	let name: string = '';
	let email: string = '';
	let password: string = '';
	import { POST } from '../../../constants/FetchDataMethods';
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
			.then((data) => console.log(data));
	};
</script>

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
