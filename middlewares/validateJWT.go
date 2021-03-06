package middlewares

import (
	"net/http"

	"github.com/durbonca/globant_twitter_back/routers"
)

func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.GetTokenFromRequest(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error de Token "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
