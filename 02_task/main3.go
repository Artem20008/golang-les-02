package main

import "fmt"

func main() {

	var A int
	var B uintptr = 5

	D := &B
	A = int(*D)
	fmt.Println(A)

}
