package main

import (
	"fmt"
)

func main() {
	testing()
	fmt.Println(addInt(1, 2))
	fmt.Println(addAllInt(1, 2, 9, 8))
	fmt.Println(addAllIntss(1, 2, 9, 8))

	hdetails := House{"myhouse", 10, 20, "fullname"}
	fmt.Println(hdetails)
	fmt.Printf("%+v\n", hdetails)
	fmt.Printf("%v\n", hdetails.Name)

	hdetails.Output()
	hdetails.AllOutput()

}

func testing() {
	fmt.Println("just testingg..")
}
func addInt(integer1 int, integer2 int) int {
	return integer1 + integer2
}

func addAllInt(integers ...int) int {
	sum := 0
	for _, integer := range integers {
		fmt.Println("--", integer)
		sum += integer

	}
	return sum
}

func addAllIntss(integers ...int) (int, int) {
	sum := 0
	for _, integer := range integers {
		fmt.Println("--", integer)
		sum += integer
	}
	return len(integers), sum
}

// example for struct
type House struct {
	Name   string
	Height int
	Weight int
	str    string
}

// method for type House
func (h House) Output() {
	fmt.Println(h.Height)
}

// method to print all
func (h House) AllOutput() {
	h.str = fmt.Sprintf("%v %v %v", h.Name, h.Height, h.Weight)
	fmt.Println(h.str)
}
