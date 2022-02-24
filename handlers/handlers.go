package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/durbonca/globant_twitter_back/middlewares"
	"github.com/durbonca/globant_twitter_back/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Set port and listen server */
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlewares.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/perfil", middlewares.CheckDB(middlewares.ValidateJWT(routers.Profile))).Methods("GET")
	router.HandleFunc("/modificar_perfil", middlewares.CheckDB(middlewares.ValidateJWT(routers.UpdateProfile))).Methods("PUT")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
