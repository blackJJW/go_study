package main

import "fmt"

func add(a, b int) int {
	return a + b

}
func mul(a, b int) int {
	return a * b
}

type OpFn func(int, int) int

func getOperator(op string) OpFn {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

func main() {
	var operator func(int, int) int
	operator = getOperator("+")

	var result = operator(3, 4)
	fmt.Println(result)
}
