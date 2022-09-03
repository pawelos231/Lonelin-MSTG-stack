<script lang="ts">
	import { onDestroy } from 'svelte';

	export let message: any;
	export let showComp, time;
	setTimeout(() => {
		showComp = false;
	}, time);
	onDestroy(() => {
		console.log('successfully destroyed Comp1');
	});
</script>

<h1 class="text-5xl fixed text-green-600 top-16">{message.text}</h1>
