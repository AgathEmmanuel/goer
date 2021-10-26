package main

import (
	"fmt"
)

func main() {
	fmt.Println("structs")
	bugatti := Car{"red", 1000, 10}
	fmt.Println(bugatti)
	fmt.Printf("%+v\n", bugatti)
	fmt.Printf("%v\n", bugatti.Colour)
	fmt.Printf("%v\n", bugatti.Cost)
	bugatti.Colour = "Blue"
	fmt.Printf("%+v\n", bugatti)
	fmt.Printf("%v\n", bugatti.Colour)
}

// defining a struct
type Car struct {
	Colour string
	Cost   int
	Weight int
}
