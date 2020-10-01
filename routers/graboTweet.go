package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/lucianoOlivera/twittor-GO-API-REST/bd"
	"github.com/lucianoOlivera/twittor-GO-API-REST/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserId:  IdUsuairo,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.IncertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intertar el registro"+err.Error(), 400)
	}
	if status == false {
		http.Error(w, "No se a logrado insertar el tweet", 400)
	}
	w.WriteHeader(http.StatusCreated)
}
