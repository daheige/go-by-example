package main

import "fmt"

func main()  {
	mess := make(chan string,2)
	mess <- "111"
	mess <- "大黑哥"

	//开始接收数据(一次性从缓存通道中获取数据)
	fmt.Println(<-mess)
	fmt.Println(<-mess)
}