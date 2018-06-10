package main

import (
	"crypto/sha1"
	"fmt"
)

func main()  {
	s := "daheige313 hello"
	//生成一个hash sha1模式
	h := sha1.New()

	//写入hash字节,如果你的参数是字符串,使用`[]byte(s)`
	h.Write([]byte(s))

	bs := h.Sum(nil) //字节数组

	//SHA1散列值经常以16进制的方式输出
	fmt.Println(bs)
	hash_str := fmt.Sprintf("%x",bs) //格式化为16进制
	fmt.Println(hash_str)

	fmt.Printf("%x\n",bs)
}
