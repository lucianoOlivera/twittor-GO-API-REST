package middlew

import (
	"net/http"

	"github.com/lucianoOlivera/twittor-GO-API-REST/bd"
)

//CheckBD un middlew que me permite conocer el estado de la base de datos
func CheckBD(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "conexion perdida con la Base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
