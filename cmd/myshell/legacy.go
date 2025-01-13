package main

// func _generateTokens(s string) []string {
// 	var tokens []string

// 	temp := s
// 	for i := range len(temp) - 1 {
// 		if i == len(temp)-2 {
// 			break
// 		}

// 		bb := []byte("'")
// 		cc := []byte("\"")

// 		bb_check := temp[i] == bb[0] && temp[i+1] == bb[0]
// 		cc_check := temp[i] == cc[0] && temp[i+1] == cc[0]

// 		if bb_check || cc_check {
// 			temp = temp[:i] + temp[i+2:]
// 		}
// 	}

// 	s = temp
// 	for {
// 		start := strings.IndexAny(s, "'\"\\")
// 		if start == -1 {
// 			tokens = append(tokens, strings.Fields(s)...)
// 			break
// 		}

// 		ch := s[start]
// 		fields := strings.Fields(s[:start])

// 		tokens = append(tokens, fields...)
// 		s = s[start+1:]
// 		end := strings.IndexByte(s, ch)

// 		token := s[:end]
// 		tokens = append(tokens, token)
// 		s = s[end+1:]
// 	}

// 	return tokens
// }

// func _parseArgs(s string) []string {
// 	var result []string

// 	temp := _removeInvalidQuotes(s)

// 	re := regexp.MustCompile(`'[^']*'|"[^"]*"|\S+`)
// 	matches := re.FindAllString(temp, -1)

// 	// if env := os.Getenv("ENV"); env == "LOCAL" {
// 	// 	for i, match := range matches {
// 	// 		fmt.Println("DEBUG: match", i, " ->", match)
// 	// 	}
// 	// }

// 	for _, match := range matches {
// 		match_single_quotes := match[0] == '\'' && match[len(match)-1] == '\''
// 		match_double_quotes := match[0] == '"' && match[len(match)-1] == '"'

// 		if match_single_quotes {
// 			result = append(result, match[1:len(match)-1])
// 		} else if match_double_quotes {
// 			var sliced string

// 			if match[len(match)-2] == '\\' {
// 				sliced = match[1:]
// 			} else {
// 				sliced = match[1 : len(match)-1]
// 			}

// 			temp := ""
// 			for i := 0; i < len(sliced); i++ {
// 				if i < len(sliced)-1 && sliced[i] == '\\' && (sliced[i+1] == '\\' || sliced[i+1] == '"' || sliced[i+1] == '&') {
// 					temp += string(sliced[i+1])
// 					i++
// 				} else {
// 					temp += string(sliced[i])
// 				}
// 			}

// 			// temp = strings.ReplaceAll(temp, `\n`, `n`)

// 			result = append(result, temp)
// 		} else {
// 			result = append(result, strings.ReplaceAll(match, `\`, ""))
// 		}
// 	}

// 	// if env := os.Getenv("ENV"); env == "LOCAL" {
// 	// 	fmt.Println("------------------------------")

// 	// 	for i, res := range result {
// 	// 		fmt.Println("DEBUG: result", i, " ->", res)
// 	// 	}

// 	// 	fmt.Println("------------------------------")
// 	// }

// 	return result
// }

// func _removeInvalidQuotes(s string) string {
// 	for i := range len(s) - 1 {
// 		if i == len(s)-2 {
// 			break
// 		}

// 		bb := []byte("'")
// 		cc := []byte("\"")

// 		bb_check := s[i] == bb[0] && s[i+1] == bb[0]
// 		cc_check := s[i] == cc[0] && s[i+1] == cc[0]

// 		if bb_check || cc_check {
// 			s = s[:i] + s[i+2:]
// 		}
// 	}

// 	return s
// }
