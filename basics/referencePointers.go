package main

import (
	"fmt"
)

func main() {
	intVar1 := 42
	var pointerVar = &intVar1
	fmt.Println("Value of intvar:", intVar1)
	fmt.Println("Value of pointer:", pointerVar)
	fmt.Println("Value inside pointer:", *pointerVar)

	fmt.Println("when you change the value through the pointer, you are also changing the value")
	*pointerVar = *pointerVar / 2
	fmt.Println("Value inside pointer:", *pointerVar)
	fmt.Println("Value of intvar:", intVar1)
}
