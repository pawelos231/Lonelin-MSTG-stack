package UserController

import (
	"BackendGo/pkg/auth"
	consts "BackendGo/pkg/constants"
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

		//decode data from client, hash password and attach the necesary informations about user to a query to database
		user := &models.User{}
		json.NewDecoder(req.Body).Decode(&user)
		password, _ := utils.HashValue(bcrypt.MinCost, user)
		user.UserId = uuid.New().String()
		user.Role = consts.USER

		emailFound, errEmail, _ := auth.FindUserByEmail(col, user, ctx)
		if errEmail == nil {
			json.NewEncoder(w).Encode(emailFound)
			return
		}

		//insert user to database
		result, insertErr := col.InsertOne(ctx, user)
		utils.RegisterInsertErrors(result, insertErr)

		//create refresh token and send it via cookie
		RefreshTokenString, _ := auth.CreateRefreshToken(user)
		auth.SendRefreshToken(w, req, RefreshTokenString)
		user.Password = string(password)

		//create access token and pass it to the user later on
		tokenString, _ := auth.CreateAccessToken(user)

		var UserInfo = map[string]interface{}{}
		UserInfo["token"] = tokenString
		UserInfo["email"] = user.Email
		UserInfo["name"] = user.Name

		var reponse = map[string]interface{}{"UserInfo": UserInfo}
		reponse["Text"] = "Udało się zalogować !"
		reponse["Status"] = 1

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

		if Info != nil {
			fmt.Println("nie ma takiego typa", Info)
			var ErrorInfo = map[string]interface{}{}
			ErrorInfo["status"] = 0
			ErrorInfo["text"] = "Nie ma takiego chłopa w bazie danych"
			json.NewEncoder(w).Encode(ErrorInfo)
			return
		}

		ErrorMessage := utils.CompareHashAndPassword(UserData.Password, user.Password, w)
		if ErrorMessage != nil {
			return
		}
		//create refresh token and send it via cookie
		RefreshTokenString, _ := auth.CreateRefreshToken(user)
		auth.SendRefreshToken(w, req, RefreshTokenString)

		//send access token to the client
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
