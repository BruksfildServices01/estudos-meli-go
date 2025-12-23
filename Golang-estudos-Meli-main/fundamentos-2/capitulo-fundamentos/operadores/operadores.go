package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Somar =>", a+b)
	fmt.Println("subtracao =>", a-b)
	fmt.Println("divisao =>", a/b)
	fmt.Println("multiplicacao =>", a*b)

	// bitwise
	fmt.Println("E =>", a&b)
	fmt.Println("OU =>", a|b)
	fmt.Println("Xor =>", a^b)

	c := 3.0
	d := 2.0

	fmt.Println("maior =>", math.Max(float64(c), float64(d)))
	fmt.Println("menor =>", math.Min(c, d))
	fmt.Println("exponenciacao =>", math.Pow(c, d))

}
