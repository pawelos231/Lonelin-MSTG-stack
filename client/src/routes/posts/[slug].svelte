<script context="module" lang="ts">
	/** @type {import('./__types/[slug]').Load} */
	export async function load({ params, fetch }: { params: any; fetch: any }): Promise<
		| {
				status: number;
				props: {
					post: any;
				};
				error?: undefined | any;
		  }
		| {
				status: number;
				error: Error;
				props?: undefined | any;
		  }
	> {
		const response: Response = await fetch(
			`http://localhost:8080/posts/getSinglePost?q=${params.slug}`
		);
		if (response.ok) {
			return {
				status: response.status,
				props: {
					post: await response.json()
				}
			};
		}
		return {
			status: response.status,
			error: new Error('Could not fetch the Posts')
		};
	}
</script>

<script lang="ts">
	import PostDetails from '../../components/PostDetails/PostDetails.svelte';
	export let post: any;
	import { onMount } from 'svelte';
	onMount(async function () {
		let respons: any = JSON.parse(localStorage.getItem('profile') || '{}');
		await fetch('http://localhost:8080/refreshToken', {
			method: 'POST',
			credentials: 'include',
			body: JSON.stringify(respons)
		})
			.then((res: Response) => res.json())
			.then((data: any) => {
				localStorage.setItem('profile', JSON.stringify(data.UserInfo));
			});
	});
</script>

<PostDetails {post} />
