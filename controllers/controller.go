package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MarcosRoch4/cheap/database"
	"github.com/MarcosRoch4/cheap/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func AllPosto(w http.ResponseWriter, r *http.Request) {
	var p []models.Posto
	fmt.Println("Procurando info na tabela Posto")
	database.DB.Find(&p)
	fmt.Println("Encontrou isso na tabela: ", &p)

	json.NewEncoder(w).Encode(p)
	fmt.Println("Transpos os dados para API")
}

func RetornaUmPosto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var posto models.Posto
	database.DB.First(&posto, id)
	json.NewEncoder(w).Encode(posto)
}
