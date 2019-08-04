package main
import "fmt"

type Books struct {
	title string
	author string
	subtitle string
	price float64
}

func main() {
	var Book1 Books
	Book1.title = "Go Programming"
	Book1.author = "kenshero"
	Book1.subtitle = "xxxxxxx"
	Book1.price = 199.24
	fmt.Println(Book1.price)

	Golang := Books {
		title: "Go Go Go",
		author: "Kenshero",
		price: 123.2,
	}

	fmt.Println(Golang.author)
}
