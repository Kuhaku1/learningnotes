package main

import "fmt"

//bool
var skipTest bool                   //全局环境变量
var enabled, disabled = true, false //忽略类型声明

//int
var high int = 180
var age int = 33
var aa = 123
var bb = 234
var cc = bb - aa

//float
var boyHigh float64 = 1.80
var girlHigh float64 = 1.60

//complex
var cNum1 complex64
var cNum2 complex128

//string
var myname string = "lisi"
var yourname string = `aaaa`

func main() {

	skipTest = true
	testa := false

	fmt.Println(testa)

	fmt.Printf("我的身高是 %d ,我的年龄是 %d", high, age)

	fmt.Println(cc)

	fmt.Printf("男孩身高是  %f , 女孩身高 %f", boyHigh, girlHigh)
	fmt.Println(yourname)
}
