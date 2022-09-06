package utils

import (
	"BackendGo/pkg/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type ErrorResponse struct {
	Err string
}

//make it more usuable later on
func HashValue(cost int, user *models.User) ([]byte, error) {
	var password, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
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
