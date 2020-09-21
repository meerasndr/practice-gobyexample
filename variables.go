package main
import "fmt"

func main(){
	// Type inference
	var a = 10;
	fmt.Println(a)

	// Explicit type declaration
	var b, c int = 20, 30
	fmt.Println(b, c)

	// More type inference
	var d = true
	var e = "somestring"
	fmt.Println(d)
	fmt.Println(e)

	// Declared but not initialized. f is 0 now
	var f int
	fmt.Println(f)
	var h bool
	fmt.Println(h) //default false
	var z string
	fmt.Println(z) //default blank string
	var y float32
	fmt.Println(y) //default 0

	//Shorthand for declaring and initializing
	g := "manor"
	fmt.Println(g)

}