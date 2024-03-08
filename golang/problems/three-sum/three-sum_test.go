package threesum

import (
	"testing"
)

func arraysElementsMatch(arr1, arr2 [][]int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for _, e := range arr1 {
		i := 0
		for i := 0; i < len(arr2); i++ {
			if arr2[i][0] == e[0] && arr2[i][1] == e[1] && arr2[i][2] == e[2] {
				break
			}
		}
		if i == len(arr2) {
			return false
		}
	}
	return true

}

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[]int{1, 0, -1, 2, -2, 1}},
			[][]int{
				{-2, 1, 1},
				{-2, 0, 2},
				{-1, 0, 1},
			},
		},
		{"148", args{[]int{1, 2, -2, -1}},
			[][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !arraysElementsMatch(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
