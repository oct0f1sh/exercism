package scrabble

import (
	"strings"
)

func Score(letters string) int {
	letters = strings.ToLower(letters)

	total := 0

	for _, c := range letters {
		switch string(c) {
		case "a", "e", "i", "o", "u", "l", "n", "r", "s", "t":
			total ++
		case "d", "g":
			total += 2
		case "b", "c", "m", "p":
			total += 3
		case "f", "h", "v", "w", "y":
			total += 4
		case "k":
			total += 5
		case "j", "x":
			total += 8
		case "q", "z":
			total += 10
		default:
		}
	}

	return total
}