package test26

import "testing"

func Test_RemoveDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"arr0", args{nums: []int{0}}, 1},
		{"arr1", args{nums: []int{0, 1, 1, 1, 2, 2}}, 3},
		{"arr2", args{nums: []int{0, 0, 1, 1, 3, 4}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
