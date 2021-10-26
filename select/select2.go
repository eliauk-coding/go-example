package main

import (
	"fmt"
	"time"
)

func main() {
	oChan1 := make(chan string)
	oChan2 := make(chan string)
	go func(ch chan string) {
		time.Sleep(time.Second * 5)
		ch <- "test1"
	}(oChan1)

	go func(ch chan string) {
		time.Sleep(time.Second * 2)
		ch <- "test2"
	}(oChan2)
	select {
	case s1 := <-oChan1:
		fmt.Println("s1=", s1)
	case s2 := <-oChan2:
		fmt.Println("s2=", s2)
	}
}
