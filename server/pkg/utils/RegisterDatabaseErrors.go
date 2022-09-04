package utils

import (
	"fmt"
	"reflect"

	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterInsertErrors(result *mongo.InsertOneResult, insertErr error) {
	if insertErr != nil {
		fmt.Println("IntertOne Error", insertErr)
		return
	}
	fmt.Println("insertOne() result type", reflect.TypeOf(result))
	fmt.Println("insertOne() api result type", result)
	newID := result.InsertedID
	fmt.Println("InsertedOne(), newID ", newID)
	fmt.Println("insertedOne(), newID type: ", reflect.TypeOf(newID))

}
