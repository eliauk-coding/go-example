package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	//trySendAndRecv()
	//firstResponseWins()
	ticker()
}

func trySendAndRecv() {
	type Book struct {
		id int
	}
	bookShelf := make(chan Book, 3)
	for i := 0; i < cap(bookShelf)*2; i++ {
		select {
		case bookShelf <- Book{id: i}:
			fmt.Println("成功将书放在书架上", i)
		default:
			fmt.Println("书架已经被占满了")
		}
	}
	for i := 0; i < cap(bookShelf)*2; i++ {
		select {
		case book := <-bookShelf:
			fmt.Println("成功从书架上取下一本书", book.id)
		default:
			fmt.Println("书架上已经没有书了")
		}
	}
}

// IsClosed 无阻塞地检查一个通道是否已经关闭
func IsClosed(c chan struct{}) bool {
	select {
	case <-c:
		return true
	default:
	}
	return false
}

func source(c chan<- int32) {
	ra, rb := rand.Int31(), rand.Intn(3)+1
	// 休眠1秒/2秒/3秒
	log.Println(ra, rb)
	time.Sleep(time.Duration(rb) * time.Second)
	select {
	case c <- ra:
	default:
	}
}

// firstResponseWins 最快回应
func firstResponseWins() {
	rand.Seed(time.Now().UnixNano())
	// 此通道容量必须至少为1
	c := make(chan int32, 1)
	for i := 0; i < 5; i++ {
		go source(c)
	}
	for {
		select {
		case rnd := <-c:
			fmt.Println(rnd)
			return
		default:
		}
	}
	// 只采用第一个成功返回的回应数据
	//rnd := <-c
	//fmt.Println(rnd)
}

func requestWithTimeout(timeout time.Duration) (int, error) {
	c := make(chan int)
	// 可能需要超出预期的时长回应
	go func(c chan int) {

	}(c)
	select {
	case data := <-c:
		return data, nil
	case <-time.After(timeout):
		return 0, errors.New("超时了！")
	}
}

func Tick(d time.Duration) <-chan struct{} {
	// 容量最好为1
	c := make(chan struct{}, 1)
	go func() {
		for {
			time.Sleep(d)
			select {
			case c <- struct{}{}:
			default:
			}
		}
	}()
	return c
}

func ticker() {
	t := time.Now()
	for range Tick(time.Second) {
		fmt.Println(time.Since(t))
	}
}
