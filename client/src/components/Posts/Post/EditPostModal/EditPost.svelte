<script lang="ts">
	//export let id;
	export let OpenModalEditHandler: any;
</script>

<div
	on:click={OpenModalEditHandler}
	class="w-screen h-screen flex justify-center items-center fixed top-0 z-50 backdrop-blur-md"
>
	<div class="h-2/3 w-1/3 bg-white rounded-sm flex items-center flex-col p-3">
		<h3 class="w-2/3">Edytuj Posta</h3>
		<div class="w-2/3 mt-5">
			<div>Tytuł</div>
			<div>Wiadomość</div>
		</div>
	</div>
</div>
