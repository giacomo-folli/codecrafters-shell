package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func searchCommandInDir(command string) (string, bool) {
	path, err := exec.LookPath(command)
	if err == nil {
		return path, true
	}

	return "", false
}

func main() {
	for {
		builtins := []string{"echo", "exit", "type"}

		fmt.Fprint(os.Stdout, "$ ")

		userInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stdout, "Error in reading string\n")
			os.Exit(1)
		}

		input_string := strings.TrimRight(userInput, "\r\n")
		parsed_string := strings.Split(input_string, " ")

		command := parsed_string[0]
		args := parsed_string[1:]

		if command == "exit" {
			os.Exit(0)
		} else if command == "echo" {
			fmt.Println(strings.Join(args, " "))
		} else if command == "type" {
			isBuildIn := false
			for i := 0; i < len(builtins); i++ {
				if builtins[i] == args[0] {
					fmt.Println(args[0], "is a shell builtin")
					isBuildIn = true
					break
				}
			}

			if isBuildIn {
				continue
			}

			path, found := searchCommandInDir(args[0])
			if found {
				fmt.Println(args[0], "is", path)
			} else {
				fmt.Printf("%s: not found\n", args[0])
			}

		} else {
			fmt.Println(command + ": command not found")
		}
	}
}
