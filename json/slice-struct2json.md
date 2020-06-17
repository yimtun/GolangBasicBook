```
package main

import (
	"encoding/json"
	"fmt"
)

type App struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
	Names    []string
}

// map序列化
func testMap() {
	var slice []App
	//定义一个map
	m1 := App{
		Name:  "sxs",
		Age:   1,
		Names: []string{"xxx", "bbb"}}
	slice = append(slice, m1)
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err is %v", err)
	}
	//输出序列化结果
	fmt.Println(string(data))
}

func main() {
	//结构体， map, slice进行序列化
	testMap()
}
```
