package main

import (
	"fmt"
	"os"
)

type MyFunc func(args []string) string

var (
	builtins = []string{"echo", "exit", "type", "pwd"}
	commands = map[string]MyFunc{
		"echo": echo,
		"exit": exit,
		"type": ttype,
		"pwd":  pwd,
		"cd":   cd,
	}
)

func task(command string, args []string) (ok bool) {
	// override args if found redirection action
	standRedirect, errRedirect, appendIt, args, file := _checkRedirection(args)
	outString, errString := "\n", ""
	var err error

	handler, ok := commands[command]
	if ok {
		outString = handler(args)
	} else {
		outString, errString, err = run(command, args)
	}

	if errRedirect {
		err = _writeToFile(file[0], errString, appendIt)
		if err != nil {
			fmt.Print("could not write in file\n")
		}

		if outString != "" {
			fmt.Print(outString)
		}
		return
	}

	if standRedirect {
		if err != nil {
			if errString != "" {
				fmt.Print(errString)
			} else {
				fmt.Print(err.Error())
			}
		}

		if outString != "" {
			err = _writeToFile(file[0], outString, appendIt)
			if err != nil {
				fmt.Print("could not write in file\n")
			}
		}
		return
	}

	if err != nil {
		if errString != "" {
			fmt.Print(errString)
		} else {
			fmt.Print(err.Error())
		}
		return
	}

	fmt.Print(outString)
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
