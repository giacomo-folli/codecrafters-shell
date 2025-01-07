package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		builtins := []string{"echo", "exit", "type"}

		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		userInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stdout, "Error in reading string")
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
			found := false
			for i := 0; i < len(builtins); i++ {
				if builtins[i] == args[0] {
					fmt.Println(args[0], "is a shell builtin")
					found = true
					break
				}
			}

			if !found {
				fmt.Println("invalid_command: not found")
			}
		} else {
			fmt.Println(command + ": command not found")
		}
	}
}
