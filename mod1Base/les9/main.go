package main

import (
	"context"
	"fmt"
	"github.com/zhashkevych/scheduler"
	"os"
	"os/signal"
	"time"
)

func main() {
	t := time.Month(2)
	fmt.Println(t)

	z := time.Local
	s := time.Date(1970, 12, 2, 3, 4, 5, 224, z)
	fmt.Println(s)
	ctx := context.Background()

	worker := scheduler.NewScheduler()
	worker.Add(ctx, parseSubscriptionData, time.Second*5)
	worker.Add(ctx, sendStatistics, time.Second*10)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit
	worker.Stop()
}

func parseSubscriptionData(ctx context.Context) {
	time.Sleep(time.Second * 1)
	fmt.Printf("subscription parsed successfuly at %s\n", time.Now().String())
}

func sendStatistics(ctx context.Context) {
	time.Sleep(time.Second * 5)
	fmt.Printf("statistics sent at %s\n", time.Now().String())
}
