package models

func SayHello() string {
	return "Hi from package dir1"
}

type PostInformiation struct {
	Title     string `json:"title"`
	CreatedAt string `json:"createdAt"`
	Message   string `json:"message"`
	UserName  string `json:"userName"`
}

type UserLoginInfo struct {
	Username string `json:"userName"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
