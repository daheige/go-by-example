package main

import (
	"fmt"
	"time"
)

func main(){
	done := make(chan bool,1)
	go worker(done)
	<-done //一直会等待数据写入完毕才解除阻塞,获取到通道中的done值
}

//done通道是单向的,只能写入数据
func worker(done chan <- bool){
	fmt.Println("开始写入数据")
	time.Sleep(time.Second)
	fmt.Println("写入完成")
	done <- true
}