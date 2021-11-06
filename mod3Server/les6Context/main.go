package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	defer func() { fmt.Println(time.Since(t)) }()
	ctx := context.Background()

	ctx = context.WithValue(ctx, "id", 1)

	//context.TODO()
	//ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	//
	//go func() {
	//	time.Sleep(time.Millisecond * 100)
	//	cancel()
	//}()

	pars(ctx)
}

func pars(ctx context.Context) {
	id := ctx.Value("id")
	//context.WithDeadline(ctx,time.Now())
	fmt.Println(id)
	for {
		select {
		case <-time.After(time.Second * 2):
			fmt.Println("parsing completed")
			return
		case <-ctx.Done():
			fmt.Println("deadline exceeded on select")
			return
		}

	}
}
