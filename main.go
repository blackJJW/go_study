package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true

}

func main() {
	fmt.Println("Hello World")

	const name string = "JJW"
	fmt.Println(name)

	//var na string = "ABC"
	//na := "ABC" // func 안에서만 작동
	// 밖에서는 작동 안함

	fmt.Println(multiply(2, 2))

	totalLength, up := lenAndUpper("blackJJW")

	fmt.Println(totalLength, up)

	repeatMe("ABC", "DEF", "GHI")

	result := superAdd(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(result)

	fmt.Println(canIDrink(16))
}
