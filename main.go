package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "raman"}
	choyeun := person{name: "choyeun", age: 18, favFood: favFood}
	fmt.Println(choyeun.name)
}

/**
go run main.go
git add .
git commit -m "."
git push
*/
