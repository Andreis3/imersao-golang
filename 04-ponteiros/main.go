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
	p := &Pessoa{
		Nome: "João",
	}

	fmt.Println(p)

	p2 := new(Pessoa)
	p2.Nome = "Maria"

	fmt.Println(p2)
}
