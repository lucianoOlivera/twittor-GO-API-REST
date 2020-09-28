package middlew

import (
	"net/http"

	"github.com/lucianoOlivera/twittor-GO-API-REST/routers"
)

//ValidoJWT valido el jwt que viene de la peticion
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorizacion"))
		if err != nil {
			http.Error(w, "Error en el token"+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
