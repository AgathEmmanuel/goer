package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Pleas enter your number: ")
	numInput, _ := reader.ReadString('\n')
	floatVar1, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The number you enetered is:", floatVar1)
	}
	var intVar1 int = 9
	var floatVar2 float64 = 44.422
	sum := float64(intVar1) + floatVar2 + floatVar1
	roundedSum := math.Round(sum)
	fmt.Printf("The sum is: %v, The Type of the output sum is: %T\n", sum, sum)
	fmt.Printf("The roundedSum is: %v, The Type of the output roundedSum is: %T\n", roundedSum, roundedSum)

	floatVar3, floatVar4, floatVar5 := 1.11, 2.2, 3.4
	floatSum := floatVar3 + floatVar4 + floatVar5
	floatSum = math.Round(floatSum)
	fmt.Println("Float sum is: ", floatSum)

	radius := 1.1
	areaCircle := radius * math.Pi * math.Pi
	fmt.Printf("The area of Circle is: %.3f\n", areaCircle)
}
