package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/JavierDominguezGomez/not/middleware"
	"github.com/JavierDominguezGomez/not/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Handlers set port, handler, listen and serve the HTTP server.*/
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middleware.CheckDB(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}