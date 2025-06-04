package routes

import (
	"api/rest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods("DELETE")
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", r))
}
