package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func exit(args string) {
	os.Exit(0)
}

func pwd(args string) {
	fmt.Println(os.Getenv("PWD"))
}

func echo(args string) {
	fmt.Println(args)
}

// search command in buildin or path nev
func ttype(args string) {
	command := strings.Split(args, " ")[0]

	isBuildIn := searchBuildin(command)
	if isBuildIn {
		fmt.Println(command, "is a shell builtin")
	} else {
		path, found := searchCommandInPath(command)
		if found {
			fmt.Println(command, "is", path)
		} else {
			fmt.Printf("%s: not found\n", command)
		}
	}
}

func cd(args string) {
	tempPath := strings.Split(args, " ")[0]

	isHome := tempPath[0] == '~'
	if isHome {
		os.Setenv("PWD", os.Getenv("HOME"))
		return
	}

	isAbsolute := tempPath[0] == '/'
	if !isAbsolute {
		tempPath = filepath.Join(os.Getenv("PWD"), tempPath)
	}

	if _, err := os.Stat(tempPath); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("%s: No such file or directory\n", tempPath)
		return
	}

	os.Setenv("PWD", tempPath)
}

// Run a general command provided by the user
func run(command string, args string) {
	// args_slice := ParseArgs(args)

	_, found := searchCommandInPath(command)
	if !found {
		fmt.Println(command + ": command not found")
	} else {
		cmd := exec.Command(command, args)
		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println("Error in execution of", command, "command")
		}

		fmt.Print(string(stdout))
	}
}
