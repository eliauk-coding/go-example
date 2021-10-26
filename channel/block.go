package main

import "runtime"

func Do(str string) {
	for {
		// 防止本协程霸占CPU不放
		println(str)
		runtime.Gosched()
	}
}

func main() {
	go Do("A")
	go Do("B")
	select {}
}
