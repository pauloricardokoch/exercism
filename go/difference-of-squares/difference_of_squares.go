package diffsquares

func SquareOfSum(n int) int {
	s := n * (1 + n) / 2
	return s * s
}
func SumOfSquares(n int) (s int) {
	return n * (n + 1) * (2*n + 1) / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
