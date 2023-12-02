package main

import (
	"fmt"
	"time"
)

func main() {

	go s_Count("nico")
	go s_Count("Lynn")

	c := make(chan string)
	people := [5]string{"nico", "Lynn", "A", "B", "C"}
	for _, person := range people {
		go is_s(person, c)
	}

	for i := 0; i < len(people); i++ {
		fmt.Print("waiting for ", i, " ")
		fmt.Println(<-c)
	}

}

func s_Count(person string) {
	for i := 1; i < 10; i++ {
		fmt.Println(person, "-- s", i)
		time.Sleep(time.Second)
	}
}

func is_s(person string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- person + "------- s"
}
