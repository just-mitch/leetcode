package twosum

import (
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.

		{
			"first",
			args{
				[]int{2, 7, 11, 17},
				9,
			},
		},
		{
			"second",
			args{
				[]int{3, 2, 4},
				6,
			},
		},
		{
			"negative",
			args{
				[]int{3, 2, 9, -3},
				6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); tt.args.nums[got[0]]+tt.args.nums[got[1]] != tt.args.target {
				t.Errorf("twoSum() = %v, sums to %v, but should sum to %v", got, tt.args.nums[got[0]]+tt.args.nums[got[1]], tt.args.target)
			}
		})
	}
}
