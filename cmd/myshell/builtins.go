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

func _base_echo(args []string) {
	stringg := strings.Join(args, " ")
	fmt.Print(stringg[1:len(stringg)-1], "\n")
}

func _adv_echo(args []string) {
	fmt.Print(strings.Join(args, " "), "\n")
}

func _generateTokens(s string) []string {
	var tokens []string

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
	// args_slice := strings.Split(args, " ")

	// start := strings.HasPrefix(args_slice[0], "'")
	// end := strings.HasSuffix(args_slice[len(args_slice)-1], "'")

	// if start && end {
	// 	_base_echo(args_slice)
	// 	return
	// }

	// formatted_Args := slices.DeleteFunc(args_slice, func(s string) bool {
	// 	return s == ""
	// })

	// _adv_echo(formatted_Args)
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
	args_slice := strings.Split(args, " ")

	_, found := searchCommandInPath(command)
	if !found {
		fmt.Println(command + ": command not found")
	} else {
		cmd := exec.Command(command, args_slice...)
		stdout, _ := cmd.Output()

		fmt.Print(string(stdout))
	}
}
