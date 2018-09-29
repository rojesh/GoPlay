package main

import (
	"fmt"
	"github.com/user/goplay/stringutil"
	"os"
	"bufio"
	"strings"
)

const inputdelimiter = '\n'

func main() {
	fmt.Println("Enter the sentence: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString(inputdelimiter)
	input = strings.Replace(input, "\n", "", -1)

	fmt.Println(stringutil.Abbreviate(input))
}