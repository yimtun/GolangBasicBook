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
}

// slice进行序列化
func testSlice() {
	var slice []map[string]interface{} // 定义了一个切片，里面是map格式 map[string]interface{}
	var m1 map[string]interface{}      //定义切片中的第一个map M1
	m1 = make(map[string]interface{})
	m1["name"] = "Matt"
	m1["age"] = 36
	m1["address"] = [2]string{"重庆沙坪坝", "上海"}
	slice = append(slice, m1)

	var m2 map[string]interface{} //定义切片中的第2个map M2
	m2 = make(map[string]interface{})
	m2["name"] = "damon"
	m2["age"] = 36
	m2["address"] = "沙坪坝"
	slice = append(slice, m2)

	// 将slice进行序列化
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err is %v", err)
	}
	//输出序列化结果
	fmt.Println(string(data))
}

func main() {
	//结构体， map, slice进行序列化
	testSlice()
}
```
