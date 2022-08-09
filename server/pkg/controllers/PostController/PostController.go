package PostControllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostInformiation struct {
	Title     string `json:"title"`
	CreatedAt string `json:"createdAt"`
	Image     string `json:"image"`
	Message   string `json:"message"`
	UserName  string `json:"userName"`
}

func GetDataFromDatabase(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*") // for CORS
			w.WriteHeader(http.StatusOK)
			cursor, err := col.Find(ctx, bson.M{})
			if err != nil {
				fmt.Println("Find errror", err)
			} else {
				var PostsData []bson.M
				if err = cursor.All(ctx, &PostsData); err != nil {
					log.Fatal(err)
				}
				fmt.Println("succes", reflect.TypeOf(cursor))
				json.NewEncoder(w).Encode(PostsData)
			}

		} else {
			println("You used a wrong method")
		}

	}
}

func PostDataToDataBase(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(http.StatusOK)

			var PostDetails PostInformiation
			json.NewDecoder(req.Body).Decode(&PostDetails)
			fmt.Println("Collection Type: ", reflect.TypeOf(col), "/n")
			result, insertErr := col.InsertOne(ctx, PostDetails)
			if insertErr != nil {
				fmt.Println("IntertOne Error", insertErr)
			} else {
				fmt.Println("insertOne() result type", reflect.TypeOf(result))
				fmt.Println("insertOne() api result type", result)
				newID := result.InsertedID
				fmt.Println("InsertedOne(), newID ", newID)
				fmt.Println("insertedOne(), newID type: ", reflect.TypeOf(newID))
			}

		} else {
			println("You used a wrong method")
		}
	}
}

func GetPostByUniqueId(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(http.StatusOK)

			queryValue := req.FormValue("q")
			println(queryValue)
			objID, _ := primitive.ObjectIDFromHex(queryValue)
			_, err := col.Find(ctx, bson.M{})
			if err != nil {
				fmt.Println("Find errror", err)
			} else {
				var SinglePost bson.M
				value := col.FindOne(ctx, bson.M{"_id": objID}).Decode(&SinglePost)
				if err != nil {
					log.Fatal("Something went wrong", err, " ", value)
				}
				json.NewEncoder(w).Encode(SinglePost)
			}
		} else {
			println("You used a wrong method")
		}
	}
}

func UpdatePostFromDatabase(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		/*TODO: write some update post functionality when frontend is ready*/
	}
}
func DeletePostFromDatabase(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodDelete {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(http.StatusOK)
			/*TODO: write some update post functionality when frontend is ready*/
		}
	}
}
