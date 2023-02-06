package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Println("Hello World")

	const name string = "JJW"
	fmt.Println(name)

	//var na string = "ABC"
	//na := "ABC" // func 안에서만 작동
	// 밖에서는 작동 안함

	fmt.Println(multiply(2, 2))

	totalLength, upperName := lenAndUpper("blackJJW")

	fmt.Println(totalLength, upperName)

	repeatMe("ABC", "DEF", "GHI")
}
