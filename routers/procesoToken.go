package routers

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/lucianoOlivera/twittor-GO-API-REST/bd"
	"github.com/lucianoOlivera/twittor-GO-API-REST/models"
)

var Email string
var IdUsuairo string

//ProcesoToken valida el token
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("lucianoOlivera")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Formato de token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.Usercheck(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IdUsuairo = claims.ID.Hex()
		}
		return claims, encontrado, IdUsuairo, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), nil

}
