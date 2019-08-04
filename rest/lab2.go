package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type addressBook struct {
	Firstname string
	Lastname  string
	Code      int
	Phone     string
}

func getAddressBookAll(w http.ResponseWriter, r *http.Request) {
	addBook := addressBook{
		Firstname: "Teeruk",
		Lastname:  "Imburanaprawat",
		Code:      1993,
		Phone:     "087879789",
	}
	json.NewEncoder(w).Encode(addBook)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/getAddress", getAddressBookAll)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
