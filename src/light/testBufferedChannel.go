package light

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10
)

func init() {
	rand.Seed(time.Now().Unix())
}

// TestBufferedChan : test buffered channel
func TestBufferedChan() {
	var wg sync.WaitGroup
	tasks := make(chan string, taskLoad)

	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr, &wg)
	}
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	close(tasks)

	wg.Wait()
}

func worker(tasks chan string, worker int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}
		fmt.Printf("Worker: %d : Started %s\n", worker, task)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}
