package main

import "fmt"

func main()  {
	randomMap := make(map[string]string)

	randomMap["key"] = "value"

	_, ok := randomMap["notExist"]
	val, exist := randomMap["key"]

	fmt.Println(ok, exist, val)
}