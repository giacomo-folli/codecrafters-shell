package main

import (
	"bufio"
	"fmt"
	"os"
)

var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error in reading string")
		os.Exit(1)
	}

	fmt.Println(command[:len(command)-1], ": command not found ")
	// _, err = exec.LookPath(userInput)
	//if err != nil {
	//	fmt.Printf("%s: Command not found", userInput)
	//}
}
