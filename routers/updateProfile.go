package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/durbonca/globant_twitter_back/db"
	"github.com/durbonca/globant_twitter_back/models"
)

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var status bool
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "error en los datos recibidos"+err.Error(), 400)
		return
	}
	fmt.Println("IDUSER", IDUser)
	status, err = db.UpdateProfile(t, IDUser)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar actualizar el perfil"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado actualizar el perfil", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
