package grains

import (
	"errors"
)

func Square(in uint64) (uint64, error) {
	if in > 64 || in < 0 {
		return 0, errors.New("Input out of range")
	}

	return in * in, nil
}

func Total() uint64 {
	var sum uint64

	for i := 0; i <= 64; i++ {
		square, _ := Square(uint64(i))
		sum += square
	}

	return sum
}