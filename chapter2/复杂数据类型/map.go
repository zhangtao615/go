package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var m1 map[int]string
	m1 = make(map[int]string)
	m1[1] = "Tom"
	fmt.Println("m1:", m1) // m1: map[1:Tom]

	var m2 map[int]string = map[int]string{}
	m2[1] = "Tom"
	fmt.Println("m2:", m2) // m2: map[1:Tom]
	toJson()
}

func toJson() {
	res := make(map[string]interface{})
	res["code"] = 200
	res["msg"] = "success"
	res["data"] = map[string]interface{}{
		"username": "Tom",
		"age":      "20",
		"hobby":    []string{"看电影", "读书"},
	}

	// 序列化
	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshal error:\n", errs)
	}
	fmt.Println("json data: \n", string(jsons))

	// 反序列化
	res2 := make(map[string]interface{})
	errs = json.Unmarshal([]byte(jsons), &res2)
	if errs != nil {
		fmt.Println("json marshal error:\n", errs)
	}
	fmt.Println("map data :\n", res2)
}
