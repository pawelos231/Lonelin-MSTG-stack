<script lang="ts">
	import Footer from '../components/Footer/Footer.svelte';
	import Navbar from '../components/Navbar/Navbar.svelte';
	import '../app.css';
	/*
	//find a way to check the current location with someting like checking it every time a location changes

	import { onMount } from 'svelte';
	let ifUserPage: boolean = false;
	onMount(function () {
		if (location.href.startsWith('http://localhost:3000/user')) {
			ifUserPage = true;
			console.log(ifUserPage);
		}
	});
	*/
</script>

<svelte:head>
	<title>Stronka</title>
	<link
		href="https://fonts.googleapis.com/css2?family=DM+Sans:ital,wght@0,400;0,500;1,400;1,500&family=Hahmlet:wght@100;200;300;500;600&family=Inter:wght@100;200;300;400;500;600;700;800;900&family=Lato:ital,wght@0,100;0,300;0,400;0,700;0,900;1,100;1,300;1,400;1,700;1,900&family=Poppins:wght@200;300;400;500;600;700;800;900&display=swap"
		rel="stylesheet"
	/>
</svelte:head>

<Navbar />

<main>
	<slot />
</main>
<Footer />

<style>
	:global(body, *) {
		margin: 0;
		padding: 0;
	}
</style>
