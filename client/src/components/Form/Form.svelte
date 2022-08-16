<script lang="ts">
	export let HandleOnClick: any;
	export let Title: any;
	export let Message: any;
	export let onFileSelected: any;
</script>

<form
	on:submit={(e) => HandleOnClick(e)}
	enctype="multipart/form-data"
	method="post"
	class="mb-10 flex flex-col w-1/5 p-4 gap-3 text-white"
>
	<input required class="border-4" bind:value={Title} type="text" placeholder="wprowadz tytuł" />
	<textarea
		required
		class="border-4"
		bind:value={Message}
		type="text"
		placeholder="wprowadz message"
	/>
	<input
		required
		name="PostFile"
		type="file"
		class="bg-red"
		accept=".jpg, .jpeg, .png"
		on:change={(e) => onFileSelected(e)}
	/>
	<button
		class="bg-slate-300 rounded p-3 cursor-pointer m-3 transition-all ease-in-out duration-300 hover:bg-black hover:text-white"
		type="submit"
	>
		Stwórz post
	</button>
</form>
