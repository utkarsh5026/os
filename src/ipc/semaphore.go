package ipc

import (
	"fmt"
	"sync"
)

type Semaphore struct {
	c chan struct{}
}

func NewSemaphore(capacity int) *Semaphore {
	return &Semaphore{c: make(chan struct{}, capacity)}
}

func (s *Semaphore) Acquire() {
	s.c <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.c
}

type CountingSemaphore struct {
	capacity int
	channel  chan struct{}
	mutex    sync.Mutex
	count    int
}

func NewCountingSemaphore(capacity int) *CountingSemaphore {
	return &CountingSemaphore{
		capacity: capacity,
		channel:  make(chan struct{}, capacity),
	}
}

func (s *CountingSemaphore) Release() {
	<-s.channel
	s.mutex.Lock()
	s.count--
	fmt.Printf("Resource released. Current count: %d\n", s.count)
	s.mutex.Unlock()
}

func (s *CountingSemaphore) Acquire() {
	s.channel <- struct{}{}
	s.mutex.Lock()
	s.count++
	fmt.Printf("Resource acquired. Current count: %d\n", s.count)
	s.mutex.Unlock()
}
