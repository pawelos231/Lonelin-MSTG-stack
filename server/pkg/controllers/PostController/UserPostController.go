package PostControllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteAllPostsOfUser(collectionOfPosts *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		var UserEmail string
		json.NewDecoder(req.Body).Decode(&UserEmail)
		res, err := collectionOfPosts.DeleteMany(ctx, bson.M{"email": UserEmail})
		if err != nil {
			fmt.Println(err, "coś poszło nie tak")
			json.NewEncoder(w).Encode("coś poszło nie tak przy usuwaniu postów")
		} else {
			fmt.Println(res)
			json.NewEncoder(w).Encode("udało się usunąć wszystkie twoje posty!")
		}

	}
}

//Will require admin user status
func DeleteAllPosts(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}

func FetchUserSpecificPosts(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		var UserEmail string
		json.NewDecoder(req.Body).Decode(&UserEmail)
		cursor, err := col.Find(ctx, bson.M{"email": UserEmail})
		if err != nil {
			fmt.Println("coś się wywaliło")
			json.NewEncoder(w).Encode("coś poszło nie tak")
		} else {
			var PostInfo []bson.M
			if (err) = cursor.All(ctx, &PostInfo); err != nil {
				log.Fatal(err)
			}
			json.NewEncoder(w).Encode(PostInfo)
		}

	}
}

func likePostByUser(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}

func CommentOnPostByUser(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}

func LikeCommentOnPostByUser(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}
