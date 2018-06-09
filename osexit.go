package main

import (
	"fmt"
	"os"
)

//Go和C语言不同,main函数并不返回一个整数来表示程序的退出
//状态,而是将退出状态作为函数的参数

func main()  {
	defer fmt.Println("fefe") //当使用`os.Exit`的时候defer操作不会被运行,

	os.Exit(3)
}