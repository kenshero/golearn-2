package main
import "fmt"

func main() {
	x := 10
	fmt.Printf("value is %d\n", x)
	fmt.Printf("address x is %x\n", &x)

	var p *int
	p = &x
	fmt.Printf("Pointer P = %x \n", p)
	*p = 20

	fmt.Printf("value X is %d\n", x)
	fmt.Printf("address x is %x\n", &x)

	x = 555

	fmt.Printf("value X is %d\n", x)
	fmt.Printf("address x is %x\n", &x)
}
