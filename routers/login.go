package routers

import (
	"encoding/json"
	"net/http"

	"github.com/lucianoOlivera/twittor-GO-API-REST/bd"
	"github.com/lucianoOlivera/twittor-GO-API-REST/jwt"
	"github.com/lucianoOlivera/twittor-GO-API-REST/models"
)

//Login realiza el login
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var t models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "usuario y/o contraseña invaido"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
	}
	docuemnto, existe := bd.Intentologin(t.Email, t.Password)
	if existe == false {
		http.Error(w, "Usuario y/o contraseña invalidos", 400)
	}
	jwtKey, err := jwt.GenerateJWT(docuemnto)
	if err != nil {
		http.Error(w, "Ocurrio erro intetnado generar el token"+err.Error(), 400)
	}
	resp := models.RespuestaLogin{
		Token: jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
	//como guardar el token en una cookie
	/*expirationTime := time.Now().Add(24*hour)
	http.SetCookie(w,&http.Cookie{
		name:"token",
		value:jwtKey,
		Expires: expirationTime,
	})*/
}
