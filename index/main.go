package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// TODO
// handle CRUD operations in orders
// respond JSON
func handleOrders(w http.ResponseWriter, r *http.Request) {

}

func handleDashboard(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/order/", handleOrders)
	r.HandleFunc("/dashboard/", handleDashboard)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
