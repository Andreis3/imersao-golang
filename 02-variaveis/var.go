package main

import "fmt"

type Pessoa struct {
	Nome     string
	Idade    uint8
	Cpf      string
	Endereco string
	Cep      string
}

func main() {
	p := Pessoa{
		Nome:     "João",
		Idade:    20,
		Cpf:      "123.456.789-00",
		Endereco: "Rua 1",
	}

	fmt.Println(p)

	p.Cep = "12345-678"

	fmt.Println(p)
}
