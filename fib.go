package fib

func Fib(u int) int {
	if u <= 1 {
		return 1
	}
	return Fib(u-2) + Fib(u-1)
}
