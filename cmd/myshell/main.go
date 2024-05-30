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
		case "type":
			_type(args)
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

func _type(args []string) {
	command := args[0]
	switch command {
	case "echo", "exit", "type":
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", command)
	default:
		path, err := searchPath(command)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Fprintf(os.Stdout, "%s is %s/%s\n", command, path, command)
		}
	}
}

func searchPath(command string) (string, error) {
	pathVar := os.Getenv("PATH")
	paths := strings.Split(pathVar, ":")
	for _, path := range paths {
		files, _ := os.ReadDir(path)

		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if file.Name() == command {
				return path, nil
			}
		}
	}
	return "", fmt.Errorf("%s: not found", command)
}
