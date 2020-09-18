package main

import (
	"log"

	"github.com/lucianoOlivera/twittor-GO-API-REST/bd"
	"github.com/lucianoOlivera/twittor-GO-API-REST/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("sin conexion a la base de datos ")
		return
	}
	handlers.Manejadores()
}
