package main


import (
	"fmt"
	"container/list"
)


func main() {
	var il list.List
	il.PushBack(10)
	il.PushBack(11)
	il.PushBack(12)
	il.PushBack(13)
	il.PushBack(14)


	fmt.Println(il)
	fmt.Println(il.Front())
	fmt.Println(il.Front().Value)
	fmt.Println(il.Front().Value.(int))

	fmt.Println()
	for e:=il.Front();e!=nil;e=e.Next(){
		fmt.Println(e.Value.(int))
	}
	fmt.Println()


}
