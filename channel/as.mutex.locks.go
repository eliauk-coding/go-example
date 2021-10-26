package main

import "fmt"

func main() {
	lockThroughSend()
	lockThroughRecv()
}

func lockThroughSend() {
	// The capacity must be one.
	mutex := make(chan struct{}, 1)

	counter := 0
	increase := func() {
		// lock
		mutex <- struct{}{}
		counter++
		// unlock
		<-mutex
	}
	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			increase()
		}
		done <- struct{}{}
	}
	done := make(chan struct{})
	go increase1000(done)
	go increase1000(done)
	<-done
	<-done
	fmt.Println(counter)
}

func lockThroughRecv() {
	// The capacity must be one.
	mutex := make(chan struct{}, 1)
	mutex <- struct{}{}

	counter := 0
	increase := func() {
		// lock
		<-mutex
		counter++
		// unlock
		mutex <- struct{}{}
	}
	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			increase()
		}
		done <- struct{}{}
	}
	done := make(chan struct{})
	go increase1000(done)
	go increase1000(done)
	<-done
	<-done
	fmt.Println(counter)
}
