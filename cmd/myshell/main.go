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

		// Wait for user input
		userInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stdout, "Error in reading string")
			os.Exit(1)
		}

		command := strings.TrimRight(userInput, "\r\n")

		if command == "exit 0" {
			os.Exit(0)
		} else {
			// Print the command not found message
			fmt.Println(command + ": command not found")
		}
	}
}
