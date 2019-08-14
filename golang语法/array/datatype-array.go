package main

import "fmt"

//array
var students [5]string
var teachers [6]int

func main() {
	students[0] = "xiaoming"
	teachers[0] = 12

	students[3] = "test"
	teachers[5] = 11
	fmt.Println(students, teachers)

	boys := [5]string{"li", "wang", "xu", "zhang", "zhou"}
	girls := [...]int{1, 2, 3, 4, 4, 5, 55}
	fmt.Println(boys, girls)

	lengths := len(boys)
	fmt.Println(lengths)

	fmt.Println(boys[4])

	fmt.Println(boys[2:4])
	fmt.Println(boys[:4])
	fmt.Println(boys[:])

}
