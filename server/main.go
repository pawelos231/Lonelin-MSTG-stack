package main

import (
	"BackendGo/pkg/controllers"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostInformiation struct {
	Title     string `json:"title"`
	CreatedAt string `json:"createdAt"`
	Message   string `json:"message"`
	UserName  string `json:"userName"`
}

func GetDataFromDatabase(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
			fmt.Println(PostsData)
			fmt.Println("succes", reflect.TypeOf(cursor))
			json.NewEncoder(w).Encode(PostsData)
		}

	}
}

func PostDataToDataBase(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)

		var PostDetails PostInformiation
		json.NewDecoder(r.Body).Decode(&PostDetails)
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
	}
}

func main() {
	//conmect to database
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10000*time.Second)
	col := client.Database("MSTG_STACK").Collection("PostInfos")
	mime.AddExtensionType(".js", "application/javascript")
	r := mux.NewRouter()

	fmt.Println(controllers.SayHello())

	r.HandleFunc("/homeland", PostDataToDataBase(col, ctx)).Methods("POST")
	r.HandleFunc("/getdata", GetDataFromDatabase(col, ctx)).Methods("GET")
	http.Handle("/", http.FileServer(http.Dir("static")))
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS()(r)))
}
