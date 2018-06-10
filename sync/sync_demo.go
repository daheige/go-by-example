package main

import (
	"sync"
	"fmt"
	"math/rand"
	"sync/atomic"
	"runtime"
	"time"
)

func main()  {
	//互斥操作
	var mutex = &sync.Mutex{}
	var state = make(map[int]int)
	var ops  int64 = 0

	//生产者
	for i := 0; i < 10; i++ {
		go func(i int) {
			key := i
			val := rand.Intn(100)+10
			//fmt.Println(key,val)
			mutex.Lock()
			fmt.Println("写入的值是",val)
			state[key] = val
			mutex.Unlock()
			atomic.AddInt64(&ops,	1)
			runtime.Gosched() //手动释放资源
		}(i)
	}

	//消费者
	for i := 0; i < 10; i++ {
		go func(i int) {
			key := i
			mutex.Lock()
			fmt.Println("读取的值是: ",state[key])
			mutex.Unlock()
			atomic.AddInt64(&ops,	1)
			runtime.Gosched()
		}(i)
	}

	time.Sleep(time.Second) //主协程Sleep,让那10个协程能够运行一段时间
	fmt.Println("hello")
	opsFinal	:=	atomic.LoadInt64(&ops)
	fmt.Println(opsFinal)

	//枷锁获取state
	mutex.Lock()
	fmt.Println(state)
	mutex.Unlock()

}
/**
写入的值是 91
写入的值是 97
写入的值是 57
写入的值是 69
写入的值是 91
写入的值是 28
写入的值是 35
写入的值是 50
写入的值是 66
写入的值是 10
读取的值是:  91
读取的值是:  97
读取的值是:  57
读取的值是:  69
读取的值是:  91
读取的值是:  28
读取的值是:  35
读取的值是:  50
读取的值是:  66
读取的值是:  10
hello
20
map[4:91 5:28 6:35 0:91 2:57 7:50 8:66 9:10 1:97 3:69]
 */

