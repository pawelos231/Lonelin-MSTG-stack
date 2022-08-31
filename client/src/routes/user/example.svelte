<script lang="ts">
	// @ts-nocheck

	const Onclick = async () => {
		const parsed = JSON.parse(localStorage.getItem('profile'));
		console.log(parsed.token);
		const res = await fetch(`http://localhost:8080/auth/userId?q=${parsed.token}`);
		let data = await res.json();
		console.log(data);
	};
</script>

<h1 on:click={() => Onclick()} class="flex justify-center text-black absolute top-12 left-96">
	Kliknij i wyślij requesta
</h1>
