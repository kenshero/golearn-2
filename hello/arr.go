package main
import "fmt"

func main() {
	var x[5] int
	x[0]=10
	x[1]=50
	x[2]=25
	x[3]=30
	x[4]=100
	fmt.Println(x[3])
	x:=[10]float64{10,10,10,10,10}
	var total float64
	fmt.Println(x)
	for _, value:= range x {
		total += value
	}
	fmt.Println(total)
}
