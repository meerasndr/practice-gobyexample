package main

import (
	"fmt"
	"math"
)
const pi = 3.14

func main() {
	var r = 7.0
	const angle  = 90
	fmt.Println("My pi and inbuilt pi", pi, math.Pi)
	fmt.Println("Sin 90: ", math.Sin(angle))
	fmt.Println("Cos 90: ", math.Cos(angle))
	fmt.Println("Area of circle: ", pi * r * r)
}