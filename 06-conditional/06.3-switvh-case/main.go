package main

func main() {
	idade := 3
	altura := 90

	if idade >= 6 {
		println("Pode brincar com acompanhante")
	} else if idade >= 4 && altura > 100 {
		println("Pode brincar")
	} else {
		println("Não pode brincar")
	}

	switch {
	case idade >= 6:
		println("Pode brincar com acompanhante")
	case idade >= 4 && altura > 100:
		println("Pode brincar")
	default:
		println("Não pode brincar")
	}
}
