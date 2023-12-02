package main

import (
	"fmt"
	"log"

	"github.com/blackJJW/learngo/mini-proj/bank/banking"
)

func main() {
	account := banking.NewAccount("nico")
	fmt.Println(account)
	account.Deposit(10)
	fmt.Println(account)
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}
