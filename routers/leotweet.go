package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/lucianoOlivera/twittor-GO-API-REST/bd"
)

//LeoTweets de routers
func LeoTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "debe inviar un parametro id", http.StatusBadRequest)
	}
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "debe inviar un parametro de pagina", http.StatusBadRequest)
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "debe inviar un parametro  mayor a 0", http.StatusBadRequest)
	}
	pag := int64(pagina)

	respuesta, correcto := bd.LeoTweets(ID, pag)
	if correcto == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
