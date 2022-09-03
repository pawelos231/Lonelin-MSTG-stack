package PostControllers

import (
	"BackendGo/pkg/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDataFromDatabase(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
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
			w.WriteHeader(http.StatusOK)
			/*
				req.ParseMultipartForm(10 << 20)
				file, handler, err := req.FormFile("PostFile")
				if err != nil {
					fmt.Println("Error retrieving file")
					fmt.Println(err)
					return
				}
				defer file.Close()
				fmt.Printf("Upload fILE: %+v\n", handler.Filename)
				fmt.Printf("Upload fILE: %+v\n", handler.Size)
				fmt.Printf("Upload fILE: %+v\n", handler.Header)
			*/
			var PostDetails models.PostInformiation
			json.NewDecoder(req.Body).Decode(&PostDetails)

			//middleware to check the user

			UserData := req.Context().Value("user")
			var temp = map[string]interface{}{}
			temp = UserData.(jwt.MapClaims)
			var Name = temp["Name"]
			var Email = temp["Email"]
			PostDetails.UserName = Name.(string)
			PostDetails.Email = Email.(string)

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
		w.WriteHeader(http.StatusOK)
		/*TODO: write some update post functionality when frontend is ready*/
	}
}
func DeletePostFromDatabase(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)

		var DataFromUser string
		json.NewDecoder(req.Body).Decode(&DataFromUser)
		fmt.Println(DataFromUser)

		objID, _ := primitive.ObjectIDFromHex(DataFromUser)
		res, err := col.DeleteOne(ctx, bson.M{"_id": objID})

		fmt.Println(res)

		var responseMessage = map[string]interface{}{}

		if err == nil {
			responseMessage["status"] = 1
			responseMessage["text"] = "udało się usunąć post :o"
			json.NewEncoder(w).Encode(responseMessage)
		} else {
			responseMessage["status"] = 0
			responseMessage["text"] = "Coś poszło masno nie tak"
			json.NewEncoder(w).Encode(responseMessage)
		}
		/*TODO: write some update post functionality when frontend is ready*/
		//res , err = col.DeleteOne(ctx, bson.M{"_id": })

	}
}
