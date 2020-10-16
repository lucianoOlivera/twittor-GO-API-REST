package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/lucianoOlivera/twittor-GO-API-REST/bd"
)

func ListaUsuario(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagtemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Enviar parameto  pagina entero mayor a 0", http.StatusBadRequest)
		return
	}
	pag := int64(pagtemp)

	resultado, status := bd.LeoUsuarioTodos(IdUsuario, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)

	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resultado)
}
