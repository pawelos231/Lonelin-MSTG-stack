package UserController

import (
	"BackendGo/pkg/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type ErrorResponse struct {
	Err string
}

func Register(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		user := &models.User{}
		json.NewDecoder(req.Body).Decode(&user)
		fmt.Println(user, "nie")
		var password, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
		if err != nil {
			fmt.Println(err)
			err := ErrorResponse{
				Err: "Password Encryption  failed",
			}
			json.NewEncoder(w).Encode(err)
		}

		dash := string(password)
		user.Password = dash
		result, insertErr := col.InsertOne(ctx, user)
		if insertErr != nil {
			fmt.Println("IntertOne Error", insertErr)
		} else {
			fmt.Println("insertOne() result type", reflect.TypeOf(result))
			fmt.Println("insertOne() api result type", result)
			newID := result.InsertedID
			fmt.Println("InsertedOne(), newID ", newID)
			fmt.Println("insertedOne(), newID type: ", reflect.TypeOf(newID))
		}
	}
}
