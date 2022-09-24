package CommentsControllers

import (
	"BackendGo/pkg/models"
	"BackendGo/pkg/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
)

func CommentOnPostByUser(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var Comment models.Comment
		json.NewDecoder(req.Body).Decode(&Comment)
		fmt.Println(Comment)
		UserData := req.Context().Value("user")
		var temp = map[string]interface{}{}
		temp = UserData.(jwt.MapClaims)
		var Name = temp["Name"]
		var UserId = temp["UserID"]

		Comment.UserId = UserId.(string)
		Comment.UserName = Name.(string)

		fmt.Println(Comment)
		result, insertErr := col.InsertOne(ctx, Comment)
		utils.RegisterInsertErrors(result, insertErr)

		json.NewEncoder(w).Encode("SIEMA")
	}
}
func LikeCommentOnPostByUser(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}
