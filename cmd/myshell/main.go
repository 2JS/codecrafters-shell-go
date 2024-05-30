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
			echo(args)
		case "exit":
			exit(args)
		default:
			commandNotFound(command)
		}
	}
}

func commandNotFound(command string) {
	fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
}

func exit(args []string) {
	status := args[0]
	if status == "0" {
		os.Exit(0)
	}
}

func echo(args []string) {
	fmt.Fprintf(os.Stdout, "%s\n", strings.Join(args, " "))
}
