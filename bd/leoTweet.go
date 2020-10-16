package bd

import (
	"context"
	"log"

	"time"

	"github.com/lucianoOlivera/twittor-GO-API-REST/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//LeoTweets leo los tweets de la base de datos mongo
func LeoTweets(ID string, pagina int64) ([]*models.DuevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	var resultados []*models.DuevuelvoTweets

	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	opciones.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}
	for cursor.Next(context.TODO()) {
		var registro models.DuevuelvoTweets
		if err := cursor.Decode(&registro); err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}
	return resultados, true
}
