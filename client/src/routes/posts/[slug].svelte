<script context="module" lang="ts">
	/** @type {import('./__types/[slug]').Load} */
	export async function load({
		params,
		fetch,
		session,
		stuff
	}: {
		params: any;
		fetch: any;
		session: any;
		stuff: any;
	}) {
		const response: Response = await fetch(`http://localhost:8080/getSinglePost?q=${params.slug}`);
		console.log(response);
		return {
			status: response.status,
			props: {
				article: response.ok && (await response.json())
			}
		};
	}
</script>

<script lang="ts">
	export let article: any;
	console.log(article);
</script>

<button> Kliknij na mnie </button>
