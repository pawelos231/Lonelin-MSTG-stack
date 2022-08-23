package models

import jwt "github.com/dgrijalva/jwt-go"

type PostInformiation struct {
	Title     string `json:"title"`
	CreatedAt string `json:"createdAt"`
	Image     string `json:"image"`
	Message   string `json:"message"`
	UserName  string `json:"userName"`
	Email     string `json:"email"`
}

//later add more things like: gender, createdPosts, liked comments, some kind of points etc.
//if this file becomes too big, move it to different smaller files

//later add role, gender, description of profile etc
type User struct {
	UserId   string `json:"userid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"Password"`
	Image    string `json:"Image"`
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

//Token struct declaration
type Token struct {
	UserID any
	Name   string
	Email  string
	*jwt.StandardClaims
}
