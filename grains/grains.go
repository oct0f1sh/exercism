package grains

import (
	"errors"
)

func Square(in int) (uint64, error) {
	if in > 64 || in <= 0 {
		return 0, errors.New("Input out of range")
	}

	var lastnum uint64 = 1

	for i := 2; i <= in; i++ {
		lastnum = lastnum * 2
	}

	return lastnum, nil
}

func Total() uint64 {
	var sum uint64

	for i := 0; i <= 64; i++ {
		square, _ := Square(i)
		sum += square
	}

	return sum
}
