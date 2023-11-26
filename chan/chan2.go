package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)

	go square()
	ch <- 9
	fmt.Println("Never Print")
}

func square() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("sleep")
	}
}
