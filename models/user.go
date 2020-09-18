package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Usuario estrucutra
type Usuario struct {
	Id              primitive.ObjectID `bson: "_id,omitempty" json:"id"`
	Nombre          string             `bson: "nombre,omitempty" json:"nombre"`
	Apellidos       string             `bson: "apellidos" json:"apellidos,omitempty"`
	FechaNacimiento time.Time          `bson: "fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email           string             `bson: "email" json:"email"`
	Password        string             `bson: "password" json:"password,omitempty"`
	Avatar          string             `bson: "avatar" json:"avatar,omitempty"`
	banner          string             `bson: "banner" json:"banner,omitempty"`
	Ubicacion       string             `bson: "ubicacion" json:"ubicacion,omitempty"`
	SitioWeb        string             `bson: "sitioweb" json:"sitioweb,omitempty"`
}