<script lang="ts">
	import { onMount } from 'svelte';
	import { createScene } from '../scene';
	let el: any;

	onMount(() => {
		createScene(el);
	});
</script>

<canvas class="siema" bind:this={el} />

<style>
	.siema {
		position: absolute;
	}
</style>
