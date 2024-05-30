package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.Trim(input, "\r\n")
	fmt.Fprintf(os.Stdout, "%s: command not found\n", input)
}
