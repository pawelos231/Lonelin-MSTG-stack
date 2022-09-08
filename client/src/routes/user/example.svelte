<script lang="ts">
	// @ts-nocheck

	const Onclick = async () => {
		const respons: string = JSON.parse(localStorage.getItem('profile') || '{}');
		console.log(respons);
		await fetch('http://localhost:8080/refreshToken', {
			method: 'POST',
			credentials: 'include',
			body: JSON.stringify(respons)
		})
			.then((res) => res.json())
			.then((data) => console.log(data.UserInfo.token));
	};
</script>

<h1 on:click={() => Onclick()} class="flex justify-center text-black absolute top-12 left-96">
	Kliknij i wyślij requesta
</h1>
