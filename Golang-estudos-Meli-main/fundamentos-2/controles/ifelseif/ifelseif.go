package main

import "fmt"

func notaParaConceito(n float64) string {

	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else {
		return "E"
	}

}

func main() {

	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.7))
	fmt.Println(notaParaConceito(2.1))

}
