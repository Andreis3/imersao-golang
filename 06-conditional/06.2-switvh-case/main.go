package main

func main() {
	animal := "gato"

	switch animal {
	case "cachorro":
		println("au au")
	case "gato":
		println("miau")
	case "papagaio":
		println("oi")
	default:
		println("...")
	}
}
