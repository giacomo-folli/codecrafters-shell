package main

import (
	"fmt"
	"os"
	"strings"
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
	os.Setenv("PWD", initPwdVar())
	os.Setenv("HOME", initHomeVar())

	for {
		fmt.Fprint(os.Stdout, "$ ")

		userInput := getUserIput()

		input_string := strings.TrimRight(userInput, "\r\n")
		parsed_string := strings.Split(input_string, " ")

		command := parsed_string[0]
		args := strings.Join(parsed_string[1:], " ")

		task(command, args)
		// task(command1, args1)
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

// func parse(s string) (string, []string) {
// 	var args []string

// 	s = strings.Trim(s, "\r\n")
// 	command, argstr, _ := strings.Cut(s, " ")

// 	if strings.Contains(s, "\"") {
// 		re := regexp.MustCompile("\"(.*?)\"")
// 		args = re.FindAllString(s, -1)

// 		for i := range args {
// 			args[i] = strings.Trim(args[i], "\"")
// 		}

// 	} else if strings.Contains(s, "'") {
// 		re := regexp.MustCompile("'(.*?)'")
// 		args = re.FindAllString(s, -1)

// 		for i := range args {
// 			args[i] = strings.Trim(args[i], "'")
// 		}

// 	} else {
// 		if strings.Contains(argstr, "\\") {
// 			re := regexp.MustCompile(`[^\\] +`)
// 			args = re.Split(argstr, -1)

// 			for i := range args {
// 				args[i] = strings.ReplaceAll(args[i], "\\", "")
// 			}

// 		} else {
// 			args = strings.Fields(argstr)
// 		}
// 	}

// 	fmt.Println("DEBUG:", command, args)
// 	return command, args
// }
