//字符串操作相关函数
package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println
func main()  {
	p("contains",strings.Contains("test","es")) //包含字符串
	p("count",strings.Count("daheige 123 daheige","dahei")) //统计子字符串出现的次数
	p("hasprefix",strings.HasPrefix("daheige","da")) //是否以某字符串开始
	p("index",strings.Index("daheige 313 da","heige")) //子字符串出现的位置
	p("join",strings.Join([]string{"a","b"},"-"))
	p("Repeat",strings.Repeat("abc",3))
	p("Replace",strings.Replace("daheige","hei","haha",-1))
	p("Replace",strings.Replace("daheige daheige","hei","haha",1)) //替换字符串,最后一个参数是退换的次数
	p("Split str",strings.Split("a-b-c-d","-"))
	p("ToLower",strings.ToLower("Abc"))
	p("ToUpper",strings.ToUpper("Abc"))
	p("len",len("fefe"))
}
