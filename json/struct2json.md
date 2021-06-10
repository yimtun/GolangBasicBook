```go
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

//写一个 testStruct()结构体的序列化方法
func testStruct() {
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2019-10-27",
		Sal:      10000.01,
		Skill:    "牛魔拳",
		Names:    []string{"xxx", "ddd"},
	}

	// 将Monster结构体序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误 err is %v", err)
	}
	//输出序列化结果
	fmt.Printf("monster序列化后 = %v", string(data))
}
func main() {
	testStruct()
}
```
