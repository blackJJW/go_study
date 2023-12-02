package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string)

	m["abc"] = "d"
	m["def"] = "g"
	m["ghi"] = "j"
	m["jkl"] = "m"

	m["jkl"] = "k"

	fmt.Printf("%s\n", m["abc"])
	fmt.Printf("%s\n", m["jkl"])
}
