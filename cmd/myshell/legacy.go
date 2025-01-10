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
