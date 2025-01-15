package main

import (
	"fmt"
	"os"
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

func task(command string, args []string) (ok bool) {
	// override args if found redirection action
	found, args, file := _checkRedirection(args)
	output := "\n"
	var err error

	handler, ok := commands[command]
	if ok {
		output = handler(args)
	} else {
		output, err = run(command, args)
	}

	if !found {
		fmt.Print(output)
	} else {
		if err != nil {
			fmt.Print(err.Error())
		}

		if output != "" {
			err = _writeToFile(file[0], output)
			if err != nil {
				fmt.Print("could not write in file\n")
			}
		}
	}
	return
}

func main() {
	_initEnv()

	for {
		fmt.Fprint(os.Stdout, "$ ")

		userInput := _getUserIput()
		if userInput == "" {
			continue
		}

		args := _parseArgs(userInput)

		task(args[0], args[1:])
	}
}
