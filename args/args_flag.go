//flag包来支持基本的命令行标记解析
package main

import (
	"flag"
	"fmt"
)

func main()  {
	wordptr := flag.String("word","word","please input word") //是string类型的指针
	numptr := flag.Int("d",1,"please input num")

	var name string
	flag.StringVar(&name,"name","daheige","please input your name")
	flag.Parse() //使用`flag.Parse()`来分解命令行选项

	fmt.Println(wordptr) //0xc42000e1f0
	fmt.Println(*wordptr) //word

	fmt.Println("您输入的数字是",*numptr)
	fmt.Println("您输入的名字是",name)
}

/**
go run args_flag.go -word=123322 -d=34
0xc42004c1d0
123322
您输入的数字是 34
您输入的名字是 daheige313

//查看帮助
go run args_flag.go -h
 */