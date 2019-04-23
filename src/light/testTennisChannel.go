package light

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// TestPlayTennis : test unbuffered channel
func TestPlayTennis() {
	var wg sync.WaitGroup

	court := make(chan int)

	wg.Add(2)

	go player("Nadal", court, &wg)
	go player("Djokovic", court, &wg)

	court <- 1

	wg.Wait()
}

func player(name string, court chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		ball, ok := <-court

		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			close(court)
			return
		}

		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		court <- ball
	}
}
