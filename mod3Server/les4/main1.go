package main

import (
	"fmt"
	"time"
)

func main() {

	msg1 := make(chan string)
	msg2 := make(chan string)
	go func() {
		for {
			msg1 <- "ch1 passed 200 millisecond"
			time.Sleep(time.Millisecond * 200)
		}
	}()
	go func() {
		for {
			msg2 <- "ch1 passed 1 second"
			time.Sleep(time.Second)
		}
	}()
	for {
		select {
		case m1 := <-msg1:
			fmt.Println(m1)
		case m2 := <-msg2:
			fmt.Println(m2)
		default:
		}
	}
}
