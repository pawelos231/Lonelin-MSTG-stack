<script lang="ts">
	let email: string = '';
	let password: string = '';
	let name: string = '';
	const Login = async () => {
		if (email != '' && password != '') {
			let ObjectLogin: object = {
				name,
				email,
				password
			};
			await fetch('http://localhost:8080/loginUser', {
				method: 'POST',
				body: JSON.stringify(ObjectLogin)
			})
				.then((response) => response.json())
				.then((data) => console.log(data));
		}
	};
</script>

<form class="bg-slate-300 p-5 m-5 rounded flex flex-col gap-4">
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
</form>

<button
	on:click={Login}
	class="p-4 bg-black text-white rounded transition ease-linear duration-75 hover:scale-105"
	>Login</button
>
