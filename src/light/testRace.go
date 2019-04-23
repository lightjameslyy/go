package light

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var counter int64
var wg sync.WaitGroup

// TestRace : test racing condition
func TestRace() {
	wg.Add(2)

	go incCounterSafe(1)
	go incCounterSafe(2)

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounterUnsafe(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter

		runtime.Gosched()

		value++

		counter = value
	}
}

func incCounterSafe(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}
