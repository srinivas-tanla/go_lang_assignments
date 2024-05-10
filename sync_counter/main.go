package synccounter

import (
	"fmt"
	"sync"
)

var count int

var lock sync.Mutex
var wg sync.WaitGroup

func increment() {
	defer wg.Done()
	lock.Lock()
	count++
	lock.Unlock()
}

func Counter(n int) {
	for i := 0; i < n; i++ {
		wg.Add(1)
		go increment()
	}
	wg.Wait()
	fmt.Println("increment value:", count)
}
