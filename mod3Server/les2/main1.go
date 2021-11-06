package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(time.Second)
		fmt.Println("concurrently ninja with delay")
	}()
	go fmt.Println("concurrently ninja")
	go fmt.Println("concurrently ninja")
	time.Sleep(time.Second * 2)
	fmt.Println("not concurrently ninja")
}
