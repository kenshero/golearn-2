package main
import "fmt"

func main() {
	for i:=1; i<=10; i++ {
		fmt.Println("kenshero ")
		if i % 2 == 0 {
			fmt.Println("kuu ")
		} else if i % 2 == 1 {
			fmt.Println("xxx ")
		}
	}
}

