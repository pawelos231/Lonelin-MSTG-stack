<script lang="ts">
	export let href: string;
</script>

<a
	{href}
	class="text-black flex items-center justify-center bg-gray-200 border-0 px-8 focus:outline-none hover:bg-gray-300 rounded"
>
	<slot />
</a>
