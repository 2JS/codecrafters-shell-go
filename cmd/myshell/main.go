package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
		case "cd":
			cd(args)
		case "echo":
			echo(args)
		case "exit":
			exit(args)
		case "type":
			_type(args)
		default:
			path, err := exec.LookPath(command)
			if err != nil {
				commandNotFound(command)
				continue
			}
			cmd := exec.Command(path, args...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
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
		path, err := exec.LookPath(command)
		if err != nil {
			fmt.Fprintf(os.Stdout, "%s: not found\n", command)
		} else {
			fmt.Fprintf(os.Stdout, "%s is %s\n", command, path)
		}
	}
}

func cd(args []string) {
	path := args[0]
	if strings.HasPrefix(path, "~") {
		path = strings.Replace(path, "~", os.Getenv("HOME"), 1)
	}
	if err := os.Chdir(path); err != nil {
		fmt.Fprintf(os.Stdout, "%s: No such file or directory\n", path)
	}
}
