package main

import "fmt"

func main() {
	animal := "cachorro"

	if animal == "gato" {
		fmt.Println("miau")
	} else {
		fmt.Println("au au")
	}

	utensilio := "garfo"

	if utensilio == "garfo" {
		fmt.Println("comer")
	} else if utensilio == "faca" {
		fmt.Println("cortar")
	} else {
		fmt.Println("nada")
	}
}
