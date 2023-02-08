package main

import (
	"fmt"
	"time"
)

func main() {
	go s_Count("nico")
	go s_Count("Lynn")
	time.Sleep(time.Second * 5)
}

func s_Count(person string) {
	for i := 1; i < 10; i++ {
		fmt.Println(person, "-- s", i)
		time.Sleep(time.Second)
	}
}
