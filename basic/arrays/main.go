package main

import "fmt"

func main() {
	//names := [5]string{"nico", "lynn", "dal"}
	names := []string{"nico", "lynn", "dal"}

	names = append(names, "ABC")

	fmt.Println(names)
}
