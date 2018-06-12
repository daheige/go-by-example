package main

import (
	"os"
	"fmt"
)

func main()  {
	params := os.Args

	fmt.Println("第一个参数是",params[0])
	//fmt.Println("第二个参数是",params[1])
}
