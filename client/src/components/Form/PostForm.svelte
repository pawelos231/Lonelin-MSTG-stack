<script lang="ts">
	export let HandleOnClick: any;
	export let Title: any;
	export let Message: any;
	export let onFileSelected: any;
</script>

<div
	class="w-screen h-screen fixed z-10 backdrop-blur-md flex justify-center
items-center"
>
	<form
		on:submit={(e) => HandleOnClick(e)}
		enctype="multipart/form-data"
		method="post"
		class="mb-10 flex flex-col w-1/5 p-4  bg-slate-600 gap-9 "
	>
		<input
			required
			class="border-2 rounded border-slate-300 border-dashed h-14"
			bind:value={Title}
			type="text"
			placeholder="wprowadz tytuł"
		/>
		<textarea
			required
			class="border-2 rounded border-slate-300 border-dashed h-14"
			bind:value={Message}
			type="text"
			placeholder="wprowadz wiadomość"
		/>
		<input
			required
			name="PostFile"
			type="file"
			class="bg-red w-50"
			accept=".jpg, .jpeg, .png"
			on:change={(e) => onFileSelected(e)}
		/>
		<button
			class="bg-slate-800 rounded p-3 cursor-pointer m-3 transition-all ease-in-out duration-300 hover:bg-black hover:text-white"
			type="submit"
		>
			Stwórz post
		</button>
	</form>
</div>
