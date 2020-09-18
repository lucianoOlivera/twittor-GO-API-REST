package routers

import (
	"encoding/json"
	"net/http"

	"github.com/lucianoOlivera/twittor-GO-API-REST/bd"
	"github.com/lucianoOlivera/twittor-GO-API-REST/models"
)

//Registro es la funcion para crear el registro usuario
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
	}
	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar uan contraseña de almenos de 6 caracteres", 400)
		return
	}
	if _, encontrado, _ := bd.Usercheck(t.Email); encontrado == true {
		http.Error(w, "Ya existe un usario con ese email", 400)
		return
	}
	_, status, err := bd.InsertRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un erro al intentar realizar el registro de usuario"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Nose a logrado insertar en registro del usuario", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
