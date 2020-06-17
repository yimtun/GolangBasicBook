```
package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
	Names    []string
}

// map序列化
func testMap() {
	//定义一个map
	var a map[string]interface{}
	//使用map之前 必须make一下
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 18
	a["address"] = "重庆洪崖洞"
	a["struct"] = Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2019-10-27",
		Sal:      10000.01,
		Skill:    "牛魔拳",
		Names:    []string{"xxx", "ddd"},
	}
	// 将a map结构体序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err is %v", err)
	}
	//输出序列化结果
	fmt.Printf("map序列化后 = %v", string(data))
}

func main() {
	//结构体， map, slice进行序列化
	testMap()
}
```
