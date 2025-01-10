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
		start := strings.IndexAny(s, "'\"\\")
		if start == -1 {
			tokens = append(tokens, strings.Fields(s)...)
			break
		}

		ch := s[start]
		fields := strings.Fields(s[:start])

		tokens = append(tokens, fields...)
		s = s[start+1:]
		end := strings.IndexByte(s, ch)

		token := s[:end]
		tokens = append(tokens, token)
		s = s[end+1:]
	}

	return tokens
}

// A non-quoted backslash ‘\’ is the Bash escape character. It preserves the literal value of
// the next character that follows, with the exception of newline.

func parseArgs(s string) []string {
	re := regexp.MustCompile(`'[^']*'|"[^"]*"|\S+`)
	matches := re.FindAllString(s, -1)

	// for i, match := range matches {
	// 	fmt.Println("DEBUG: match", i, " ->", match)
	// }

	var result []string

	for _, match := range matches {
		match_single_quotes := match[0] == '\'' && match[len(match)-1] == '\''
		match_double_quotes := match[0] == '"' && match[len(match)-1] == '"'

		if match_single_quotes || match_double_quotes {
			result = append(result, match[1:len(match)-1])
			// } else if match[0] == '\\' {
			// 	result = append(result, "")
		} else {
			result = append(result, strings.ReplaceAll(match, "\\", ""))
		}
	}

	// for i, res := range result {
	// 	fmt.Println("DEBUG: result", i, " ->", res)
	// }

	return result
}
