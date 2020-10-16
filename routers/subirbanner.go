package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/lucianoOlivera/twittor-GO-API-REST/bd"
	"github.com/lucianoOlivera/twittor-GO-API-REST/models"
)

func SubirBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/banners/" + IdUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "erro al subir la imagen!"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "erro al copiar la imagen!"+err.Error(), http.StatusBadRequest)
		return
	}
	var usuario models.Usuario
	var status bool

	usuario.Banner = IdUsuario + "." + extension
	status, err = bd.ModificoRegistro(usuario, IdUsuario)
	if err != nil || status == false {
		http.Error(w, "erro al grabar el banners!"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
