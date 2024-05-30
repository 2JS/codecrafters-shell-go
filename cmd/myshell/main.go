package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.Trim(input, "\r\n")
		args := strings.Split(input, " ")
		command, args := args[0], args[1:]
		switch command {
		case "echo":
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(args, " "))
		case "exit":
			if args[0] == "0" {
				os.Exit(0)
			}
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", input)
		}
	}
}
