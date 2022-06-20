package main

import "fmt"

func main() {

	var A int = 5.0
	var B uintptr

	D := &A
	B = uintptr(*D)

	fmt.Println(B)
}
