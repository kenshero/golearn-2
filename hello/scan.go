package main
import "fmt"

func main() {
	fmt.Println("Input Your Number:")
	var input float64
	fmt.Scanf("%f", &input)
	output:=input*2
	fmt.Println(output)
}