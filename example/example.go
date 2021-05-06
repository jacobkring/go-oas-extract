package example

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	//
	r.HandleFunc("/pet/{petId}", getPet).Methods("GET")

	http.Handle("/", r)
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getPet (w http.ResponseWriter, r *http.Request) {

}