```
package main

import (
	"encoding/json"
	"fmt"
)

// map序列化
func testMap() {
	//定义一个map
	var a map[string]string
	//使用map之前 必须make一下
	a = make(map[string]string)
	a["name"] = "红孩儿"
	a["age"] = "18"
	a["address"] = "重庆洪崖洞"
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
