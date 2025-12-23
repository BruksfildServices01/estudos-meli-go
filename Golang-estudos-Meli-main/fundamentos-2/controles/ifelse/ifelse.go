package main

import "fmt"

func imprimirResultado(nota float64) {

	if nota >= 7 {
		fmt.Println("aprovado com nota:", nota)
	} else {
		fmt.Println("reprovado com nota:", nota)
	}
}

func main() {
	imprimirResultado(6.6)
	imprimirResultado(7)
}
