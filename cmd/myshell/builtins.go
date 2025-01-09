package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
)

func exit(args []string) {
	os.Exit(0)
}

func pwd(args []string) {
	fmt.Println(os.Getenv("PWD"))
}

func base_echo(args []string) {
	stringg := strings.Join(args, " ")
	fmt.Print(stringg[1:len(stringg)-1], "\n")
}

func adv_echo(args []string) {
	fmt.Print(strings.Join(args, " "), "\n")
}

func echo(args []string) {
	start := strings.HasPrefix(args[0], "'")
	end := strings.HasSuffix(args[len(args)-1], "'")

	if start && end {
		base_echo(args)
		return
	}

	formatted_Args := slices.DeleteFunc(args, func(s string) bool {
		return s == ""
	})

	adv_echo(formatted_Args)
}

func run(command string, args []string) {
	_, found := searchCommandInPath(command)
	if !found {
		fmt.Println(command + ": command not found")
	} else {
		cmd := exec.Command(command, args...)
		stdout, _ := cmd.Output()

		fmt.Print(string(stdout))
	}
}

func ttype(args []string) {
	// search command in buildin or path nev
	command := args[0]

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

func cd(args []string) {
	tempPath := args[0]

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
