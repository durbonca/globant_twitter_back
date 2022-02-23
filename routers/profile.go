package routers

import (
	"encoding/json"
	"net/http"

	"github.com/durbonca/globant_twitter_back/db"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	perfil, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "Error searching profile: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)

}
