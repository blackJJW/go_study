package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("입력")

		var number int
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("숫자를 입력")

			stdin.ReadString('\n')
			continue
		}
		fmt.Printf("입력된 숫자 : %d\n", number)

		if number%2 == 0 {
			break
		}
	}
	fmt.Println("for exited")
}
