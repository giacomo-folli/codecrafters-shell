package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
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

func _parseArgs(s string) []string {
	var result []string

	temp := _removeInvalidQuotes(s)

	re := regexp.MustCompile(`'[^']*'|"[^"]*"|\S+`)
	matches := re.FindAllString(temp, -1)

	// if env := os.Getenv("ENV"); env == "LOCAL" {
	// 	for i, match := range matches {
	// 		fmt.Println("DEBUG: match", i, " ->", match)
	// 	}
	// }

	for _, match := range matches {
		match_single_quotes := match[0] == '\'' && match[len(match)-1] == '\''
		match_double_quotes := match[0] == '"' && match[len(match)-1] == '"'

		if match_single_quotes {
			result = append(result, match[1:len(match)-1])
		} else if match_double_quotes {
			sliced := match[1 : len(match)-1]
			// \, $, "

			sliced = strings.ReplaceAll(sliced, `\\`, `\`)
			sliced = strings.ReplaceAll(sliced, `\$`, `$`)
			sliced = strings.ReplaceAll(sliced, `\"`, `"`)
			// temp = strings.ReplaceAll(temp, `\n`, `n`)

			result = append(result, sliced)
		} else {
			result = append(result, strings.ReplaceAll(match, `\`, ""))
		}
	}

	// if env := os.Getenv("ENV"); env == "LOCAL" {
	// 	fmt.Println("------------------------------")

	// 	for i, res := range result {
	// 		fmt.Println("DEBUG: result", i, " ->", res)
	// 	}

	// 	fmt.Println("------------------------------")
	// }

	return result
}

func _removeInvalidQuotes(s string) string {
	for i := range len(s) - 1 {
		if i == len(s)-2 {
			break
		}

		bb := []byte("'")
		cc := []byte("\"")

		bb_check := s[i] == bb[0] && s[i+1] == bb[0]
		cc_check := s[i] == cc[0] && s[i+1] == cc[0]

		if bb_check || cc_check {
			s = s[:i] + s[i+2:]
		}
	}

	return s
}
