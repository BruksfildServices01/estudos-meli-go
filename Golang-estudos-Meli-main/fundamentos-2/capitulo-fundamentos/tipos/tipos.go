package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	// inteiro
	fmt.Println(1, 100, 1000)
	fmt.Println("literal inteiro é", reflect.TypeOf(32000))
	// int sem sinal (so positivos)
	var b byte = 255
	fmt.Println("o byte é", reflect.TypeOf(b))
	// com sinal int8... int64
	i1 := math.MaxInt64
	fmt.Println("o valor maximo do int é", i1)
	// numeros reais ( float32, float64)
	var x = 49.99
	fmt.Println("o tipo de x é", reflect.TypeOf(x))
	fmt.Println("o tipo do literal 49.99 é", reflect.TypeOf(49.99))

	// boolean

	bo := true
	fmt.Println("o tipo da var bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "ola meu nome é lucas"
	fmt.Println(s1 + "!")

	//string com varias linhas
	s2 := `ola
	       me
		   chamo
		   lucas`
	fmt.Println("o tamanho da string é", len(s2))

}
