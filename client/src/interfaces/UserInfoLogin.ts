export interface UserInfo {
    email: string
    name: string
    token: string
}
export interface ErrorMessage  {
    text: string
    status: number
}

export interface StatusLogin {
    UserInfo: UserInfo
    text: string,
    status : number,
}

export interface UserLoginInfo {
    name :string,
    email :string,
    password :string,
}