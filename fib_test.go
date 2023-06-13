package fib

import (
	"fmt"
	"testing"
)

func BenchmarkFib(b *testing.B) {
	testCases := []struct {
		n int
	}{
		{10},
		{20},
		{30},
	}
	for _, tt := range testCases {
		b.Run(fmt.Sprintf("fib(%d)", tt.n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var _ = Fib(tt.n)
			}
		})
	}
}
