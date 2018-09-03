package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	fmt.Println("Waiting")

	for {
		if v, ok := <-ch; !ok {
			break
		} else {
			fmt.Println(5 - v)
		}
	}
	fmt.Println("DONE")
}
