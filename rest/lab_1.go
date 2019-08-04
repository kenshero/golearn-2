package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type test_struct struct {
	Test string
}

func parseGhPost(rw http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var t test_struct
	err := decoder.Decode(&t)

	if err != nil {
		fmt.Println("Error")
		panic(err)
	}

	fmt.Println(t.Test)
}

func main() {
	log.Print("Starting the service")

	http.HandleFunc("/home", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	http.HandleFunc("/testpost", parseGhPost)

	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
