package UserController

import (
	"BackendGo/pkg/auth"
	"BackendGo/pkg/models"
	"BackendGo/pkg/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func RefreshTokenHandler(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		user := &models.User{}
		json.NewDecoder(req.Body).Decode(&user)

		if !utils.ValidateToken(w, req) {
			return
		}

		_, _, User := auth.FindUserByEmail(col, user, ctx)

		RefreshTokenString, _ := auth.CreateRefreshToken(User)
		auth.SendRefreshToken(w, req, RefreshTokenString)
		tokenString, _ := auth.CreateAccessToken(User)

		var UserInfo = map[string]interface{}{}
		UserInfo["token"] = tokenString
		UserInfo["email"] = User.Email
		UserInfo["name"] = User.Name
		var reponse = map[string]interface{}{"UserInfo": UserInfo}
		reponse["text"] = "Udało się zalogować !"
		reponse["status"] = 1
		fmt.Println(reponse)
		json.NewEncoder(w).Encode(reponse)
	}
}
