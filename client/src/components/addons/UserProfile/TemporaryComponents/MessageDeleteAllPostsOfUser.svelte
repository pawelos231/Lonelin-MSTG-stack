<script lang="ts">
	export let message: any;
	export let showComp: any;
	setTimeout(() => {
		showComp = false;
	}, 1600);
</script>

<h1 class="left-1/2 -translate-x-1/2 absolute top-10 p-5 text-2xl z-20 text-green-500  rounded-md">
	{message}
</h1>
