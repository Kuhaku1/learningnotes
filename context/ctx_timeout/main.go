package main

import (
	"context"
	"fmt"
	"time"
)

func process() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	c := make(chan interface{}, 1)

	go func() {
		// 做任务
		// for {
		// 	time.Sleep(time.Second)
		// }
		c <- "vghh"
	}()
	select {
	case <-ctx.Done():
		fmt.Println("Timeout! err:")
	case res := <-c:
		fmt.Printf("Server Response: %#v\n", res)
	}
	return
}
func main() {
	process()
}
