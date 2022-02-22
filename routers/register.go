package routers

import (
	"encoding/json"
	"net/http"

	"github.com/durbonca/globant_twitter_back/db"
	"github.com/durbonca/globant_twitter_back/models"
)

/* Register users on DB */
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Se requiere email de usuario", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Se requiere password de al menos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := db.ExisteUsuario(t.Email)
	if encontrado {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := db.InsertUser(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado registrar el usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
