package main

import "fmt"

func main() {
	fmt.Println("start")

	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(0)

	fmt.Println("done")
}
