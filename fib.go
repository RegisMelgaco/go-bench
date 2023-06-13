package fib

func Fib(u int) int {
	return fib(u, map[int]int{})
}

func fib(u int, mem map[int]int) int {
	if u <= 1 {
		return 1
	}
	a, ok := mem[u-2]
	if !ok {
		a = fib(u-2, mem)
		mem[u-2] = a
	}

	b, ok := mem[u-1]
	if !ok {
		b = fib(u-1, mem)
		mem[u-1] = b
	}

	return a + b
}
