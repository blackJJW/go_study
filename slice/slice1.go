package main

import "fmt"

// slice는 go에서 제공하는 동적 배열 타입

func main() {
	slice := []int{1, 2, 3}

	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)

	}

	slice[0] = 10
	fmt.Println(slice)
}
