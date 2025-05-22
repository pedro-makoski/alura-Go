package controllers

import (
	"api/rest/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(models.Personalidades)
	if err != nil {
		panic(err.Error())
	}
}
