package main

import "fmt"

func add(a int) int {
	a = a + 1
	return a
}

func add2(a *int) int {
	*a = *a + 1
	return *a
}

func main() {

	//初始化number等于20
	number := 20
	//打印number值
	fmt.Println("number = ", number)

	//调用方法
	number2 := add(number)
	fmt.Println("number + 1 = ", number2)

	//打印number
	fmt.Println("number = ", number)

	number3 := add2(&number)
	fmt.Println("number + 1 =  ", number3)
	fmt.Println("number is = ", number)

	// 匿名函数
	result := func(a int, b int) { fmt.Println(a, b) }
	result(2, 7)

}
