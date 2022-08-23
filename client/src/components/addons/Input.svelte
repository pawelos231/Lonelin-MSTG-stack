<script lang="ts">
	export let value: string;
	export let label: string;
</script>

<input
	bind:value
	type="text"
	required
	placeholder={label}
	class="bg-white p-2 rounded-lg transition duration-150 ease-linear hover:bg-black hover:text-white"
/>
