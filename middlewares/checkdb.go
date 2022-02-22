package middlewares

import (
	"net/http"

	"github.com/durbonca/globant_twitter_back/db"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "No se pudo conectar a la base de datos", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}
