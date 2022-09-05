package UserController

import (
	"BackendGo/pkg/auth"
	"BackendGo/pkg/models"
	"BackendGo/pkg/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type ErrorResponse struct {
	Err string
}

func Register(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			json.NewEncoder(w).Encode("U used wrong method")
			return
		}

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

		if Error == nil {
			println(User, Error, "udało się znaleźć innego uzytkoniwka o tym mailu")
			var resp = map[string]interface{}{"status": false, "text": "Podany uzytkownik juz istnieje w bazie"}
			resp["status"] = 0
			json.NewEncoder(w).Encode(resp)
			return
		}

		result, insertErr := col.InsertOne(ctx, user)
		utils.RegisterInsertErrors(result, insertErr)

		tokenString, _ := auth.CreateAccessToken(user)

		var UserInfo = map[string]interface{}{}
		UserInfo["token"] = tokenString
		UserInfo["email"] = user.Email
		UserInfo["name"] = user.Name

		var reponse = map[string]interface{}{"UserInfo": UserInfo}
		reponse["text"] = "Udało się zalogować !"
		reponse["status"] = 1

		json.NewEncoder(w).Encode(reponse)

	}

}

func Login(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			json.NewEncoder(w).Encode("U used wrong method")
			return
		}

		user := &models.User{}
		json.NewDecoder(req.Body).Decode(&user)
		var UserData = &models.User{}
		Info := col.FindOne(ctx, bson.M{"email": user.Email}).Decode(&UserData)
		fmt.Println(Info)

		if Info != nil {
			fmt.Println("nie ma takiego typa", Info)
			var ErrorInfo = map[string]interface{}{}
			ErrorInfo["status"] = 0
			ErrorInfo["text"] = "Nie ma takiego chłopa w bazie danych"
			json.NewEncoder(w).Encode(ErrorInfo)
			return
		}

		errorPass := bcrypt.CompareHashAndPassword([]byte(UserData.Password), []byte(user.Password))

		if errorPass != nil && errorPass == bcrypt.ErrMismatchedHashAndPassword {
			var ErrorInfo = map[string]interface{}{}
			ErrorInfo["status"] = 0
			ErrorInfo["text"] = "Niepoprawne Hasło"
			json.NewEncoder(w).Encode(ErrorInfo)
			return
		}

		tokenString, _ := auth.CreateAccessToken(user)

		var UserInfo = map[string]interface{}{}
		UserInfo["token"] = tokenString
		UserInfo["email"] = UserData.Email
		UserInfo["name"] = UserData.Name
		var reponse = map[string]interface{}{"UserInfo": UserInfo}
		reponse["text"] = "Udało się zalogować !"
		reponse["status"] = 1
		json.NewEncoder(w).Encode(reponse)

	}
}

func FetchUserById(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			data := req.Context().Value("user")
			json.NewEncoder(w).Encode(data)
		}
	}
}
