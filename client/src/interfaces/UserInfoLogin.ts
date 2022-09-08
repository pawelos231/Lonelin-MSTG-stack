export interface UserInfo {
    email: string
    name: string
    token: string
}
export interface ErrorMessage  {
    Text: string
    Status: number
}

export interface StatusLogin {
    UserInfo: UserInfo
    Text: string,
    Status : number,
}

export interface UserLoginInfo {
    name :string,
    email :string,
    password :string,
    image: any
}