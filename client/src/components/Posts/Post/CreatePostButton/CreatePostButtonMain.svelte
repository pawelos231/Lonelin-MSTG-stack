<script lang="ts">
	export let OpenModalPostFormHandler: any;
	export let trimmedName: string;
</script>

<div class="w-full flex justify-center mt-10">
	<div
		class=" rounded-sm w-[45%] text-white p-3 m-5 z-10 border-2 bg-stone-800 ease-linear duration-100  cursor-pointer border-stone-500 hover:border-stone-600"
	>
		<div class="flex w-full h-full items-center">
			<p class="basis-[20%] text-center  mb-0 ease-linear duration-150 hover:text-stone-300">
				<a href="/user"> {trimmedName}</a>
			</p>
			<button
				class=" flex-grow bg-stone-600 easy-in duration-150 p-3 rounded-sm hover:bg-stone-900"
				on:click={OpenModalPostFormHandler}>Stwórz post</button
			>
			<div class="basis-[13%]" />
		</div>
	</div>
</div>
