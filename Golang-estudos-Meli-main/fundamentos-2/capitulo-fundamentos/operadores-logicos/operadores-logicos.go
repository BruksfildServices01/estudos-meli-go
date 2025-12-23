package main

import "fmt"

func compras(trampo1, trampo2 bool) (bool, bool, bool) {
	compraTv50 := trampo1 && trampo2
	compraTv32 := trampo1 != trampo2
	comprasorverte := trampo1 || trampo2

	return compraTv50, compraTv32, comprasorverte
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("tv50: %t, tv32: %t, sorvete: %t, saudavel: %t\n", tv50, tv32, sorvete, !sorvete)
}
