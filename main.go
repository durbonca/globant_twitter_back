package main

import (
	"log"

	db "github.com/durbonca/globant_twitter_back/db"
	handlers "github.com/durbonca/globant_twitter_back/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("No se pudo conectar a la base de datos")
		return
	}
	handlers.Handlers()
}
