<script lang="ts">
	export let value: string;
	export let link: string;
</script>

<li
	class="text-purple-600 ml-3 transition-colors duration-150 cursor-pointer hover:text-purple-300 hover:decoration-solid hover:underline z-20"
>
	<a href={link}>
		{value}
	</a>
</li>
