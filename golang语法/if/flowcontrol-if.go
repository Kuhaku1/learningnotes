package main

/*

if t1 {

} else {

}

if t2 {

} else if t3 {

} else if t4 {

} else {

}

*/

import "fmt"

var enabled, disabled = false, true
var name string = "zhangsan"

func main() {

	if enabled == true {
		fmt.Println(enabled)
	} else {
		fmt.Println(disabled)

	}

	if name == "zhangsan" {
		fmt.Println("error , my name is zhangsan")
	} else if name == "wangwu" {
		fmt.Println("error , my name is wangwu ")
	} else {
		fmt.Println("success! ")
	}

}
