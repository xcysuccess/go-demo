package main

import (
	"encoding/json"
	"fmt"
)

// Message TODO
type Message struct {
	Name string `json:"some_field"`
	Body string `json:"-"`
	Time int64
}

func main() {
	m := Message{"Alice", "Hello", 234234234}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}
	var a string = string(b)
	fmt.Printf("%s", a)
}
