//defer在函数结束之前执行,一般用来释放资源
package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("defer使用")

	file := createFile("/tmp/log-defer.txt")
	defer closeFile(file)

	writeFile(file)

	fmt.Println("sucess")
}

func createFile(p string) *os.File{
	f,err := os.Create(p)
	if err != nil{
		panic(err)
	}

	return f
}

func writeFile(f *os.File){
	f.WriteString("大黑哥313")
}

func closeFile(f *os.File)  {
	f.Close()
}
