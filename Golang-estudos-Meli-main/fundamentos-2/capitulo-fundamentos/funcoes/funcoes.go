package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func imprimir(valor int) {
	fmt.Println(valor)
}

func main() {
	resultado := somar(3, 4)
	imprimir(resultado)
}
