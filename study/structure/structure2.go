package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"avb", "hana", 23}
	vip := VIPUser{
		User{"vbfds", "hwarang", 48},
		3,
		250,
	}

	fmt.Printf("%v\n", user)
	fmt.Printf("%v\n", vip)
}
