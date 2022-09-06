package auth

import (
	"BackendGo/pkg/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoginIfno struct {
	Text   string
	Status int
}

//modify this funciton to be able to hadnle login as well as register
func FindUserByEmail(col *mongo.Collection, user *models.User, ctx context.Context) (LoginIfno, error, *models.User) {
	var UserData = &models.User{}
	Error := col.FindOne(ctx, bson.M{"email": user.Email}).Decode(&UserData)
	resp := LoginIfno{}
	if Error == nil {
		println(UserData, Error, "udało się znaleźć innego uzytkoniwka o tym mailu")
		resp.Text = "Podany uzytkownik istnieje juz w bazie"
		resp.Status = 0
		return resp, nil, UserData
	}
	resp.Text = ""
	resp.Status = 0
	return resp, Error, nil
}
