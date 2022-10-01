package CommentsControllers

import (
	"BackendGo/pkg/models"
	"BackendGo/pkg/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllCommentsOfGivenPost(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		var PostId string
		json.NewDecoder(req.Body).Decode(&PostId)
		fmt.Println(PostId)
		cursor, err := col.Find(ctx, bson.M{"postid": PostId})

		if err != nil {
			fmt.Println("find error")
			return
		}
		var CommentData []bson.M

		if err = cursor.All(ctx, &CommentData); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(CommentData)

	}
}

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

		//change it later
		var Response = map[string]interface{}{}
		Response["status"] = 1
		Response["message"] = "Udało się opublikować komentarz !"

		json.NewEncoder(w).Encode(Response)
	}
}

func LikeCommentOnPostByUser(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}
