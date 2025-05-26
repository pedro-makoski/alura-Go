package controllers

import (
	"api/rest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personalidade := range models.Personalidades {
		if strconv.Itoa(personalidade.Id) == id {
			json.NewEncoder(w).Encode(personalidade)
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Personalidade não encontrada"})
	w.WriteHeader(http.StatusNotFound)
}
