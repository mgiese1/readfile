// readfile.go
//https://www.golangprograms.com/go-program-to-reading-plain-text-files.html
// reading text line-by-line from the plain text file
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("pepe.txt") //open text file in read-only mode and returns a pointer of type os.File

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for _, eachline := range txtlines { //_ is the blank identifier https://golang.org/doc/effective_go.html#blank
		fmt.Println(eachline)
	}
}
