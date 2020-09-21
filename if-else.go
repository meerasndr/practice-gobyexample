package main
import "fmt"
func main(){
	num1 := 24
	if num1 % 2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}
	
	if num2 := 9; num2 < 0 {
		fmt.Println("Number is negative")
	} else if num2 < 10 {
		fmt.Println("Number has single digit")
	}else{
		fmt.Println("Number has multiple digits")
	}
}
