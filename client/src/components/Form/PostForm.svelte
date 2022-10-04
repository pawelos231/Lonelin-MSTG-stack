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
	<div class="bg-stone-800 w-[40%] h-[50%] h border-2 rounded-sm">
		<form
			on:submit={(e) => HandleOnClick(e)}
			enctype="multipart/form-data"
			method="post"
			class=" w-full h-full mb-10 flex flex-col  rounded-sm  p-4   gap-9  z-10  relative"
		>
			<div class="border-2 rounded border-stone-500 text-white">
				<input
					required
					class=" outline-none h-14 p-2"
					bind:value={Title}
					type="text"
					placeholder="wprowadz tytuł"
				/>
			</div>

			<div class="w-full border-2 rounded border-stone-500 text-white">
				<textarea
					required
					class="0  h-14 w-full outline-none p-2"
					bind:value={Message}
					type="text"
					placeholder="wprowadz wiadomość"
				/>
			</div>
			<div class="flex justify-center">
				<input
					required
					name="PostFile"
					type="file"
					class=" w-50"
					accept=".jpg, .jpeg, .png"
					on:change={(e) => onFileSelected(e)}
				/>
			</div>
			<div
				class="bg-stone-800 border-2 border-stone-500 text-white   rounded p-3 cursor-pointer m-3 transition-all ease-in-out duration-100 hover:bg-stone-900 hover:text-white flex justify-center bottom-6 absolute w-[90%]"
			>
				<button type="submit"> Stwórz post </button>
			</div>
		</form>
	</div>
</div>
