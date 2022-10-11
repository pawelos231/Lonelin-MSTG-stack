package utils

import (
	"BackendGo/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type ErrorResponse struct {
	Err string
}


//make it more usuable later on
func HashValue(cost int, user *models.User) ([]byte, error) {
	var password, err = bcrypt.GenerateFromPassword([]byte(user.Password), cost)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Password Encryption  failed",
		}
		fmt.Println(err)
		return nil, fmt.Errorf("coś poszło nie tak w hashowaniu")
	}
	return password, nil
}

func CompareHashAndPassword(hashedPasword string, password string, w http.ResponseWriter) error {
	errorPass := bcrypt.CompareHashAndPassword([]byte(hashedPasword), []byte(password))
	if errorPass != nil && errorPass == bcrypt.ErrMismatchedHashAndPassword {
		var ErrorInfo = map[string]interface{}{}
		ErrorInfo["status"] = 0
		ErrorInfo["text"] = "Niepoprawne Hasło"
		json.NewEncoder(w).Encode(ErrorInfo)
		return errorPass
	}
	return nil
}
