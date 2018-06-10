package main

import (
	"os"
	"syscall"
	"os/signal"
	"fmt"
)

//信号量
func main()  {
	//Go信号通知通过向一个channel发送``os.Signal`来实现
	sigs := make(chan os.Signal,1)
	done := make(chan bool,1) //接收到信号量标志

	//设置可接收的信号量
	signal.Notify(sigs,syscall.SIGINT,syscall.SIGKILL,syscall.SIGTERM)

	//这个goroutine阻塞等待信号的到来,当信号到来的时候
	go func(){
		sig := <-sigs
		fmt.Println("接收到信号量",sig)
		done <- true
	}()

	fmt.Println("等待信号量的到来")
	<-done
	fmt.Println("程序将退出")
}

/**当程序运行后,按住ctrl+c终端程序
等待信号量的到来
^C接收到信号量 interrupt
程序将退出
 */