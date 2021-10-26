package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Maps")

	keyValue := make(map[string]string)
	fmt.Println(keyValue)
	keyValue["key1"] = "value1"
	keyValue["key2"] = "value2"
	keyValue["key3"] = "value3"
	keyValue["key4"] = "value4"
	fmt.Println(keyValue)
	fmt.Println(keyValue["key2"])
	delete(keyValue, "key1")
	fmt.Println(keyValue)

	// In Maps there is no specifed iterational order
	// The order of iteration can be manually managed
	for k, v := range keyValue {
		fmt.Printf("%v : %v \n", k, v)
	}
	keys := make([]string, len(keyValue))
	i := 0
	for k := range keyValue {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println(keys)

	for i := range keys {
		fmt.Println(keyValue[keys[i]])
	}
}
