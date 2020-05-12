package main

import "testing"

func TestSumFiboEven(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test for 10",
			args: args{n: 10},
			want: 10,
		},
		{
			name: "Test for 10",
			args: args{n: 4000000},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumFiboEven(tt.args.n); got != tt.want {
				t.Errorf("SumFiboEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
