<script lang="ts">
	// @ts-nocheck
	import { mode } from '$app/env';

	const Onclick = async () => {
		const parsed = JSON.parse(localStorage.getItem('profile'));
		console.log(parsed.token);
		const res = await fetch('http://localhost:8080/auth/userId', {
			headers: { 'access-token': parsed.token, 'Access-Control-Allow-Origin': '*' }
		});

		console.log(res);
	};
</script>

<h1 on:click={() => Onclick()} class="flex justify-center">Kliknij i wyślij requesta</h1>
