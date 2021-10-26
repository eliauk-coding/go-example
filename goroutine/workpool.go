package main

import (
	"fmt"
	"math/rand"
)

type job struct {
	Id      int
	RandNum int
}

type Result struct {
	job *job
	sum int
}

func main() {
	// new two channel
	jobChan := make(chan *job, 128)
	resultChan := make(chan *Result, 128)
	createPool(64, jobChan, resultChan)
	go func(resultChan chan *Result) {
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id, result.job.RandNum, result.sum)
		}
	}(resultChan)
	var id int
	for {
		id++
		r_num := rand.Int()
		job := &job{
			Id:      id,
			RandNum: r_num,
		}
		jobChan <- job
	}
}

// create pool
func createPool(num int, jobChan chan *job, resultChan chan *Result) {
	for i := 0; i < num; i++ {
		go func(jobChan chan *job, resultChan chan *Result) {
			for job := range jobChan {
				// 随机数接过来
				r_num := job.RandNum
				// 随机数每一位相加
				// 定义返回值
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				// 想要的结果是Result
				r := &Result{
					job: job,
					sum: sum,
				}
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}
