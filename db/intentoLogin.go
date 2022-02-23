package db

import (
	"github.com/durbonca/globant_twitter_back/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) (models.Usuario, bool) {
	usuario, encontrado, _ := ExisteUsuario(email)
	if encontrado == false {
		return usuario, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usuario.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return usuario, false
	}

	return usuario, true
}
