package main

import (
	"fmt"

	"github.com/valyala/fastjson"
)

func main() {
	var json fastjson.Parser
	jsonData := `{"name": "John", "age": 30, "city": "New York", "arr": [1, 2, 3, 4, 5]}`
	value, err :=json.Parse(jsonData)

	if err != nil {
		panic(err)
	}

	fmt.Println(value)
	fmt.Println(string(value.GetStringBytes("name")))
	fmt.Println(int(value.GetInt("age")))
	fmt.Println(value.GetArray("arr"))
}