package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Get("url")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	responseValue, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	responseValueStr := string(responseValue)
	fmt.Print(responseValueStr)

	decoded := filterJson(responseValueStr)
	for _, i := range decoded {
		fmt.Println(i.key1)
	}

}

func filterJson(responseValueStr string) []DecodedSlice {
	decoded := make([]DecodedSlice, 0)
	decoder := json.NewDecoder(strings.NewReader(responseValueStr))
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}
	var decodedOut DecodedSlice
	for decoder.More() {
		err := decoder.Decode(&decoded)
		if err != nil {
			panic(err)
		}
		decoded = append(decoded, decodedOut)
	}
	return decoded
}

type DecodedSlice struct {
	key1, key2, key3 string
}
