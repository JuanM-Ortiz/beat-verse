package db

import (
	"context"
	"time"

	"github.com/JuanM-Ortiz/beat-verse/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckUser(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := Mongo.Database("beat-verse")
	collection := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.User

	err := collection.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
