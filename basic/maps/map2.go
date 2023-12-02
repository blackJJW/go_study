package main

import (
	"fmt"
)

type Product struct {
	Name  string
	Price int
}

func main() {
	m := make(map[int]Product)

	m[16] = Product{"a", 500}
	m[46] = Product{"b", 200}
	m[78] = Product{"c", 1000}
	m[345] = Product{"d", 3000}
	m[897] = Product{"e", 500}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
