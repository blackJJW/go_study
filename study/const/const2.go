package main

import "fmt"

const Pig int = 0
const Cow int = 1
const Chicken int = 2

func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("꿀꿀")
	} else if animal == Cow {
		fmt.Println("음메")
	} else {
		fmt.Println("...")
	}
}

func main() {
	PrintAnimal(Cow)
	PrintAnimal(Pig)
	PrintAnimal(7)
	PrintAnimal(0)
}
