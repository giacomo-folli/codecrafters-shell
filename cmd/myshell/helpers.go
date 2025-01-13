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
			var sliced string

			if match[len(match)-2] == '\\' {
				sliced = match[1:]
			} else {
				sliced = match[1 : len(match)-1]
			}

			temp := ""
			for i := 0; i < len(sliced); i++ {
				if i < len(sliced)-1 && sliced[i] == '\\' && (sliced[i+1] == '\\' || sliced[i+1] == '"' || sliced[i+1] == '&') {
					temp += string(sliced[i+1])
					i++
				} else {
					temp += string(sliced[i])
				}
			}

			// temp = strings.ReplaceAll(temp, `\n`, `n`)

			result = append(result, temp)
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

// ParseArgs parses a command string into arguments, similar to bash's echo
func ParseArgs(s string) []string {
	var args []string
	var current []rune
	var i int

	for i < len(s) {
		char := rune(s[i])

		switch char {
		case ' ', '\t':
			// Handle whitespace
			if len(current) > 0 {
				args = append(args, string(current))
				current = current[:0]
			}
			i++

		case '\'':
			// Handle single quoted string
			content, newPos := parseSingleQuoted(s, i+1)
			current = append(current, []rune(content)...)
			i = newPos

		case '"':
			// Handle double quoted string
			// if i+1 < len(s) {
			// 	if s[i+1] == ' ' {
			content, newPos := parseDoubleQuoted(s, i+1)
			current = append(current, []rune(content)...)
			i = newPos
			// 	} else {
			// 		i++
			// 	}
			// } else {
			// 	current = append(current, '"')
			// 	i++
			// }

		case '\\':
			// Handle escaped character
			if i+1 < len(s) {
				if s[i+1] == ' ' {
					current = append(current, ' ')
					i += 2
				} else {
					current = append(current, '\\')
					i++
				}
			} else {
				current = append(current, '\\')
				i++
			}

		default:
			current = append(current, char)
			i++
		}
	}

	// Add final argument if exists
	if len(current) > 0 {
		args = append(args, string(current))
	}

	return args
}

// parseSingleQuoted handles content within single quotes
// Returns the parsed content and the position after the closing quote
func parseSingleQuoted(s string, start int) (string, int) {
	var content []rune

	for i := start; i < len(s); i++ {
		if s[i] == '\'' {
			return string(content), i + 1
		}
		if s[i] == '\\' {
			content = append(content, '\\')
			if i+1 < len(s) {
				content = append(content, rune(s[i+1]))
				i++
			}
			continue
		}
		content = append(content, rune(s[i]))
	}

	// If no closing quote is found, return the content up to the end
	return string(content), len(s)
}

// parseDoubleQuoted handles content within double quotes
// Returns the parsed content and the position after the closing quote
func parseDoubleQuoted(s string, start int) (string, int) {
	var content []rune

	for i := start; i < len(s); i++ {
		if s[i] == '"' {
			if i+1 < len(s) {
				if s[i+1] == ' ' {
					return string(content), i + 1
				} else if s[i+1] == '"' {
					i += 2
				} else {
					i++
				}
			} else {
				return string(content), i + 1
			}
		}

		if s[i] == '\\' && i+1 < len(s) {
			switch s[i+1] {
			case 'n':
				content = append(content, '\\', 'n')
			case '"':
				content = append(content, '"')
			case '\'':
				content = append(content, '\'')
			case '\\':
				content = append(content, '\\')
			default:
				content = append(content, '\\', rune(s[i+1]))
			}
			i++
			continue
		}

		if s[i] != '"' {
			content = append(content, rune(s[i]))
		}
	}

	return string(content), len(s)
}
