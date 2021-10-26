package main

import (
	"fmt"
)

const constVar1 int = 64

func main() {

	var stringVar1 string = "Hellloooo"
	fmt.Println(stringVar1)
	fmt.Printf("variable's type is %T\n", stringVar1)

	var integerVar1 int = 42
	fmt.Println(integerVar1)

	var integerDefauldValue int
	fmt.Println(integerDefauldValue)

	var stringVar2 = "Hii.....iiiiii..."
	fmt.Println(stringVar2)
	fmt.Printf("variable's type is %T\n", stringVar2)

	var integerVar2 = 53
	fmt.Println(integerVar2)
	fmt.Printf("variable's type is %T\n", integerVar2)

	stringVar3 := "This is also a string"
	fmt.Println(stringVar3)
	fmt.Printf("variable's type is %T\n", stringVar3)

	fmt.Println(constVar1)
	fmt.Printf("variable's type is %T\n", constVar1)

}
