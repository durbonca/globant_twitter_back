package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/durbonca/globant_twitter_back/db"
	"github.com/durbonca/globant_twitter_back/models"
)

var Email, IDUser string

func GetTokenFromRequest(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("elGueboAnalfabeta_Criptico_69")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Invalid token")
	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := db.ExisteUsuario(claims.Email)

		if encontrado {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, encontrado, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("Invalid token")
	}
	return claims, false, string(""), err
}
