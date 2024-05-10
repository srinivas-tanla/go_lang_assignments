package timeout

import (
	"fmt"
	"time"
)

func Test(n int) {
	done := make(chan bool)
	timer := time.After(3 * time.Second)
	go worker(done, n)
	select {
	case res := <-done:
		fmt.Println("worker:", res)
	case <-timer:
		fmt.Println("timeout...")
	}
}

func worker(done chan bool, n int) {
	time.Sleep(time.Second * time.Duration(n))
	done <- true
}
