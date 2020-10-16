package routers

import (
	"net/http"

	"github.com/lucianoOlivera/twittor-GO-API-REST/bd"
)

func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) > 1 {
		http.Error(w, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}
	err := bd.BorroTweet(ID, IdUsuario)
	if err != nil {
		http.Error(w, "ocurrio un erro al intentar borrar el tweet"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
