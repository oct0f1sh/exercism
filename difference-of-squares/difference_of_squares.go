package diffsquares

func SquareOfSum(num int) int {
	squareOfSum := 0

	for i := 1; i <= num; i++ {
		squareOfSum += i
	}

	return squareOfSum * squareOfSum
}

func SumOfSquares(num int) int {
	sumOfSquares := 0

	for i := 1; i <= num; i++ {
		sumOfSquares += i * i
	}

	return sumOfSquares
}

func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}