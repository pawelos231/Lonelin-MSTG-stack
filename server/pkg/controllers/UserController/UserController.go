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
	"github.com/google/uuid"
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
			id := uuid.New().String()
			dash := string(password)
			user.Password = dash
			user.UserId = id

			var User bson.M
			Error := col.FindOne(ctx, bson.M{"email": user.Email}).Decode(&User)

			type LoginIfno struct {
				Text   string `json:"text"`
				Status int    `json:"status"`
			}

			if Error != nil {

				result, insertErr := col.InsertOne(ctx, user)
				RegisterErrors(result, insertErr)
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
					println(err)
				}

				var UserInfo = map[string]interface{}{}
				UserInfo["token"] = tokenString
				UserInfo["email"] = user.Email
				UserInfo["name"] = user.Name

				var reponse = map[string]interface{}{"UserInfo": UserInfo}
				reponse["text"] = "Udało się zalogować !"
				reponse["status"] = 1

				json.NewEncoder(w).Encode(reponse)

			} else {
				println(User, Error, "udało się znaleźć innego uzytkoniwka o tym mailu")
				var resp = map[string]interface{}{"status": false, "text": "Podany uzytkownik juz istnieje w bazie"}
				resp["status"] = 0
				json.NewEncoder(w).Encode(resp)
			}

		} else {
			json.NewEncoder(w).Encode("U used wrong method")
		}
	}
}

func Login(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			user := &models.User{}
			json.NewDecoder(req.Body).Decode(&user)
			var UserData = &models.User{}
			Info := col.FindOne(ctx, bson.M{"email": user.Email}).Decode(&UserData)
			fmt.Println(Info)

			if Info == nil {
				errorPass := bcrypt.CompareHashAndPassword([]byte(UserData.Password), []byte(user.Password))
				if errorPass != nil && errorPass == bcrypt.ErrMismatchedHashAndPassword {
					var ErrorInfo = map[string]interface{}{}
					ErrorInfo["status"] = 0
					ErrorInfo["text"] = "Niepoprawne Hasło"
					json.NewEncoder(w).Encode(ErrorInfo)
				} else {
					expirationTime := time.Now().Add(time.Hour * 24)
					tkToHash := &models.Token{
						UserID: UserData.UserId,
						Name:   UserData.Name,
						Email:  UserData.Email,
						StandardClaims: &jwt.StandardClaims{
							ExpiresAt: expirationTime.Unix(),
						},
					}
					token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tkToHash)
					tokenString, err := token.SignedString([]byte("secret"))
					if err != nil {
						println(err)
					}
					var UserInfo = map[string]interface{}{}
					UserInfo["token"] = tokenString
					UserInfo["email"] = UserData.Email
					UserInfo["name"] = UserData.Name
					var reponse = map[string]interface{}{"UserInfo": UserInfo}
					reponse["text"] = "Udało się zalogować !"
					reponse["status"] = 1
					json.NewEncoder(w).Encode(reponse)

				}
			} else {
				fmt.Println("nie ma takiego typa", Info)
				var ErrorInfo = map[string]interface{}{}
				ErrorInfo["status"] = 0
				ErrorInfo["text"] = "Nie ma takiego chłopa w bazie danych"
				json.NewEncoder(w).Encode(ErrorInfo)
			}
		} else {
			json.NewEncoder(w).Encode("U used wrong method")
		}
	}
}

func FetchUserById(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			fmt.Println("siemaController")
			data := req.Context().Value("user")
			fmt.Println(data)
			json.NewEncoder(w).Encode(data)
		}
	}
}
