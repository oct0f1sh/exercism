package raindrops

import "strconv"

func Convert(in int) string {
	output := ""

	if in%3 == 0 {
		output += "Pling"
	}

	if in%5 == 0 {
		output += "Plang"
	}

	if in%7 == 0 {
		output += "Plong"
	}

	if output == "" {
		output = strconv.Itoa(in)
	}

	return output
}
