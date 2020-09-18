package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN va ser igual la conecion de la bd
var MongoCN = ConnectionBD()

var clientOptions = options.Client().ApplyURI("mongodb+srv://twittor:38209056lo@cluster0.8qtlb.gcp.mongodb.net/?retryWrites=true&w=majority")

//ConnectionBD la funcion que permite conectar con la funcion
func ConnectionBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Printf("conecion exitosa con la db")
	return client
}

//CheckConnection e sun ping a la bd
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
