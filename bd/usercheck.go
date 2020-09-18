package bd

import (
	"context"
	"time"

	"github.com/lucianoOlivera/twittor-GO-API-REST/models"
	"go.mongodb.org/mongo-driver/bson"
)

//Usercheck c
func Usercheck(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("users")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.Id.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
