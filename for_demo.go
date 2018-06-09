//for三种写法
package main

import "fmt"

func main()  {

	//相当于while
	i := 1
	for i < 10{
		fmt.Println(i)
		i++
	}

	//初始化条件
	for j := 1;j<10;j++{
		fmt.Printf("当前j是%d\n",j)
	}

	//无限循环,除非有break,return结束
	j :=1;
	for{
		if j >100{
			break
		}

		fmt.Println("当前j是",j)
		j++
	}


}