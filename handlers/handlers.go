package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/ccmorenov/microservicesounds/middlew"
	"github.com/ccmorenov/microservicesounds/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Manejadores*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ViewBD(routes.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ViewBD((routes.Login))).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ViewBD(routes.VerPerfil)).Methods("GET")
	router.HandleFunc("/modifyperfil", middlew.ViewBD(routes.ModifyPerfil)).Methods("PUT")
	router.HandleFunc("/eliminaruser", middlew.ViewBD(routes.EliminarUser)).Methods("DELETE")
	router.HandleFunc("/uploadavatar", middlew.ViewBD(routes.UploadAvatar)).Methods("POST")
	router.HandleFunc("/getavatar", middlew.ViewBD(routes.GetAvatar)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
