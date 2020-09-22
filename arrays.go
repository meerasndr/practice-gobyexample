package main
import "fmt"
func main(){
	var a [5] int
	fmt.Println("emp: ", a)

	a[4] = 99
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	fmt.Println("Length: ", len(a))

	// declare and initialize array in one line:
	b := [5]int {1, 2, 3, 4, 5}
	fmt.Println("declare and initialize: ", b)

	var c[] int
	fmt.Println("Array without size: ", c)
	c[0] = 0
	c[1] = 1
	fmt.Println(c)

}