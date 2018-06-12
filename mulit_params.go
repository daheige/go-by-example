package main

import "fmt"

func main()  {
	fmt.Println(sum(1,3,4))

	//通过展开操作符...,将int类型的切片传入sum
	arr := []int{1,2,3,4}
	fmt.Println(sum(arr...))
}

//可变参数一般来说是函数最后一个参数
func sum(nums ...int) (total int) {
	for _,v := range nums{
		total += v
	}

	return
}