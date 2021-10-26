package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	TestAtomicAdd()
	TestAtomicValue()
	TestCAS()
	var count int32 = 9
	atomic.AddInt32(&count, -1)
	fmt.Println(count)
}

func TestAtomicAdd() {
	var count int32
	fmt.Println("main start...")
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("goroutine:", i, "start...")
			atomic.AddInt32(&count, 1)
			fmt.Println("goroutine:", i, "count:", count, "end...")
		}(i)
	}
	wg.Wait()
	fmt.Println("main end...")
}

func TestAtomicValue() {
	var v atomic.Value
	type info struct {
		num  int32
		num2 int64
	}
	in := info{num: 1}
	v.Store(&in)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			atomic.AddInt32(&in.num, 1)
			v.Store(&in)
		}
	}()
	for i := 0; i < 10; i++ {
		go func() {
			for {
				res := v.Load()
				fmt.Println(res)
			}
		}()
	}

	select {}
	fmt.Println("main end...")
}

func TestCAS() {
	var share uint64 = 1
	wg := sync.WaitGroup{}
	wg.Add(3)
	// 协程1，期望值是1,欲更新的值是2
	go func() {
		defer wg.Done()
		swapped := atomic.CompareAndSwapUint64(&share, 1, 2)
		fmt.Println("goroutine 1", swapped)
	}()
	// 协程2，期望值是1，欲更新的值是2
	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Millisecond)
		swapped := atomic.CompareAndSwapUint64(&share, 1, 2)
		fmt.Println("goroutine 2", swapped)
	}()
	// 协程3，期望值是2，欲更新的值是1
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Millisecond)
		swapped := atomic.CompareAndSwapUint64(&share, 2, 1)
		fmt.Println("goroutine 3", swapped)
	}()
	wg.Wait()
	fmt.Println("main exit")
}
