package main

import (
	"fmt"
)

func Asum(base int) func(int) int {
	sum := 0
	return func(a int) int {
		sum = sum + a
		return sum
	}
}

type Iasum func(int) (int, Iasum)

func orthodoxfunction(base int) Iasum {
	return func(v int) (int, Iasum) {
		return base + v, orthodoxfunction(base + v)
	}
}

func main() {
	fmt.Println("函数式编程")
	a := Asum(0)
	for i := 0; i < 10; i++ {
		fmt.Println(i, "  ", a(i))
	}

	fmt.Println("正统函数式编程")
	box := orthodoxfunction(0)
	var s int
	for index := 0; index < 10; index++ {
		s, box = box(index)
		fmt.Println(s)
	}
}
