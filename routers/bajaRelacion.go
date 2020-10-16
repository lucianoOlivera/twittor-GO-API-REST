package routers

import (
	"net/http"

	"github.com/lucianoOlivera/twittor-GO-API-REST/bd"
	"github.com/lucianoOlivera/twittor-GO-API-REST/models"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relacion
	t.UsuarioID = IdUsuario
	t.UsuarioRelacionID = ID
	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al borrar relacion "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "no se a logrado borrar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
