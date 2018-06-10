package main

import (
	"encoding/json"
	"fmt"
	"os"
	"log"
)

type person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main()  {
	var stu = person{
		Name:"daheige",
		Age:27,
	}

	json_data,err := json.Marshal(stu)
	if err != nil{
		panic(err)
	}

	fmt.Println(string(json_data)) //{"name":"daheige","age":27}

	str :=`{"name":"daheige","age":27}`
	var p interface{}
	err = json.Unmarshal([]byte(str),&p)
	if err != nil{
		panic(err)
	}
	fmt.Println(p)
	m := p.(map[string]interface{}) //通过类型断言访问
	fmt.Println(m["name"])

	//将json数据写入os.writer中
	enc := json.NewEncoder(os.Stdout)
	data := map[string]int{
		"a":1,
		"b":2,
	}

	enc.Encode(data) //{"a":1,"b":2}

	//通过json编码器将内容写入文件中
	tmp, _ := os.OpenFile("tmp.json", os.O_CREATE|os.O_WRONLY | os.O_APPEND, 0644)
	defer tmp.Close()
	enc2 := json.NewEncoder(tmp)
	hg_err := enc2.Encode(data)
	if hg_err != nil {
		log.Println("Error in encoding json")
	}

	fmt.Println("写入成功")
}
