//for the time being it will be only refresh token validator, later it will be able to validate

package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func ValidateToken(w http.ResponseWriter, req *http.Request) bool {
	tokenCookie2, errCookie := req.Cookie("jid")
	if errCookie != nil {
		fmt.Println(errCookie)
		return false
	}

	value := tokenCookie2.Value

	tkClaims := jwt.MapClaims{}
	refreshToken, errParsed := jwt.ParseWithClaims(value, tkClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("REFRESH_TOKEN_SECRET")), nil
	})
	if errParsed != nil || refreshToken == nil {
		if errParsed == jwt.ErrSignatureInvalid {
			json.NewEncoder(w).Encode("invalid token")
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("zły token")
		println("blad")
		return false
	}
	return true
}
