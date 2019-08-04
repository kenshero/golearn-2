package main
import "fmt"

func main() {
	var myAge int = 200
	myAge2 := 500
	age1,age2:=35,80
	fmt.Println("Value Variable = ", myAge)
	fmt.Println("Value Variable = ", myAge2)
	// fmt.Println(text)
	fmt.Println(age1,age2)

	p1:="kenshero"
	p2:="555"

	p3:=p1+p2
	fmt.Println(p3)
	fmt.Println(p3[1:3])
	fmt.Println(p3[4:])
}


