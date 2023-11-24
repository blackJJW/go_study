package main

import "fmt"

func HasRichFriend() bool {
	return true
}

func GetFriendCount() int {
	return 3
}

func main() {
	price := 35000

	if price >= 50000 {
		if HasRichFriend() {
			fmt.Println("신발끈 풀림")
		} else {
			fmt.Println("더치페이")
		}
	} else if price >= 30000 {
		if GetFriendCount() > 3 {
			fmt.Println("어이쿠 신팔끈이")
		} else {
			fmt.Println("더치2")
		}
	} else {
		fmt.Println("더치3")
	}
}
