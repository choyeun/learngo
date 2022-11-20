package main

import (
	"fmt"
	"learngo/banking"
)

func main() {
	account := banking.Account{Owner: "choyeun", Balance: 1000}
	fmt.Println(account)
}

/**
go run main.go
git add .
git commit -m "."
git push
*/
