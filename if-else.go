package main

import "fmt"

func main()  {
	if num := 10;num < 10{
		fmt.Println("num <10")
	} else if num < 100 {
		fmt.Println("num >=10 and num < 100")
	}else{
		fmt.Println("num > 100")
	}
}
