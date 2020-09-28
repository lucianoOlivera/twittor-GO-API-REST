package routers

import (
	"encoding/json"
	"net/http"

	"github.com/lucianoOlivera/twittor-GO-API-REST/bd"
)

//VerPerfil a
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro de la id", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Error al buscar el registro "+err.Error(), http.StatusNotFound)
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
