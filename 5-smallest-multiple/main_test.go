package main

import (
	"testing"
)

func TestSmallestMultiple(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Testing for a:10",
			args: args{a: 10},
			want: 2520,
		},
		{
			name: "Testing for a:10",
			args: args{a: 20},
			want: 232792560,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SmallestMultiple(tt.args.a); got != tt.want {
				t.Errorf("SmallestMultiple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDivisibleByRange(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Testing for a:10 b:2",
			args: args{a: 2520, b: 10},
			want: true,
		},
		{
			name: "Testing for a:10 b:2",
			args: args{a: 2, b: 10},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDivisibleByRange(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("IsDivisibleByRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
