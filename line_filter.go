package main

import (
    "bufio"
    "os"
    "strings"
    "fmt"
)

func main() {
    //行数据过滤器
    scanner := bufio.NewScanner(os.Stdin) //使用缓冲scanner来包裹无缓冲的`os.Stdin`
    for scanner.Scan(){ //使用`Scan`方法,这个方法会将scanner定位到下一行
        if scanner.Text() == "q"{ //`Text`方法从输入中返回当前行
            fmt.Println("程序即将退出")
            break
        }

        ucl := strings.ToUpper(scanner.Text())
        //转换为大写
        fmt.Println(ucl)
    }

    if err := scanner.Err();err != nil{
        fmt.Fprintln(os.Stderr,"err",err)
        os.Exit(1)
    }
}
