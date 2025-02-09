package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"golang.org/x/term"
)

func _initEnv() {
	home, home_err := os.UserHomeDir()
	wd, wd_err := os.Getwd()

	if home_err != nil || wd_err != nil {
		log.Panicln("Error while getting  env variables! :(")
	}

	if err := os.Setenv("HOME", home); err != nil {
		log.Panicln("Error while setting home env variables! :(")
	}

	if err := os.Setenv("PWD", wd); err != nil {
		log.Panicln("Error while setting working dir env variables! :(")
	}
}

func _searchCommandInPath(command string) (string, bool) {
	path, err := exec.LookPath(command)
	if err == nil {
		return path, true
	}

	return "", false
}

func _searchBuildin(command string) bool {
	found := false

	for i := 0; i < len(builtins); i++ {
		if builtins[i] == command {
			found = true
			break
		}
	}

	return found
}

func _getUserInput() (input string) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(os.Stdin)

loop:
	for {
		c, _, err := r.ReadRune()
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer term.Restore(int(os.Stdin.Fd()), oldState)

		switch c {
		case '\x03': // Ctrl+C
			term.Restore(int(os.Stdin.Fd()), oldState)
			fmt.Print("\n")
			os.Exit(0)

		case '\r', '\n': // Enter
			fmt.Fprint(os.Stdout, "\r\n")
			break loop

		case '\x7F': // Backspace
			if length := len(input); length > 0 {
				input = input[:length-1]
				fmt.Fprint(os.Stdout, "\b \b")
			}

		case '\t': // Tab
			suffix := _autocomplete(input)
			if suffix != "" {
				input += suffix + " "
				fmt.Fprint(os.Stdout, suffix+" ")
			} else {
				fmt.Fprint(os.Stdout, "\a")
			}

		default:
			input += string(c)
			fmt.Fprint(os.Stdout, string(c))
		}
	}

	return
}

func _autocomplete(input string) (suffix string) {
	if input == "" {
		return
	}

	for i := range builtins {
		if strings.Contains(builtins[i], input) {
			a, _ := strings.CutPrefix(builtins[i], input)
			return a
		}
	}

	c, found := _searchCommandInPath(input)
	if found {
		a, _ := strings.CutPrefix(path.Base(c), input)
		fmt.Println("ei", c, found, a)
		return a
	}

	return
}

func _parseArgs(s string) []string {
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
			content, newPos := _parseSingleQuoted(s, i+1)
			current = append(current, []rune(content)...)
			i = newPos

		case '"':
			// Handle double  quoted string
			content, newPos := _parseDoubleQuoted(s, i+1)
			current = append(current, []rune(content)...)
			i = newPos

		case '\\':
			// Handle escaped character
			if i+1 < len(s) {
				if s[i+1] == ' ' {
					current = append(current, ' ')
					i += 2
				} else {
					current = append(current, rune(s[i+1]))
					i += 2
				}
			} else {
				current = append(current, rune(s[i]))
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

// _parseSingleQuoted handles content within single quotes
// Returns the parsed content and the position after the closing quote
func _parseSingleQuoted(s string, start int) (string, int) {
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

// _parseDoubleQuoted handles content within double quotes
// Returns the parsed content and the position after the closing quote
func _parseDoubleQuoted(s string, start int) (string, int) {
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
			case '\\':
				content = append(content, '\\')
			case '\'':
				content = append(content, '\\')
				content = append(content, '\'')
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

func _checkRedirection(args []string) (bool, bool, bool, []string, []string) {
	var redirection_string []string
	args_string := args

	for i := range len(args) {
		if args[i] == "2>" {
			args_string = args[:i]
			redirection_string = args[i+1:]

			return false, true, false, args_string, redirection_string
		}

		if args[i] == "2>>" {
			args_string = args[:i]
			redirection_string = args[i+1:]

			return false, true, true, args_string, redirection_string
		}

		if args[i] == "1>" || args[i] == ">" {
			args_string = args[:i]
			redirection_string = args[i+1:]

			return true, false, false, args_string, redirection_string
		}

		if args[i] == "1>>" || args[i] == ">>" {
			args_string = args[:i]
			redirection_string = args[i+1:]

			return true, false, true, args_string, redirection_string
		}
	}

	return false, false, false, args_string, nil
}

func _writeToFile(file string, data string, appendIt bool) error {
	var flags int = os.O_CREATE | os.O_WRONLY
	if appendIt {
		flags = flags | os.O_APPEND
	}

	if _, err := os.Stat(file); os.IsNotExist(err) {
		path := path.Dir(file)
		os.MkdirAll(path, 0700)
	}

	f, err := os.OpenFile(file, flags, 0644)
	if err != nil {
		return err
	}

	if _, err := f.Write([]byte(data)); err != nil {
		f.Close()
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
