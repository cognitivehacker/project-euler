package largestprime

import (
	"testing"
)

func TestLargestPrimeFactor(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test for 9",
			args: args{n: 9},
			want: 3,
		},
		{
			name: "Test for 13195",
			args: args{n: 13195},
			want: 29,
		},
		{
			name: "Test for 600851475143",
			args: args{n: 600851475143},
			want: 29,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestPrimeFactor(tt.args.n); got != tt.want {
				t.Errorf("LargestPrimeFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test for 1",
			args: args{n: 1},
			want: true,
		},
		{
			name: "Test for 2",
			args: args{n: 2},
			want: true,
		},
		{
			name: "Test for 3",
			args: args{n: 3},
			want: true,
		},
		{
			name: "Test for 4",
			args: args{n: 4},
			want: false,
		},
		{
			name: "Test for 5",
			args: args{n: 5},
			want: true,
		},
		{
			name: "Test for 6",
			args: args{n: 6},
			want: false,
		},
		{
			name: "Test for 7",
			args: args{n: 7},
			want: true,
		},
		{
			name: "Test for 8",
			args: args{n: 8},
			want: false,
		},
		{
			name: "Test for 9",
			args: args{n: 9},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrime(tt.args.n); got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
