package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lucianoOlivera/twittor-GO-API-REST/middlew"
	"github.com/lucianoOlivera/twittor-GO-API-REST/routers"

	"github.com/rs/cors"
)

//Manejadores sete mi puerto y pongo a escuchar al servidor
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.CheckBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT != "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
