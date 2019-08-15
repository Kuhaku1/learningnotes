package main

import "fmt"

//map
var students map[string]int

func main() {
	students := make(map[string]int)
	students["name"] = 345

	delete(students, "name")
	fmt.Println(students)
	fmt.Println(students["sdfghgfd"])
	v, ok := students["sdfg"]
	if !ok {
		fmt.Println(ok)
	}
	fmt.Println(v)

}
