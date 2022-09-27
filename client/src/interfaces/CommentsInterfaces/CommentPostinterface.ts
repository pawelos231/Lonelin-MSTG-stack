type CommentInfo = {
    likes: number,
    parentId: string,
    nestedlevel: number, 
    respondto: string,
    updatedat: string,
    comment: string,
    createdat: string,
    postid: string,
}

type UserData = {
    _id: string,
    username: string,
}
export type AllDataComments = CommentInfo & UserData
export type CommentPostDetails =  CommentInfo | UserData 
