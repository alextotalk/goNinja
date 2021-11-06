package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	defer func() { fmt.Println(time.Since(t)) }()
	jobNum, workerCount := 15, 30
	job := make(chan int, jobNum)
	result := make(chan int, jobNum)
	for i := 0; i < jobNum; i++ {
		go worker(i, job, result)
	}
	for i := 0; i < workerCount; i++ {
		job <- i + 1
	}
	close(job)
	for i := 0; i < jobNum; i++ {
		fmt.Printf("Result num %d and vol %d \n", i+1, <-result)
	}
	close(result)
}

func worker(i int, job <-chan int, result chan<- int) {
	for v := range job {
		time.Sleep(time.Second)
		fmt.Printf("Worker num %d  and job %d \n", i+1, v)
		result <- v * v
	}

}
