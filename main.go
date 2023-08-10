package main

import (
	"fmt"
)

func main() {
	var arr [3]string
	arr[0] = "1"
	arr[1] = "2"
	arr[2] = "3"
	fmt.Println(arr)

	// var m map[string]string
	m := make(map[string]string)
	m["key1"] = "value1"
	m["key2"] = "value2"
	m["key3"] = "value3"
	fmt.Println(m)

	for x := 0; x < 10; x++ {
		fmt.Println("Hello GO")
	}
}
