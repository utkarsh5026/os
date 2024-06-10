package main

import (
	"fmt"
	"osimpl/ipc"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	semaphore := ipc.NewSemaphore(1)

	process := func(id int) {
		defer wg.Done()
		fmt.Printf("Process %d waiting to enter critical section\n", id)
		semaphore.Acquire() // Acquire a resource
		fmt.Printf("Process %d entered critical section\n", id)
		time.Sleep(2 * time.Second) // Simulate critical section work
		fmt.Printf("Process %d leaving critical section\n", id)
		semaphore.Release() // Release the resource
	}

	wg.Add(6)
	for i := 1; i <= 6; i++ {
		go process(i)
	}

	wg.Wait()
}
