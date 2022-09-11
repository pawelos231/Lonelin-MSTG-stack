<script lang="ts">
	export let HandleOnClickModal: any;
	export let src: any;
</script>

<div
	class="w-full flex justify-center items-center h-full fixed top-0 backdrop-blur-md z-50"
	on:click={HandleOnClickModal}
>
	<img class="h-[95%]" {src} alt="" />
</div>
