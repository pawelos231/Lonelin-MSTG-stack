package PostControllers

import (
	"BackendGo/pkg/models"
	"BackendGo/pkg/utils"
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

func GetPostsData(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		if req.Method != http.MethodGet {
			println("You used a wrong method")
			return
		}
		
		w.WriteHeader(http.StatusOK)

		cursor, err := col.Find(ctx, bson.M{})
		if err != nil {
			fmt.Println("Find errror", err)
			return
		}

		var PostsData []bson.M
		if err = cursor.All(ctx, &PostsData); err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(PostsData)

	}

}

func PostDataToDataBase(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			println("You used a wrong method")
			return
		}
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
		utils.RegisterInsertErrors(result, insertErr)

	}
}

func GetPostByUniqueId(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			println("You used a wrong method")
			return
		}

		w.WriteHeader(http.StatusOK)
		queryValue := req.FormValue("q")
		println(queryValue)
		objID, _ := primitive.ObjectIDFromHex(queryValue)
		_, err := col.Find(ctx, bson.M{})
		if err != nil {
			fmt.Println("Find errror", err)
			return
		}

		var SinglePost bson.M
		value := col.FindOne(ctx, bson.M{"_id": objID}).Decode(&SinglePost)
		if err != nil {
			log.Fatal("Something went wrong", err, " ", value)
		}
		json.NewEncoder(w).Encode(SinglePost)

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

		if err != nil {
			responseMessage["status"] = 0
			responseMessage["text"] = "Coś poszło masno nie tak"
			json.NewEncoder(w).Encode(responseMessage)
		}
		responseMessage["status"] = 1
		responseMessage["text"] = "udało się usunąć post :o"
		json.NewEncoder(w).Encode(responseMessage)
	}
}
