<script lang="ts">
	let ProfileObj: any = {};
	if (typeof localStorage !== 'undefined') {
		ProfileObj = JSON.parse(localStorage.getItem('profile') || '{}');
	}
</script>

<nav class="w-full bg-neutral-900 p-5 z-20 fixed text-white">
	<ul
		class=" flex items-center  justify-end gap-12 mr-11
		child:transition duration-150 cursor-pointer 
		child-hover:text-blue-500 "
	>
		{#if Object.keys(ProfileObj).length !== 0}
			<li class="absolute left-5 flex">
				<a href="/user"> {ProfileObj.name}</a>
			</li>
		{/if}
		<li class="absolute left-1/2 -translate-x-1/2 w-1/4">
			<input
				class="bg-neutral-700 w-full p-2 rounded-sm text-white"
				type="text"
				placeholder="Szukaj po Lonlinie"
			/>
		</li>
		<li>
			<a href="/"> Główna </a>
		</li>
		<li class="relative">
			<a href="/games">
				Games
				<p id="smalltext" class="absolute -top-2 -right-5  text-red-700 ">Nowość</p>
			</a>
		</li>
	</ul>
</nav>

<style>
	#smalltext {
		font-size: 10px;
	}
</style>
