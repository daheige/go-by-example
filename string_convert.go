//String转换到Byte数组时,每个byte(byte类型其实就是uint8)保
//存字符串对应字节的数值
package main

import "fmt"

func main()  {
	//unicode,每个中文
	//字符会由三个byte组成
	str := "大黑哥313"
	fmt.Println([]byte(str))
	fmt.Println(len([]byte(str)))

	//rune一个字一个数据 4个字节
	fmt.Println([]rune(str))
	fmt.Println(len([]rune(str))) //6个长度

}