package main
import "fmt"

func main() {
	// x:=make([]float64, 5)
	// fmt.Println(x)

	x := []int{5,4,3}
	x2 := make([]int, 2)
	copy(x2, x)

	// arr := [5]float64{1,2,3,4,5}

	s := []int{1,2,4,5}
	s2 := append(s, 9, 10)

	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(x2)

}
