package main

import (
	"testing"
)

func TestLargestPalindromeProduct(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Testing for a:91, b:99",
			args: args{a: 91, b: 99},
			want: 9009,
		},
		{
			name: "Testing for a:91, b:99",
			args: args{a: 999, b: 999},
			want: 9009,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestPalindromeProduct(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("LargestPalindromeProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Testing for 101",
			args: args{n: 101},
			want: true,
		},
		{
			name: "Testing for 10",
			args: args{n: 10},
			want: false,
		},
		{
			name: "Testing for 9009",
			args: args{n: 9009},
			want: true,
		},
		{
			name: "Testing for 546",
			args: args{n: 546},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.n); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
