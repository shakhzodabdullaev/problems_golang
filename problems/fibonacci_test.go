package problems_test

import (
	"testing"

	"github.com/shakhzodabdullaev/problem_solving/problems"
)

func TestFibonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first fib",
			args: args{1},
			want: 1,
		},
		{
			name: "first fib",
			args: args{2},
			want: 2,
		},
		{
			name: "first fib",
			args: args{3},
			want: 3,
		},
		{
			name: "first fib",
			args: args{5},
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problems.Fibonacci(tt.args.n); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
