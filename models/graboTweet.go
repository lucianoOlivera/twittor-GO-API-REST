package models

import "time"

type GraboTweet struct {
	UserId  string    `bson:"userid" json:"userid,omitemty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitemty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitemty"`
}
