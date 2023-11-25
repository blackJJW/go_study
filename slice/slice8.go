package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2

	copy(slice[idx:], slice[idx+1:])

	slice = slice[:len(slice)-1]

	fmt.Println("slice", slice)
}
