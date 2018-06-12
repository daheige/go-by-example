package main

import (
	"io/ioutil"
	"os"
	"bufio"
	"fmt"
)

func main()  {
	data := []byte("大黑哥313333")
	//1. 将内容写入文件中
	err := ioutil.WriteFile("tmp.log",data,0644)
	check(err)

	//2. 以更小的粒度将内容写入文件中
	f,err := os.Create("tmp_2.log")
	check(err)
	defer f.Close()

	f.Write([]byte("2333\n")) //以字节切片的方式写入
	f.WriteString("daheige") //以字符串的方式写入
	f.Sync() //调用Sync方法来将缓冲区数据写入磁盘

	//3. 通过`bufio`缓存的方式写入
	w := bufio.NewWriter(f)
	n1,err := f.WriteString("哈哈")
	fmt.Println("写入的字节num ",n1)
	w.Flush() //将缓存区中的内容刷到文件中

}

func check(e error)  {
	if e != nil{
		panic(e)
	}
}
