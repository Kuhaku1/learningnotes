package main

import "fmt"

var options string = "restart"

func main() {

	switch content := "Python"; content {
	case "Python":
		fmt.Println("I am python!")
	case "Shell":
		fmt.Println("I am Shell")
	default:
		fmt.Println("use [Python|Shell]")

	}

	switch options {
	case "start":
		fmt.Printf("开始%sService ", options)

	case "stop":
		fmt.Printf("开始%sService ", options)

	case "restart":
		fmt.Printf("开始%sService ", "stop")
		fmt.Printf("开始%sService ", "start")

	default:
		fmt.Println("Please use [start|stop|restart]")

	}

}
