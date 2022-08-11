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
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"Password"`
}

type Comment struct {
	UserName    string `json:username`
	UserId      string `json:userid`
	CreatedAt   string `json:createdat`
	PostId      string `json:postid`
	Likes       string `json:likes`
	parentId    string `json:parentid`
	NestedLevel string `json:nestedlevel`
	Repondsto   string `json:respondsto`
	updatedAt   string `json:updatedats`
}
