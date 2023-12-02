package main

import "fmt"

func main() {
	a := 2
	b := &a
	a = 5

	*b = 20

	fmt.Println(&a, b)
	fmt.Println(a, b)
	fmt.Println(*b)
}
