package recursions

func Factorial(n uint) uint64 {
	if n == 0 || n == 1 {
		return 1
	}

	return uint64(n) * Factorial(n-1)
}
