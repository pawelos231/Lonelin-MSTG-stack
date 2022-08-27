package models

type PostInformiation struct {
	Title     string `json:"title"`
	CreatedAt string `json:"createdAt"`
	Image     string `json:"image"`
	Message   string `json:"message"`
	UserName  string `json:"userName"`
	Email     string `json:"email"`
}

type Comment struct {
	UserName    string `json:"username"`
	UserId      string `json:"userid"`
	CreatedAt   string `json:"createdat"`
	PostId      string `json:"postid"`
	Likes       string `json:"likes"`
	ParentId    string `json:"parentid"`
	NestedLevel string `json:"nestedlevel"`
	Repondsto   string `json:"respondsto"`
	UpdatedAt   string `json:"updateat"`
}
