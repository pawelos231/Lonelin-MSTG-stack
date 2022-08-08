package models

func SayHello() string {
	return "Hi from package dir1"
}

type PostInformiation struct {
	Title     string `json:"title"`
	CreatedAt string `json:"createdAt"`
	Image     string `json:"image"`
	Message   string `json:"message"`
	UserName  string `json:"userName"`
}

type UserLoginInfo struct {
	Username string `json:"userName"`
	Email    string `json:"Email"`
	Image    string `json:"Image"`
	Password string `json:"Password"`
}
