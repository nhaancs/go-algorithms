package recursions

// Given an index N, return is the value at the that index of the Fibonacci sequence, where the sequence is:
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144 ...
// the pattern os the sequence is that each value is the sum of the 2 previous values, that means that for N=5 -> 2+3

func FibonacciNumber(n uint) uint { //O(2^n)
	if n < 2 {
		return n
	}
	return FibonacciNumber(n-1) + FibonacciNumber(n-2)
}