package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)


func main() {

	p:=flag.String("path","sample.txt","The path to the file to be filtered")
	v:=flag.String("value","broken","The vale to be filtered from the file")
	// flags can be used to pass in input arguments 
	// >> go run lineFilter.go -value broken
	// Here we are passing value available options can be [broken] [empty] [null]
	// >> go run lineFilter.go -value broken -path ../sample.txt
	// >> go run lineFilter.go -value broken -path sample.txt
	// Here we are passing value and path for the file location

	//file,err:=os.Open("sample.txt")

	flag.Parse()
	// provids command line utilities for example
	// >> go run lineFilter.go -help
	// provides with descriptions of what is expected that we specified above

	file,err:=os.Open(*p)
	if err!=nil {
		log.Fatal(err)
	}
	defer file.Close() // make sure everythings cleaned up after file variable usage and on close
	reader:=bufio.NewReader(file)
	for {
		string,err:=reader.ReadString('\n')
		if err!=nil {
			break
		}
		// if strings.Contains(string,"broken") {
		if strings.Contains(string,*v) {
			fmt.Println(string)
		}
	}
}