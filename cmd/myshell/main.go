package main

import (
	"fmt"
	"os"
	"strings"
)

var builtins = []string{"echo", "exit", "type", "pwd"}
var commands = map[string]MyFunc{
	"echo": echo,
	"exit": exit,
	"type": ttype,
	"pwd":  pwd,
	"cd":   cd,
}

func main() {
	os.Setenv("PWD", initPwdVar())
	os.Setenv("HOME", initHomeVar())

	for {
		fmt.Fprint(os.Stdout, "$ ")

		userInput := getUserIput()
		input_string := strings.TrimRight(userInput, "\r\n")
		parsed_string := strings.Split(input_string, " ")

		command := parsed_string[0]
		args := parsed_string[1:]

		task(command, args)
	}
}

// ---------------------------------------------------------
// TODOTODOTODOTODOTODOTODOTODOTODOTODOTODOTODOTODOTODOTODOT
// ---------------------------------------------------------

type MyFunc func(args []string)

func task(command string, args []string) (ok bool) {
	handler, ok := commands[command]
	if ok {
		handler(args)
	} else {
		run(command, args)
	}

	return

}
