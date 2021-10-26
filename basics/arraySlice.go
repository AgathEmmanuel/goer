package main

import (
	"fmt"
	"sort"
)

func main() {
	var namsArray [4]string
	namsArray[0] = "name0"
	namsArray[1] = "name1"
	namsArray[2] = "name2"
	namsArray[3] = "name3"
	fmt.Println(namsArray)
	fmt.Println(namsArray[0])

	var numbersArray = [4]int{11, 22, 33, 44}
	fmt.Println(numbersArray)
	fmt.Println(len(numbersArray))
	fmt.Println("Array is an object, when passed to a function \n a copy will be made")

	var namesSlice = []string{"names0", "names1", "names2"}
	fmt.Println(namesSlice)
	namesSlice = append(namesSlice, "names3")
	fmt.Println(namesSlice)
	fmt.Println(len(namesSlice))

	namesSlice = append(namesSlice[2 : len(namesSlice)+0])
	fmt.Println(namesSlice)
	namesSlice = append(namesSlice[:len(namesSlice)-1])
	fmt.Println(namesSlice)

	// namesSliceMake := make([]string, 3 (initialsize), 3 (capsize))
	namesSliceMake := make([]string, 3, 3)
	namesSliceMake[0] = "zero"
	namesSliceMake[1] = "one"
	namesSliceMake[2] = "three"
	fmt.Println(namesSliceMake)

	namesSliceMakeNocap := make([]string, 3)
	namesSliceMakeNocap[0] = "zero"
	namesSliceMakeNocap[1] = "one"
	namesSliceMakeNocap[2] = "three"
	fmt.Println(namesSliceMakeNocap)
	namesSliceMakeNocap = append(namesSliceMakeNocap, "four")
	fmt.Println(namesSliceMakeNocap)

	sort.Strings(namesSliceMakeNocap)
	fmt.Println(namesSliceMakeNocap)

}
