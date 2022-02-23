package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/durbonca/globant_twitter_back/models"
)

func GenerateJWT(t models.Usuario) (string, error) {
	miClave := []byte("elGueboAnalfabeta_Criptico_69")
	payload := jwt.MapClaims{
		"email":     t.Email,
		"nombre":    t.Nombre,
		"apellidos": t.Apellidos,
		"fechaNac":  t.FechaNac,
		"bio":       t.Bio,
		"ubicacion": t.Ubicacion,
		"web":       t.Web,
		"_id":       t.ID,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(miClave)

	if err != nil {
		return tokenString, err
	}

	return tokenString, nil

}
