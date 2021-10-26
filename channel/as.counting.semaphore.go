package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	server1()
}

type Seat int
type Bar chan Seat

func (bar Bar) ServeCustomer(consumers chan int) {
	for c := range consumers {
		seatId := <-bar
		//log.Print("顾客#", c, "进入酒吧")
		//log.Print("++ customer#", c, " drinks at seat#", seat)
		log.Print("++ 顾客#", c, "在第", seatId, "个座位开始饮酒")
		time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))
		log.Print("-- 顾客#", c, "离开了第", seatId, "个座位")
		bar <- seatId // 释放座位离开酒吧
	}
}

func server1() {
	rand.Seed(time.Now().UnixNano())
	// 此酒吧有10个座位
	bar24x7 := make(Bar, 10)
	// 摆放10个座位
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		// 均不会阻塞
		bar24x7 <- Seat(seatId)
	}

	consumers := make(chan int)
	for i := 0; i < cap(bar24x7); i++ {
		go bar24x7.ServeCustomer(consumers)
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		consumers <- customerId
	}
}
