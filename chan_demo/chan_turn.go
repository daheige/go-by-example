package main

import "fmt"

//通道的方向,通道可以只读,也可以只写

func main()  {
	pings := make(chan string,1)
	pongs := make(chan string,1)
	//写入通道
	pingWrite(pings,"daheige")
	pongRead(pings,pongs) //将pings通道中的值写入pongs中

	fmt.Println(<-pongs)



}

//只写通道,只写的定义方式为`chan <- string`
func pingWrite(pings chan <- string,msg string) {
	pings <- msg
}

//只读定义 <-chan string
//只写的pongs,使用`chan<-	string`来定义
func pongRead(pings <-chan string,pongs chan <- string){
	//读出pings
	msg := <- pings //从pings通道读取值给msg
	pongs <- msg //将读取的pings写入pongs中
}

