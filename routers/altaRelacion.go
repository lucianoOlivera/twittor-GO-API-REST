package routers

import (
	"net/http"

	"github.com/lucianoOlivera/twittor-GO-API-REST/bd"
	"github.com/lucianoOlivera/twittor-GO-API-REST/models"
)

func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "debe enviar el id", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IdUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al inserta relacion "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "no se a logrado insertar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
