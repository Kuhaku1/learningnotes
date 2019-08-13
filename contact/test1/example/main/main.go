package main

import (
	"time"
	"contact/test1/example/balance"
	"os"
)

func main() {
	name := "random"
	if len(os.Args) > 1 {
		name = os.Args[1]

	}
	for {
		balance.Dobalance(name, 5)
		time.Sleep(time.Second)
	}
}
