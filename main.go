package main

import (
	"fmt"
	"strconv"
)

func main() {

	w, _ := strconv.Atoi("104")
	a := strconv.Itoa(35)

	fmt.Println("%q\n", w)
	fmt.Printf("%q\n", a)
}
