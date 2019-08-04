package main
import "fmt"

func main() {
	// dosomething("Kenshero")
	// dosomething("King Of Everything")
	// addition(9, 2)
	// var result int = addition2(5, 3)
	// fmt.Println(result)
	// summation(1, 2, 5)
	fmt.Println(factorial(5))
}

func dosomething(str string) {
	fmt.Println(str)
}

func addition(a int, b int) {
	fmt.Println(a+b)
}

func addition2(a int, b int) int {
	output:=a+b
	return output
}

func summation(args...int) {
	var total int
	for _ , n:=range args{
		total+=n
	}
	fmt.Println(total)
}

func factorial(a int) int {
	if a == 0 {
		return 1
	}
	return a*factorial(a - 1)
}