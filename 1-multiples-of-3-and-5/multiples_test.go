package multiples

import "testing"

func TestFindMultiples(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test for n = 10",
			args: args{n: 10},
			want: 23,
		},
		{
			name: "Test for n = 13",
			args: args{n: 13},
			want: 45,
		},
		{
			name: "Test for n = 13",
			args: args{n: 1000},
			want: 233168,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMultiples(tt.args.n); got != tt.want {
				t.Errorf("FindMultiples() = %v, want %v", got, tt.want)
			}
		})
	}
}
