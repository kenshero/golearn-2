package main
import "fmt"

func main() {
	go f(0)
	f2()
	// var input string
	// fmt.Scanln(&input)
}

func f(n int) {
	for i := 0; i < 100; i ++ {
		fmt.Println("ZZZZ")
	}
}


func f2() {
	for i := 0; i < 100; i ++ {
		fmt.Println("SSSS")
	}
}
