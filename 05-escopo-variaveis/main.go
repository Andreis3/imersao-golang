package main

import "fmt"

type Pessoa struct {
	Nome     string
	Idade    uint8
	Cpf      string
	Endereco string
	Cep      string
}

var nome string = "João"

func main() {
	//nome := "Maria"
	var nome string
	if len(nome) == 0 {
		nome = "Maria"
	}
	fmt.Println(nome)
	fmt.Println("=====================")
	showName()
}

func showName() {
	fmt.Println(nome)
}
