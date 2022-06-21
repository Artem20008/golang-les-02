package main

import (
	"fmt"
	"strconv"
)

func main() {
	var stringToInt string = "104"
	var intToString int = 34

	w, _ := strconv.Atoi(stringToInt)
	a := strconv.Itoa(intToString)

	fmt.Printf("%d \n", w)
	fmt.Printf("%q \n", a)
}
