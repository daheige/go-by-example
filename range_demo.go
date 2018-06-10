package main

import "fmt"

func main()  {
	arr := []int{1,2,3}
	for k,v := range arr{
		fmt.Println(k,v)
	}

	for _,v := range arr{
		fmt.Println("当前v是",v)
	}
}