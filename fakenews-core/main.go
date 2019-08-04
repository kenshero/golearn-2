package main

import (
	"fakenews-core/news/controllers"
	"fakenews-core/news/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	book := controllers.GetBook()
	routes := routes.GetRoutes()
	fmt.Println("book is %s", book)
	// fmt.Println("routes is %s", routes)
	// fmt.Println("types is %s\n", types.Book)
	log.Print("The service is ready to listen and serve. 555555555555")
	http.ListenAndServe(":8080", routes)
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
