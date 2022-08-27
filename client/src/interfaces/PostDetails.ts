type PostInfo = {
    title: string,
    image: any,
    createdat: string, 
    message: string
}

type specificData = {
    _id: string,
    username: string,
}
export type PostDetails = PostInfo | specificData