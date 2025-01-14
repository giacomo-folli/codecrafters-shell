package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type MyFunc func(args []string) string

var builtins = []string{"echo", "exit", "type", "pwd"}
var commands = map[string]MyFunc{
	"echo": echo,
	"exit": exit,
	"type": ttype,
	"pwd":  pwd,
	"cd":   cd,
}

func main() {
	err := godotenv.Load()
	if err == nil {
		environment := os.Getenv("ENV")
		fmt.Println("APP RUNNING IN", environment, "MODE")
	}

	os.Setenv("PWD", _initPwdVar())
	os.Setenv("HOME", _initHomeVar())

	for {
		fmt.Fprint(os.Stdout, "$ ")

		userInput := _getUserIput()
		parsedInput := _parseArgs(strings.TrimRight(userInput, "\r\n"))

		command := parsedInput[0]
		args := parsedInput[1:]

		task(command, args)
	}
}

func task(command string, args []string) (ok bool) {
	handler, ok := commands[command]
	output := "\n"

	var err error
	found, args, file := _checkRedirection(args)

	if ok {
		output = handler(args)
	} else {
		output, err = run(command, args)
	}

	if found && err != nil {
		_writeToFile(file[0], output)
		return
	}

	fmt.Print(output)
	return
}
