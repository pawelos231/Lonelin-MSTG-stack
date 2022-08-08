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
				article: await response.json()
			}
		};
	}
</script>

<script lang="ts">
	export let article: any;
	console.log(article);
</script>

<h1 class="flex justify-center text-5xl m-8">{article.title}</h1>
<img src={article.image} alt={article.title} />
