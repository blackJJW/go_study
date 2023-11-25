package main

import "fmt"

func main() {
	temp := 33

	if temp > 28 {
		fmt.Println("Ac on")
	} else if temp <= 3 {
		fmt.Println("heater on")
	} else if temp <= 18 {
		fmt.Println("get out")
	} else {
		fmt.Println("hot")
	}
}
