package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func exit(args []string) string {
	os.Exit(0)
	return "\n"
}

func pwd(args []string) string {
	return fmt.Sprintln(os.Getenv("PWD"))
}

func echo(args []string) string {
	return fmt.Sprintln(strings.Join(args, " "))
}

// search command in buildin or path nev
func ttype(args []string) string {
	command := args[0]

	isBuildIn := _searchBuildin(command)
	if isBuildIn {
		return fmt.Sprintln(command, "is a shell builtin")
	}

	path, found := _searchCommandInPath(command)
	if found {
		return fmt.Sprintln(command, "is", path)
	}

	return fmt.Sprintf("%s: not found\n", command)
}

func cd(args []string) string {
	tempPath := args[0]

	isHome := tempPath[0] == '~'
	if isHome {
		os.Setenv("PWD", os.Getenv("HOME"))
		return "\n"
	}

	isAbsolute := tempPath[0] == '/'
	if !isAbsolute {
		tempPath = filepath.Join(os.Getenv("PWD"), tempPath)
	}

	if _, err := os.Stat(tempPath); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("%s: No such file or directory\n", tempPath)
		return "\n"
	}

	os.Setenv("PWD", tempPath)
	return "\n"
}

// Run a general command provided by the user
func run(command string, args []string) (string, error) {
	_, found := _searchCommandInPath(command)
	if !found {
		return fmt.Sprintln(command + ": command not found"), nil
	}

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(command, args...)
	// stdout, err := cmd.Output()

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return fmt.Sprint(string(exitErr.Stderr)), err
		}

		return fmt.Sprint(err), err
	}

	output := out.Bytes()
	return fmt.Sprint(string(output)), nil
}
