package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//goroutine池
/*
编写代码实现一个计算机数的每个位置数字之和的程序，
要求使用goroutine和channel构建生产者模式和消费者模式，
可以指定启动的goroutine数量-worker pool模式
在工作中我们通常会使用work pool模式来控制goroutine的数量，防止goroutine泄漏和暴涨






使用goroutine和channel实现一个计算int64随机数各位数和的程序。
	1.开启一个goroutine循环生成int64类型的随机数，发送到jobChan
	2.开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
	3.主goroutine从resultChan取出结果并打印到终端输出
*/
type job struct {
	x int64
}
type result struct {
	job *job
	sum int64
}

func j(a chan<- *job) {
	defer wg.Done()
	defer close(a)
	// 循环生成int64类型的随机数，发送到jobChan\
	for {
		x := rand.Int63()
		newjob := &job{
			x: x,
		}
		a <- newjob
		time.Sleep(time.Millisecond * 500)
	}
}
func res(a <-chan *job, b chan<- *result) {
	defer wg.Done()
	defer close(b)

	// 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
	for {
		job := <-a
		sum := int64(0)
		n := job.x
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newres := &result{
			job: job,
			sum: sum,
		}
		b <- newres

	}
}

var wg sync.WaitGroup

func main() {
	jobChan := make(chan *job, 100)
	resultChan := make(chan *result, 100)
	wg.Add(1)
	go j(jobChan)
	wg.Add(24)
	for i := 1; i < 25; i++ {
		go res(jobChan, resultChan)
	}
	// 从resultChan取出结果并打印到终端输出
	for i := range resultChan {
		fmt.Println(i.job.x, i.sum)
	}
	wg.Wait()

}
