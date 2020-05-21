package main

import (
	"fmt"
	"log"

	"stacew/gostudy/nomadcoders/2.0.BankProjects/accounts"
)

func main() {
	account := accounts.NewAccount("yslee")
	fmt.Println(account)

	account.Deposit(100)
	fmt.Println(account.Balance())

	err := account.WithDraw(10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())

	err = account.WithDraw(1000)
	if err != nil {
		fmt.Println(err)
		//log.Fatalln(err)
	}
	fmt.Println(account.Balance())

	account.ChangeOwner("zzz")
	fmt.Println(account.Balance(), account.Owner())

	fmt.Println(account)

}
