package _89_rotate_array

import "testing"

func Test_rotate2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{"arr0", args{nums: []int{1, 2, 3}, k: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate1(tt.args.nums, tt.args.k)
			rotate2(tt.args.nums, tt.args.k)
		})
	}
}
