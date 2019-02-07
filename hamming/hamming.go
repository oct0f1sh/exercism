package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Both strands must be the same size")
	}

	differences := 0

	for i := range a {
		if a[i] != b[i] {
			differences++
		}
	}

	return differences, nil
}
