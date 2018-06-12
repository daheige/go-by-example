package main

import "fmt"

func main()  {
	message := make(chan string,1)
	sigs := make(chan bool)

	//第一次没有消息可以接收,就会走default逻辑
	select {
	case msg :=<-message:
		fmt.Println("received message : ",msg)
	default:
		fmt.Println("no message!")
	}

	//开始发送消息
	msg := "hi dahei"
	select{
	case message <- msg:
		fmt.Println("send message:",msg)
	default:
		fmt.Println("no sent message")
	}

	//接收信息
	select {
	case msg:= <-message:
		fmt.Println("current received message: ",msg)
	case sig :=<-sigs:
		fmt.Println("received sig:",sig)
	default:
		fmt.Println("no activity")
	}
}

/**
no message!
send message: hi dahei
current received message:  hi dahei
 */