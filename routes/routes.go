package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guilherme2912/go-rest-api/controllers"
)

func HandleResquet() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))

}
