<script lang="ts">
	export let ParsedUserObject: any;
	export let LogOutUser: any;
	let UserObjLen: any = Object.keys(ParsedUserObject).length;
</script>

<!--

    it is a temporary file, it will look different in the end, the logout funciton as will be in the navbar on the top

-->

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
	on:click={() => {
		if (UserObjLen === 0) {
			return null;
		} else {
			return LogOutUser();
		}
	}}
>
	{#if UserObjLen === 0}
		<a href="./login">Zaloguj się</a>
	{/if}

	{#if UserObjLen !== 0}
		Wyloguj się
	{/if}
</div>
