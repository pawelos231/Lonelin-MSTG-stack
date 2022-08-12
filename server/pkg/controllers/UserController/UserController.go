package UserController

import (
	"BackendGo/pkg/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/dgrijalva/jwt-go"
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
		if req.Method == http.MethodPost {
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

			type LoginIfno struct {
				Text   string `json:"text"`
				Status int    `json:"status"`
			}

			if Error != nil {
				result, insertErr := col.InsertOne(ctx, user)
				RegisterErrors(result, insertErr)
				println(User)
				tkToHash := &models.Token{
					Name:  user.Name,
					Email: user.Email,
					StandardClaims: &jwt.StandardClaims{
						ExpiresAt: time.Now().Add(time.Minute * 100000).Unix(),
					},
				}

				token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tkToHash)
				println(token)
				LoginSuccesInfo := &LoginIfno{Text: "udało się utworzyć konto", Status: 1}

				//Later send here token

				json.NewEncoder(w).Encode(&LoginSuccesInfo)

			} else {
				println(User, Error, "udało się znaleźć innego uzytkoniwka o tym mailu")
				ErrorInfo := &LoginIfno{Text: "Podany email juz istnieje w bazie danych", Status: 0}
				json.NewEncoder(w).Encode(&ErrorInfo)
			}

		}
	}
}
