package main

import "fmt"

var counter = func(n int) chan<- chan<- int {
	requests := make(chan chan<- int)
	go func() {
		for request := range requests {
			if request == nil {
				// 递增计数
				n++
			} else {
				// 返回当前计数
				fmt.Println("return")
				request <- n
			}
		}
	}()
	// 隐式转换到类型chan<- (chan<- int)
	return requests
}(0)

func main() {
	fmt.Println(cap(counter))
	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			counter <- nil
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go increase1000(done)
	go increase1000(done)
	<-done
	<-done
	fmt.Println(cap(counter))
	request := make(chan int, 1)
	counter <- request
	// 2000
	fmt.Println(<-request)
}
