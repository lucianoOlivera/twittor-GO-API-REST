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
	router.HandleFunc("/login", middlew.CheckBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.CheckBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/moficarPerfil", middlew.CheckBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweet", middlew.CheckBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminoTweet", middlew.CheckBD(middlew.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")

	router.HandleFunc("/ ", middlew.CheckBD(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.CheckBD(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/subirBanner", middlew.CheckBD(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.CheckBD(routers.ObtenerBanner)).Methods("GET")

	router.HandleFunc("/altaRelacion", middlew.CheckBD(routers.AltaRelacion)).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.CheckBD(routers.BajaRelacion)).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlew.CheckBD(routers.ConsultoRelacion)).Methods("GET")

	router.HandleFunc("/listaUsuario", middlew.CheckBD(routers.ListaUsuario)).Methods("GET")

	os.Setenv("PORT", "8080")
	PORT := os.Getenv("PORT")
	if PORT != "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
