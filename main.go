package main

import (
	"fmt"
	"sync"

	"github.com/srinivas365/day2/pipeline"
	"github.com/srinivas365/day2/pubsub"
	synccounter "github.com/srinivas365/day2/sync_counter"
	"github.com/srinivas365/day2/timeout"
)

// type Container struct {
// 	counter int
// 	mu      sync.Mutex
// }

var counter int
var wg sync.WaitGroup

var lock sync.Mutex

func foo() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		lock.Lock()
		counter++
		lock.Unlock()
		// time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
	}
}

func boo() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		lock.Lock()
		counter--
		lock.Unlock()
		// time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

// var cg sync.WaitGroup

func worker(channel chan int, low int, high int) {
	val := 0
	for i := low; i < high; i++ {
		val += i
	}
	channel <- val
}

func CalculateConcurrently(n int) int {
	channel := make(chan int)

	go worker(channel, 0, n/2)
	go worker(channel, n/2, n+1)

	count := 0
	for i := 0; i < 2; i++ {
		count += <-channel
	}

	// fmt.Printf("sum of %d numbers: %d \n", n, count)
	return count
}

func Calculate(n int) int {
	count := 0
	for i := 0; i < n+1; i++ {
		count += i
	}
	return count
}

func main() {
	fmt.Println("hello world")
	wg.Add(2)
	go foo()
	go boo()
	wg.Wait()
	fmt.Println("counter:", counter)
	CalculateConcurrently(9)                               // assign 1
	pubsub.Print(10)                                       // assign 2
	synccounter.Counter(10)                                // assign 3
	timeout.Test(1)                                        // assign 4
	timeout.Test(5)                                        // assign 4
	pipeline.Print(pipeline.Square(pipeline.Generate(10))) // assign 5
}
