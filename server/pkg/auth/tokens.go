package auth

import (
	"BackendGo/pkg/models"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func initializeToken(user *models.User, expirationTime time.Time) (*jwt.Token, error) {
	claims := &models.Token{
		UserID: user.UserId,
		Name:   user.Name,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	return token, nil
}

func CreateAccessToken(user *models.User) (string, error) {

	expirationTime := time.Now().Add(time.Minute * 15)
	token, _ := initializeToken(user, expirationTime)
	tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
	if err != nil {
		return "", fmt.Errorf("couldnt sign the token")
	}

	return tokenString, nil
}

func CreateRefreshToken(user *models.User) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 500)

	token, _ := initializeToken(user, expirationTime)
	tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
	if err != nil {
		return "", fmt.Errorf("couldnt sign the token")
	}

	return tokenString, nil
}
