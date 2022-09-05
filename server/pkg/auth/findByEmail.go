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

func FindUserByEmail(col *mongo.Collection, user *models.User, ctx context.Context) (LoginIfno, error) {
	var User bson.M
	Error := col.FindOne(ctx, bson.M{"email": user.Email}).Decode(&User)
	resp := LoginIfno{}
	if Error == nil {
		println(User, Error, "udało się znaleźć innego uzytkoniwka o tym mailu")
		resp.Text = "Podany uzytkownik istnieje juz w bazie"
		resp.Status = 0
		return resp, nil
	}
	resp.Text = ""
	resp.Status = 0
	return resp, Error
}
