<script lang="ts">
	export let SwitchModal: any;
	import Register from './Register/Register.svelte';
	import Login from './Login/Login.svelte';
	let isRegister: boolean = false;
	const SetIsRegister = () => {
		isRegister = !isRegister;
		console.log(isRegister);
	};
</script>

<div
	class="fixed flex flex-col justify-center items-center text-2xl  top-0 left-0 w-screen h-screen bg-slate-400"
>
	Login Form
	<div class="absolute top-4 left-4 text-5xl" on:click={SwitchModal}>X</div>
	{#if isRegister}
		<Register />
	{/if}
	{#if !isRegister}
		<Login />
	{/if}
	<p class="cursor-pointer ease-linear duration-150 hover:text-blue-900" on:click={SetIsRegister}>
		{#if isRegister}
			I want to login
		{/if}
		{#if !isRegister}
			I want to register
		{/if}
	</p>
</div>
