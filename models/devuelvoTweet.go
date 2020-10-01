package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//DuevuelvoTweets devuelve una estuctura de devuevloTweets
type DuevuelvoTweets struct {
	ID     primitive.ObjectID `bson:"userid" json:"userid,omitemty"`
	UserId string             `bson:"id" json:"id,omitemty"`
	Mesaje string             `bson:"mensaje" json:"mensaje,omitemty"`
	fecha  time.Time          `bson:"time" json:"time,omitemty"`
}
