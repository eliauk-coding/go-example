package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	m := sync.Map{}
	go func() {
		for i := 0; i < 10000; i++ {
			m.Store(strconv.Itoa(i), i)
		}
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			fmt.Println(m.Load(strconv.Itoa(i)))
		}
	}()

	//unSyncM := make(map[string]int)
	//go func() {
	//	for i := 0; i < 10000; i++ {
	//		unSyncM[strconv.Itoa(i)] = i
	//	}
	//}()
	//go func() {
	//	for i := 0; i < 10000; i++ {
	//		fmt.Println(unSyncM[strconv.Itoa(i)])
	//	}
	//}()
	time.Sleep(time.Second * 20)
}
