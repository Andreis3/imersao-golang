package main

import "fmt"

func main() {
	slice := make([]string, 0)

	array := [3]string{"a", "b", "c"}

	slice = append(slice, "a")
	slice = append(slice, "b")
	slice = append(slice, "c")

	slice2 := make([]string, 2)

	slice2[0] = "d"
	slice2[1] = "e"
	slice2 = append(slice2, "f")

	fmt.Println(slice, len(slice), cap(slice))

	fmt.Println(slice2, len(slice2), cap(slice2))

	fmt.Println(array, len(array), cap(array))
}
