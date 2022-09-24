package models

import jwt "github.com/dgrijalva/jwt-go"

//Token struct declaration
type Token struct {
	UserID string
	Name   string
	Email  string
	*jwt.StandardClaims
}

//later add role, gender, description of profile etc
type User struct {
	UserId   string `json:"userid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"Password"`
	Image    string `json:"Image"`
	Role     string `json:"Role"`
}
