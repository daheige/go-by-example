//多协程原子访问计数器
package main

import (
	"sync/atomic"
	"time"
	"fmt"
	"runtime"
)

func main()  {
	var ops uint64 = 0

	//使用50个协程,对ops不停的计数
	for i :=0;i<50;i++{
		go func() {
			for{
				fmt.Println(ops)
				atomic.AddUint64(&ops,1)
				//允许其他的协程来处理
				runtime.Gosched()
			}
		}()
	}

	//让协程运行一段时间
	time.Sleep(time.Second)

	//`LoadUint64`读取uint64变量
	cnt := atomic.LoadUint64(&ops)
	fmt.Println("当前的ops是",cnt)

}

/**
607281
当前的ops是 607281
 */
