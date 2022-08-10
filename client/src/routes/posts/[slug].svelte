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
				post: await response.json()
			}
		};
	}
</script>

<script lang="ts">
	import PostDetails from '../../components/PostDetails/PostDetails.svelte';
	export let post: any;
	console.log(post);
</script>

<svelte:head>
	<title>{post.title}</title>
</svelte:head>

<PostDetails {post} />
