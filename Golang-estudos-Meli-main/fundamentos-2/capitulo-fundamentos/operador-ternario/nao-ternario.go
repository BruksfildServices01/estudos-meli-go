package main

import "fmt"

// operador ternarios nao existe no go
func obterresultado(nota float64) string {
	if nota >= 6 {
		return "aprovador"
	}
	return "reprovado"
}

func main() {
	fmt.Println(obterresultado(6.2))
}
