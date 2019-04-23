package light

import (
	"fmt"
	"runtime"
	"sync"
)

// TestGoroutine01 : test goroutine
func TestGoroutine01() {
	//runtime.GOMAXPROCS(1)
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func() {
		defer wg.Done()

		for count := 0; count < 10; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 10; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Waiting To Finish")

	wg.Wait()
	fmt.Println("\nTerminating Program")
}

// TestGoroutineSwitching : goroutines that take too much time will be switched
func TestGoroutineSwitching() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Create Goroutines")

	go printPrimes("A", &wg)
	go printPrimes("B", &wg)

	fmt.Println("Waiting To Finish")

	wg.Wait()
	fmt.Println("\nTerminating Program")
}

func printPrimes(prefix string, wg *sync.WaitGroup) {
	defer wg.Done()

next:
	for outer := 2; outer < 2000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s: %d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
