package pubsub

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Producer(channel chan int, n int) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		channel <- i
	}
	close(channel)
}

func Consumer(channel chan int) {
	defer wg.Done()
	for i := range channel {
		fmt.Println("consumer:", i)
	}
}

func Print(n int) {
	channel := make(chan int)
	wg.Add(2)
	go Consumer(channel)
	go Producer(channel, n)
	wg.Wait()
}
