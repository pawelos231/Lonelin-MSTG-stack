package UserController

import (
	"BackendGo/pkg/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type ErrorResponse struct {
	Err string
}

func RegisterErrors(result *mongo.InsertOneResult, insertErr error) {
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

func Register(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		user := &models.User{}
		json.NewDecoder(req.Body).Decode(&user)
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

		var User bson.M
		Error := col.FindOne(ctx, bson.M{"email": user.Email}).Decode(&User)
		if Error != nil {
			result, insertErr := col.InsertOne(ctx, user)
			RegisterErrors(result, insertErr)
			println(User)
			//Later send here token
			json.NewEncoder(w).Encode("logowanie powiodło się")

		} else {
			println(User, Error, "chuj")
			json.NewEncoder(w).Encode("Jest juz w bazie danych ten mail")
		}

	}
}
