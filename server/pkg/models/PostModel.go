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

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"Password"`
}
