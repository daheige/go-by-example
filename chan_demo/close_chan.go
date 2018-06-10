//关闭通道的意思是该通道将不再允许写入数据。这个方法可以让通道数
//据的接受端知道数据已经全部发送完成了

package main

import "fmt"

func main(){
	jobs := make(chan int,5)
	done := make(chan bool,1)

	//数据接收端协程,检查通道是否关闭
	go func(){
		for {
			j,ok := <-jobs
			if ok{
				//fmt.Println("通道未关闭")
				fmt.Println("接收通道",j)
			}else{
				fmt.Println("数据全部发送完毕")
				done <- true
				break
			}
		}
	}()

	//开始发送数据
	for i := 0;i<3;i++{
		fmt.Println("开始发送",i)
		jobs <- i
	}

	close(jobs)
	fmt.Println("开始发送所有的jobs")
	//一直会等所有的数据发送完毕
	<-done //one通道在接收数据的时候会阻塞,所以在所有的数据发送接收完成后,写入done的数据将在这里被接收,然后程序结束
}
