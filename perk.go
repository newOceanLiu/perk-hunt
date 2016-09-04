package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func viewCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	fmt.Println("Key: " + key)
	perk := Perk{Description: "food is great!", Like: 3, Heart: 9}
	company := Company{
		Id:     1,
		Name:   key,
		Detail: "great!",
		Perks:  []Perk{perk},
	}
	fmt.Println(company)
	json.NewEncoder(w).Encode(company)
}

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/viewCompany/{key}", viewCompany)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	fmt.Println("Rest API - mux Routers")
	handleRequests()
}
