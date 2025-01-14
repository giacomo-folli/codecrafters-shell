package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type MyFunc func(args string)

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

	os.Setenv("PWD", initPwdVar())
	os.Setenv("HOME", initHomeVar())

	for {
		fmt.Fprint(os.Stdout, "$ ")

		userInput := getUserIput()
		parsedInput := ParseArgs(strings.TrimRight(userInput, "\r\n"))

		command := parsedInput[0]
		args := strings.Join(parsedInput[1:], " ")

		task(command, args)
	}
}

func task(command string, args string) (ok bool) {
	handler, ok := commands[command]

	if ok {
		handler(args)
	} else {
		run(command, args)
	}

	return
}
