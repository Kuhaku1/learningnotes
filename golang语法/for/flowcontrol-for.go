package main

/*

for 初始语句; 条件表达式; 结束语句{
	循环体语句
}

*/
import (
	"fmt"
)

func main() {
	//忽略初始语句

	for number := 2; number > 10; number++ {
		fmt.Println(number)
	}

	//忽略条件表达式
	age := 10

	// break 跳出for循环
	// cintiune 跳出本次循环
	for ; ; age++ {
		fmt.Println(age)
		if age > 20 {
			break
		}
		if age%2 == 0 {
			continue
		}

	}

	//结束语句
	// 死循环
	boys := 20
	for {
		fmt.Println(boys)
	}

	number := 1

	for ; number <= 9; number++ {
		x := 1
		for ; x <= number; x++ {
			fmt.Printf("%d*%d=%d ", x, number, x*number)
		}

		fmt.Println()
	}

}
