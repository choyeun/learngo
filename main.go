package main

import (
	"fmt"
	"learngo/accounts"
)

func main() {
	account := accounts.NewAccount("choyeun")
	fmt.Println(account)
}

/**
go run main.go
git add .
git commit -m "."
git push
*/
