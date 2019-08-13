package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	a := fibonacci()
	for index := 0; index < 10; index++ {
		val := a()
		fmt.Println(val)
	}
}
