package main

import (
	"fmt"
	"main/accounts"
)

func main() {
	account := accounts.NewAccount("Bacon")
	fmt.Println(account)
	if account.Deposit(1000) {
		fmt.Println("Deposit succeeded!")
	}
	err := account.Withdraw(500)
	if err != nil {
		fmt.Println(err)
	}
	account.ChangeOwner("Hoon")
	fmt.Println(account.String())
}
