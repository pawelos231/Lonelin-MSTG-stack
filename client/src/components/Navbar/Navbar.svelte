<script lang="ts">
	let ProfileObj: any = {};
	if (typeof localStorage !== 'undefined') {
		ProfileObj = JSON.parse(localStorage.getItem('profile') || '{}');
	}
</script>

<nav class="w-100 bg-slate-600 p-3 z-20">
	<ul
		class=" flex items-center  justify-end gap-12 mr-11
		child:transition duration-150 cursor-pointer 
		child-hover:text-blue-500 "
	>
		{#if Object.keys(ProfileObj).length !== 0}
			<li>
				<a href="/user/user"> {ProfileObj.name}</a>
			</li>
		{/if}
		<li>
			<a href="/"> Główna </a>
		</li>
		<li>siema</li>
		<li>siema</li>
		<li>siema</li>
	</ul>
</nav>
