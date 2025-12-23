package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("String:", "banana" == "banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 4)
	fmt.Println("=>", 3 >= 4)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas: ", d1 == d2)
	fmt.Println("Data: ", d1.Equal(d2))

	type Pessoa struct {
		nome  string
		idada int
	}
	p1 := Pessoa{"lucas", 24}
	p2 := Pessoa{"kaue", 17}

	fmt.Println("pessoas:", p1 == p2)

}
