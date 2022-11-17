package main

import "fmt"

func main() {
	names := []string{"a", "b"}
	names = append(names, "c")
	fmt.Println(names)
}
