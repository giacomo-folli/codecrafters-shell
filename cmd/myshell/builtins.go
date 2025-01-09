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

func _generateTokens(s string) []string {
	var tokens []string

	temp := s
	for i := range len(temp) - 1 {
		if i == len(temp)-2 {
			break
		}

		bb := []byte("'")
		if temp[i] == bb[0] && temp[i+1] == bb[0] {
			temp = temp[:i] + temp[i+2:]
		}
	}

	s = temp
	for {
		start := strings.Index(s, "'")
		if start == -1 {
			tokens = append(tokens, strings.Fields(s)...)
			break
		}

		tokens = append(tokens, strings.Fields(s[:start])...)
		s = s[start+1:]
		end := strings.Index(s, "'")
		token := s[:end]
		tokens = append(tokens, token)
		s = s[end+1:]
	}

	return tokens
}

func echo(args string) {
	s := args
	tokens := _generateTokens(s)

	fmt.Println(strings.Join(tokens, " "))
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
	args_slice := _generateTokens(args)

	_, found := searchCommandInPath(command)
	if !found {
		fmt.Println(command + ": command not found")
	} else {
		cmd := exec.Command(command, args_slice...)
		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println("Error in execution of", command, "command")
		}

		fmt.Print(string(stdout))
	}
}
