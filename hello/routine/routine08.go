package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	var mutex sync.Mutex
	wait := sync.WaitGroup{}

	fmt.Println("Locked Start")
	mutex.Lock()

	for i := 1; i <= 5; i++ {
		wait.Add(1)

		go func(i int) {
			fmt.Println("Not lock:", i)

			mutex.Lock()
			fmt.Println("Lock:", i)

			time.Sleep(time.Second)

			fmt.Println("Unlock:", i)
			mutex.Unlock()

			defer wait.Done()
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("Unlocked finish")
	mutex.Unlock()

	wait.Wait()

}
