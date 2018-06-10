package main

import (
	"fmt"
	"time"
)

func main()  {
	//当做单一值判断
	i := 12
	switch i {
	case 1:
		fmt.Println("i = 1")
	case 12:
		fmt.Println("i = 12")
	default:
		fmt.Println("i is not 12")
	}

	//初始化表达式
	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
		fmt.Println("it is the weekend")
	default :
		fmt.Println("it is not the weekend")
	}

	//相对于if--else
	t := time.Now()
	switch{
	case t.Hour() < 12:
		fmt.Println("中午之前")
	default:
		fmt.Println("中午之后了")
	}


}