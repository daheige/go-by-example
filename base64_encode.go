package main

import (
	"encoding/base64"
	"fmt"
)

func main()  {
	//Go支持标准的和兼容URL的base64编码
	//标准的编码后面是+
	data := "abced@#$!"
	sEnc := base64.StdEncoding.EncodeToString([]byte(data)) //编码为64
	fmt.Println(sEnc)

	//解码成base64原来的额字符串
	str ,_:= base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(str))

	//兼容url,使用兼容URL的base64编码和解码
	s := "sfefefe###!sss"
	str_enc := base64.URLEncoding.EncodeToString([]byte(s))
	fmt.Println("编码之后",str_enc) //c2ZlZmVmZSMjIyFzc3M=

	//解码经过url编码的字符串
	str_dec,err := base64.URLEncoding.DecodeString(str_enc)
	if err != nil{
		panic(err)
	}

	fmt.Printf("解码之后的s是%s",string(str_dec))

}