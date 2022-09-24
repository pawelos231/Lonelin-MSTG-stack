<script lang="ts">
	export let HandleOnClick: any;
	export let Title: any;
	export let Message: any;
	export let onFileSelected: any;
	export let OpenModalPostFormHandler: any;
</script>

<div
	class="w-screen h-screen fixed  backdrop-blur-md flex justify-center
items-center z-20  "
>
	<div
		on:click={OpenModalPostFormHandler}
		class="absolute top-4 left-4 text-5xl text-white cursor-pointer transition-all duration-100 hover:rotate-[360deg]"
	>
		X
	</div>
	<form
		on:submit={(e) => HandleOnClick(e)}
		enctype="multipart/form-data"
		method="post"
		class="mb-10 flex flex-col w-1/5 p-4  bg-slate-600 gap-9 rounded z-10 "
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
			class=" bg-slate-800 rounded p-3 cursor-pointer m-3 transition-all ease-in-out duration-300 hover:bg-black hover:text-white "
			type="submit"
		>
			Stwórz post
		</button>
	</form>
</div>
