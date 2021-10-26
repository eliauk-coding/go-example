package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	TestMutexFairness()
}

func TestMutexFairness() {
	var mu sync.Mutex
	stop := make(chan bool)
	defer close(stop)
	go func() {
		for {
			mu.Lock()
			log.Println("A acquire Mutex")
			time.Sleep(100 * time.Microsecond)
			mu.Unlock()
			select {
			case <-stop:
				return
			default:
			}
		}
	}()
	done := make(chan bool, 1)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(100 * time.Microsecond)
			mu.Lock()
			log.Printf("B%d: acquire Mutex\n", i)
			mu.Unlock()
		}
		done <- true
	}()
	select {
	case <-done:
	case <-time.After(10 * time.Second):
		log.Fatalf("can't acquire Mutex in 10 seconds")
	}
}
