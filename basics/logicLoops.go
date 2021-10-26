package main

import (
	"fmt"
)

func main() {
	a := 1
	if a < 0 {
		fmt.Println("less than 0")
	} else if a == 0 {
		fmt.Println("equal to 0")
	} else {
		fmt.Println("greater than 0")
	}
	if b := 0; b < 0 {
		fmt.Println("less than 0")
	} else if b == 0 {
		fmt.Println("equal to 0")
	} else {
		fmt.Println("greater than 0")
	}

	switch testNum := 2; testNum {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("default")
	}

	testNum1 := 3
	switch testNum1 {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("two")
		fallthrough
	case 3:
		fmt.Println("three")
		fallthrough
	default:
		fmt.Println("default")
	}

	names := []string{"name1", "name2", "name3"}
	fmt.Println(names)

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	fmt.Println()
	for i := range names {
		fmt.Println(names[i])
	}

	fmt.Println()
	for _, name := range names {
		fmt.Println(name)
	}

	fmt.Println()
	for i, name := range names {
		fmt.Println(i, name)
	}

	fmt.Println()
	for i := 0; i < len(names); {
		fmt.Println(names[i])
		i++
	}

	var aa int = 10
	for aa < 20 {
		if aa == 15 {
			fmt.Println("fifteen")
		} else if aa > 17 {
			fmt.Println("seventeen")
			goto exitPoint
		} else if aa == 18 {
			fmt.Println("eighteen")
			break
		}
		fmt.Printf("value of aa: %d\n", aa)
		aa++
	}
exitPoint:
	fmt.Println("Exit Poin reached")
}
