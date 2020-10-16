package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/lucianoOlivera/twittor-GO-API-REST/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//LeoUsuarioTodos a
func LeoUsuarioTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("users")

	var resultado []*models.Usuario

	findOption := options.Find()
	findOption.SetSkip((page - 1) * 20)
	findOption.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}
	cur, err := col.Find(ctx, query, findOption)
	if err != nil {
		fmt.Println(err.Error())
		return resultado, false
	}
	var encontrado, incluir bool
	for cur.Next(ctx) {
		var s models.Usuario
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return resultado, false
		}
		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex()
		incluir = false

		encontrado, err = ConsultoRelacion(r)
		if tipo == "new" && encontrado == false {
			incluir = true
		}
		if tipo == "follow" && encontrado == true {
			incluir = true
		}
		if r.UsuarioRelacionID == ID {
			incluir = false
		}
		if incluir == true {
			s.Password = ""
			s.Biografia = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Email = ""
			s.Banner = ""

			resultado = append(resultado, &s)
		}

	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return resultado, false
	}
	cur.Close(ctx)
	return resultado, true
}
