<script lang="ts">
	export let ParsedUserObject: any;
	export let LogOut: any;
</script>

<!--

    it is a temporary file, it will look different in the end, the logout funciton as will be in the navbar on the top

-->

{#if Object.keys(ParsedUserObject).length === 0}
	<div
		class=" 
    text-white border-2 border-stone-500 
    w-[12%] 
    bg-stone-800
    z-20 
    text-center 
    transition ease-in duration-100 cursor-pointer 
    absolute top-20 right-20  p-4 
    rounded-md hover:bg-stone-900 
    hover:scale-105"
	>
		<a href="/register"> Zaloguj się </a>
	</div>
{/if}
{#if Object.keys(ParsedUserObject).length !== 0}
	<div
		class=" 
    text-white border-2 border-stone-500 
    w-[12%] 
    bg-stone-800
    z-20 
    text-center 
    transition ease-in duration-100 cursor-pointer 
    absolute top-20 right-20  p-4 
    rounded-md hover:bg-stone-900 
    hover:scale-105"
		on:click={LogOut}
	>
		Wyloguj się
	</div>
{/if}
