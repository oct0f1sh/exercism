package isogram

import(
	"strings"
	"unicode"
) 

func IsIsogram(input string) bool {
	chars := make(map[string]int)

	input = strings.ToLower(input)

	for _, char := range input {
		if unicode.IsLetter(char) {
			char_count := chars[string(char)]

			if char_count == 1 {
				return false
			}
	
			chars[string(char)] ++
		}
	}

	return true
}