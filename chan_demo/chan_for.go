package main

import "fmt"

func main()  {
	queue := make(chan string,2)
	queue<- "daheige"
	queue<- "golang"
	close(queue) //关闭通道,停止接收数据

	//如果不关闭通道for...range是无法获取数据的
	//遍历通道中的值
	for elem := range queue{
		fmt.Println(elem)
	}

}