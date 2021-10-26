package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files")
	writeFile()
	defer readFile("./exWords.txt")
}

func writeFile() {
	words := "words need to be written to file"
	file, err := os.Create("./exWords.txt")
	showError(err)
	length, err := io.WriteString(file, words)
	showError(err)
	fmt.Println("The length of file is", length)
	defer file.Close()
}

func readFile(fileName string) {
	content, err := ioutil.ReadFile(fileName)
	showError(err)
	fmt.Println(string(content))
	fmt.Println("Completed Reading")

}

func showError(err error) {
	if err != nil {
		panic(err)
	}
}
