package main

import "fmt"

var (
	idade uint8  = 20
	cpf   string = "123.456.789-00"
)

func main() {
	var (
		nome     string
		endereco string
	)

	nome = "João"
	endereco = "Rua 1"

	nome2 := "Maria"

	fmt.Println(nome, idade, cpf, endereco)
	fmt.Println(nome2)
}
