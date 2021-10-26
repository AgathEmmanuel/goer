package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Pleas enter your text input: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("The value you entered is ", input)

}
