package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func initPwdVar() string {
	dir, err := os.Getwd()
	if err == nil {
		return dir
	}

	return ""
}

func initHomeVar() string {
	dir, err := os.UserHomeDir()
	if err == nil {
		return dir
	}

	return ""
}

func searchCommandInPath(command string) (string, bool) {
	path, err := exec.LookPath(command)
	if err == nil {
		return path, true
	}

	return "", false
}

func searchBuildin(command string) bool {
	found := false

	for i := 0; i < len(builtins); i++ {
		if builtins[i] == command {
			found = true
			break
		}
	}

	return found
}

func getUserIput() string {
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error in reading string\n")
		os.Exit(1)
	}

	return input
}

func _generateTokens(s string) []string {
	var tokens []string

	temp := s
	for i := range len(temp) - 1 {
		if i == len(temp)-2 {
			break
		}

		bb := []byte("'")
		cc := []byte("\"")

		bb_check := temp[i] == bb[0] && temp[i+1] == bb[0]
		cc_check := temp[i] == cc[0] && temp[i+1] == cc[0]

		if bb_check || cc_check {
			temp = temp[:i] + temp[i+2:]
		}
	}

	s = temp
	for {
		start := strings.IndexAny(s, "'\"")
		if start == -1 {
			tokens = append(tokens, strings.Fields(s)...)
			break
		}

		ch := s[start]
		tokens = append(tokens, strings.Fields(s[:start])...)
		s = s[start+1:]
		end := strings.IndexByte(s, ch)
		token := s[:end]
		tokens = append(tokens, token)
		s = s[end+1:]
	}

	return tokens
}
