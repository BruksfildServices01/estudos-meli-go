package main

import "fmt"

type Saudador interface {
	Saudacao() string
}

type Pessoa struct {
	nome  string
	idade int
}

func (p Pessoa) Saudacao() string {
	return "Olá, " + p.nome
}

func cumprimentar(s Saudador) {
	fmt.Println(s.Saudacao())
}

func definiIdade(p *Pessoa) {
	p.idade = 24
}

func main() {
	p := Pessoa{nome: "Lucas", idade: 25}
	fmt.Println(p)

	definiIdade(&p)
	fmt.Println(p)

	pessoa := Pessoa{nome: "Luis"}
	cumprimentar(pessoa)
}
