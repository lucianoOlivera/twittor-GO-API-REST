package routers

import (
	"encoding/json"
	"net/http"

	"github.com/lucianoOlivera/twittor-GO-API-REST/bd"
	"github.com/lucianoOlivera/twittor-GO-API-REST/models"
)

func ConsultoRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relacion
	t.UsuarioID = IdUsuario
	t.UsuarioRelacionID = ID
	var res models.RespuestaConsultaRelacion

	status, err := bd.ConsultoRelacion(t)

	if err != nil || status == false {
		res.Status = false
	} else {
		res.Status = true
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
