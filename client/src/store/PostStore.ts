import { writable, type Writable } from "svelte/store";

export const PostsFetched: Writable<any> = writable({})
const loading: Writable<boolean> = writable(false)
const error: Writable<boolean | any>  = writable(false)

export function PostsStoreHandler(url: string){
    async function get(){
        loading.set(true)
        error.set(false)
        try{const response = await fetch(url)
            PostsFetched.set(await response.json())
    } catch(e) {
        error.set(e)
    }
    loading.set(false)
    }
    get()
    return {PostsFetched, loading, error, get}
}