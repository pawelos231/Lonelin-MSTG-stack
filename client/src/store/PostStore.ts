import { writable } from "svelte/store";
export const PostStore = writable([{
    title: "",
    username: "",
    createdat: "", 
    message: "",
}])
