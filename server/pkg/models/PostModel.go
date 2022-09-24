package models

type PostInformiation struct {
	Title     string   `json:"title"`
	CreatedAt string   `json:"createdAt"`
	Image     string   `json:"image"`
	Message   string   `json:"message"`
	UserName  string   `json:"userName"`
	Email     string   `json:"email"`
	Tags      []string `json:"tags"`
	Likes     []string `json:"likes"`
}

type Comment struct {
	UserName    string `json:"username"`
	UserId      string `json:"userid"`
	CreatedAt   string `json:"createdat"`
	PostId      string `json:"postid"`
	Likes       int    `json:"likes"`
	ParentId    string `json:"parentid"`
	NestedLevel int    `json:"nestedlevel"`
	Respondto   string `json:"respondto"`
	UpdatedAt   string `json:"updatedat"`
	Comment     string `json:"comment"`
}
