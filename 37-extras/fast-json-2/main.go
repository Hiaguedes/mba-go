package main

import (
	"fmt"

	"github.com/valyala/fastjson"
)

func main() {
	json := `{"user": {"name": "John", "age": 30}, "city": "New York"}`
	var p fastjson.Parser
	v, err := p.Parse(json)
	if err != nil {
		panic(err)
	}

	user := v.Get("user")
	name := user.GetStringBytes("name")
	age := user.GetInt("age")
	fmt.Println(user)
	fmt.Println(string(name), user.Get("name")) // Output: John
	fmt.Println(age)          // Output: 30
}