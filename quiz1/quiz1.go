package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

const inputdelimiter = '\n'

var count int

func main() {
	fmt.Printf("LETS TAKE ON THE QUIZ\n")
	fmt.Printf("---------------------\n")
	fmt.Printf("---------------------\n")
	ReadCsvFile("quiz1.csv")
}

func ReadCsvFile(filePath string) {
	csvFile, _ := os.Open(filePath)
	csvReader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			fmt.Printf("You got %d answers correct\n", count)
			break
		}
		for index, value := range record {
			if index == 0 {
				fmt.Printf("What is %v\n", value)
			} else {
				reader := bufio.NewReader(os.Stdin)
				input, _ := reader.ReadString(inputdelimiter)
				input = strings.Replace(input, "\n", "", -1)
				if input == value {
					count++
				}
			}
		}
	}
}
