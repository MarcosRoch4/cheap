package routes

import (
	"log"
	"net/http"

	"github.com/MarcosRoch4/cheap/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/posto", controllers.AllPosto).Methods("Get")
	r.HandleFunc("/api/posto/{id}", controllers.RetornaUmPosto).Methods("Get")

	log.Fatal(http.ListenAndServe(":8000", r))
}
