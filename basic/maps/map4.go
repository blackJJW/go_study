package main

import (
	"fmt"
)

const M = 10

func hash(d int) int {
	return d % M
}

func main() {
	m := [M]string{}

	m[hash(23)] = "abc"
	m[hash(259)] = "def"

	fmt.Printf("%d = %s\n", 23, m[hash(23)])
	fmt.Printf("%d = %s\n", 259, m[hash(259)])
}
