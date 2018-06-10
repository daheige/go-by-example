package main

import (
	urlTool "net/url"
	"fmt"
	"strings"
)

func main()  {
	url := "http://mygo.com:8080/path?k=123&name=daheige#tag"

	u,_ := urlTool.Parse(url) //解析url,可以获取到scheme,host,path,query
	fmt.Println("协议: ",u.Scheme)
	fmt.Println("host(包含端口): ",u.Host)
	fmt.Println("path: ",u.Path)
	fmt.Println("query str: ",u.RawQuery)

	//手动分解主机名和端口
	h := strings.Split(u.Host,":")
	fmt.Println(h)

	fmt.Println("host is: ",h[0])
	if len(h) >1 {
		fmt.Println("port is: ",h[1])
	}else{
		fmt.Println("port default is: ",80)
	}

	//将query解析到一个map中,key是字符串,value是一个字符串切片
	m,_:= urlTool.ParseQuery(u.RawQuery)
	fmt.Println(m) //map[k:[123] name:[daheige]]
	fmt.Println(m["k"][0])
	fmt.Println(m["name"][0])

	//快速获取query,和parseQuery的结果一样
	params := u.Query()
	fmt.Println("参数是: ",params)
	fmt.Println(params["k"][0])

	//获取查询片段
	fmt.Println(u.Fragment)

	enc :=urlTool.QueryEscape("a=1&b=2")
	fmt.Println(enc) //编码

	//解码
	query_str,err := urlTool.QueryUnescape(enc)
	if err != nil{
		panic(err)
	}

	fmt.Println(query_str)

	//将map的参数列表格式化为a=1&b=2形式
	q := urlTool.Values{}
	q.Add("a","111")
	q.Add("b","2")
	q.Set("c","3")
	fmt.Println(q)
	fmt.Println(q.Encode()) //a=111&b=2&c=3

}