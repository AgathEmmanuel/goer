package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	var ss string = "Helllooo\n"
	fmt.Println(ss)
        for i := 0; i < len(ss); i++ {
		fmt.Printf("Character at %d Index Position = %c\n", i, ss[i])
	}
	fmt.Printf("%c\n", ss[1])
	fmt.Printf("%T\n",ss[1])
}
