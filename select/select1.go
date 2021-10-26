package main

import (
	"fmt"
	"time"
)

func main() {
	// create channel
	oChan := make(chan string, 10)
	go write(oChan)
	for c := range oChan {
		fmt.Println("res:", c)
		time.Sleep(time.Second * 2)
	}
	close(oChan)
}

func write(ch chan string) {
	var i int
	value := "hello"
	for {
		select {
		case ch <- value:
			fmt.Printf("write %v\n", value)
			i++
			value = fmt.Sprintf("hello%d", i)
			if i > 20 {
				return
			}
		default:
			fmt.Println("channel full")
		}
		time.Sleep(time.Second * 1)
	}
}
