package main

import (
	"fmt"
	"strconv"
	"sync"
)

//sync.map
var mp = sync.Map{}
var wg sync.WaitGroup

func main() {
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			mp.Store(key, n)         //必须使用sync.map内置的store方法设置键值对
			value, _ := mp.Load(key) //必须使用sync.map提供的Load方法根据键值对取值
			fmt.Printf("key:%d, value:%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
