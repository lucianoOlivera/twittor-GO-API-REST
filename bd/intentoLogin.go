package bd

import (
	"github.com/lucianoOlivera/twittor-GO-API-REST/models"
	"golang.org/x/crypto/bcrypt"
)

//Intentologin s
func Intentologin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := Usercheck(email)
	if encontrado == false {
		return usu, false
	}
	passwordbytes := []byte(password)
	passwordBD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordbytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
