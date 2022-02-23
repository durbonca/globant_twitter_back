package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/durbonca/globant_twitter_back/db"
	"github.com/durbonca/globant_twitter_back/jwt"
	"github.com/durbonca/globant_twitter_back/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o Contraseña invalida"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Se requiere email de usuario", 400)
		return
	}

	documento, existe := db.Login(t.Email, t.Password)

	if !existe {
		http.Error(w, "Usuario y/o Contraseña invalida", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(documento)

	if err != nil {
		http.Error(w, "Ocurrió un error"+err.Error(), 400)
		return
	}

	resp := models.JWTLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(time.Hour * 24)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
