package auth

import (
	"BackendGo/pkg/models"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateAccessToken(user *models.User) (string, error) {

	expirationTime := time.Now().Add(time.Hour * 24)
	tkToHash := &models.Token{
		UserID: user.UserId,
		Name:   user.Name,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tkToHash)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", fmt.Errorf("couldnt sign the token")
	}
	return tokenString, nil
}
