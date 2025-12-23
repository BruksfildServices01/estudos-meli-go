package main

import "fmt"

func main() {
	fmt.Print("mesma")
	fmt.Print("linha")

	fmt.Println("nova")
	fmt.Println("linha")

	x := 3.141516

	xs := fmt.Sprint(x)
	fmt.Println("o valor de x é " + xs)
	fmt.Println("o valor de x é", x)

	fmt.Printf("O valor de X é %.2f.\n", x)

	a := 1      // int
	b := 1.9999 // float64
	c := false  // bool
	d := "opa"  // string

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

}
