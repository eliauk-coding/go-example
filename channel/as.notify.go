package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"time"
)

func main() {
	//oneToOneUseSend()
	//oneToOneUseRecv()
	//nToOneAndOneToN()
	timerNotify()
}

// one-to-one notification use send
func oneToOneUseSend() {
	values := make([]byte, 32*1034*1024)
	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	done := make(chan struct{})
	// The sorting goroutine
	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		// Notify sorting is done.
		done <- struct{}{}
	}()
	// do some other things ...
	<-done
	fmt.Println(values[0], values[len(values)-1])
}

// one to one notification use recv
func oneToOneUseRecv() {
	done := make(chan struct{})
	// The capacity of the signal channel can
	// also be one. If this is true, then a
	// value must be sent to the channel before
	// creating the following goroutine.
	go func() {
		fmt.Println("start")
		time.Sleep(time.Second * 2)
		// Receive a value from the done
		// channel, to unblock the second
		// send in main goroutine.
		<-done
	}()
	// Blocked here, wait for a notification.
	done <- struct{}{}
	fmt.Println("end")
}

func worker(id int, ready <-chan struct{}, done chan<- struct{}) {
	<-ready
	log.Print("worker#", id, "开始工作")
	time.Sleep(time.Second * time.Duration(id+1))
	log.Print("worker#", id, "工作完成")
	// Notify the main goroutine (N-to-1),
	done <- struct{}{}
}

// N-to-1 and 1-to-N notifications
func nToOneAndOneToN() {
	log.SetFlags(0)

	ready, done := make(chan struct{}), make(chan struct{})
	go worker(1, ready, done)
	go worker(2, ready, done)
	go worker(3, ready, done)

	time.Sleep(time.Second * 3 / 2)
	// 1-to-N notifications.
	ready <- struct{}{}
	ready <- struct{}{}
	ready <- struct{}{}
	// Being N-to-1 notified.
	<-done
	<-done
	<-done
}

func afterDuration(d time.Duration) <-chan struct{} {
	c := make(chan struct{}, 1)
	go func() {
		time.Sleep(d)
		c <- struct{}{}
	}()
	return c
}

// timer notification
func timerNotify() {
	fmt.Println("Hi!")
	<-afterDuration(time.Second)
	fmt.Println("hello!")
	<-afterDuration(time.Second)
	fmt.Println("Bye!")
}
