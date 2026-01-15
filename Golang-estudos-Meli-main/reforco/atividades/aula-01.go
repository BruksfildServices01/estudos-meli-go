package main

import "fmt"

var nome_test = "rafa"
var idade_test = 30
var ativado_test = false
var preco_test float64 = 30.30
var inteiro_test int = 20

const ApiVersion = "v1"

type User struct {
	ID     int
	Name   string
	Email  string
	Active bool
}

func main() {
	fmt.Println(nome_test, idade_test, ativado_test, preco_test, inteiro_test)
	idade := 24
	preco := 20.20
	inteiro := 10
	fmt.Println(idade, preco, inteiro)

	user := Create_User()
	if user.Active == true {
		fmt.Println(user.Name)
	}
}

func Create_User() User {

	id := 1
	email := "lucas@gmail.com"
	ativado := false
	nome := "lucas"

	user := User{
		ID:     id,
		Name:   nome,
		Email:  email,
		Active: ativado,
	}

	return user
}
