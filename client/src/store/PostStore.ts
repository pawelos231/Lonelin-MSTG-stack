import { writable } from "svelte/store";
const PostStore= writable([{
    title: "",
    username: "",
    createdat: "", 
    message: "",
}])
export default PostStore