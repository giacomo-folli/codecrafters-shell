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

		userInput := getUserIput()
		input_string := strings.TrimRight(userInput, "\r\n")
		parsed_string := strings.Split(input_string, " ")

		command := parsed_string[0]
		args := parsed_string[1:]

		switch command {
		case "exit":
			// terminate session
			os.Exit(0)

		case "echo":
			// echo command (buildin)
			fmt.Println(strings.Join(args, " "))

		case "type":
			// search command in buildin or path nev
			isBuildIn := searchBuildin(args[0])
			if isBuildIn {
				fmt.Println(args[0], "is a shell builtin")
			} else {
				path, found := searchCommandInPath(args[0])
				if found {
					fmt.Println(args[0], "is", path)
				} else {
					fmt.Printf("%s: not found\n", args[0])
				}
			}

		default:
			_, found := searchCommandInPath(command)
			if !found {
				fmt.Println(command + ": command not found")
			} else {
				cmd := exec.Command(command, args...)
				stdout, _ := cmd.Output()

				fmt.Println(string(stdout))
			}
		}
	}
}

func searchCommandInPath(command string) (string, bool) {
	path, err := exec.LookPath(command)
	if err == nil {
		return path, true
	}

	return "", false
}

func searchBuildin(command string) bool {
	builtins := []string{"echo", "exit", "type"}
	found := false

	for i := 0; i < len(builtins); i++ {
		if builtins[i] == command {
			found = true
			break
		}
	}

	return found
}

func getUserIput() string {
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error in reading string\n")
		os.Exit(1)
	}

	return input
}
