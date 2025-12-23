package main

import (
	"fmt"
	"math"
)

func main() {
	// forma - 1
	const pi float64 = 3.1415
	var raio = 3.2

	area := pi * math.Pow(raio, 2)

	fmt.Println("area da circurferencia é", area)

	// forma - 2
	const (
		a = 5
		b = 10
		c = 0
	)
	var (
		d = 2
		e = 100
	)
	fmt.Println("numero", a, b, c, d, e)

	// forma - 3

	var verdade, falso bool = true, false
	fmt.Println(verdade, "e", falso)

	n1, bole, letra := 2, false, "opa"
	fmt.Println(n1, bole, letra)

}
