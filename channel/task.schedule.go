package main

import (
	"fmt"
	"time"
)

func printTask(msg string, recv, send chan struct{}) {
	for {
		<-recv
		fmt.Println(msg)
		send <- struct{}{}
		if msg == "A" {
			time.Sleep(time.Second)
		}
	}
}

func main() {
	chA := make(chan struct{}, 1)
	chB := make(chan struct{}, 1)
	chC := make(chan struct{}, 1)
	chD := make(chan struct{}, 1)
	chA <- struct{}{}
	go printTask("A", chA, chB)
	go printTask("B", chB, chC)
	go printTask("C", chC, chD)
	go printTask("D", chD, chA)
	for {

	}
}
