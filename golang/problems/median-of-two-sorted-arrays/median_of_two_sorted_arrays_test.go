package medianoftwosortedarrays

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"simple",
			args{[]int{1, 2, 3, 4}, []int{6, 7, 8}},
			4},
		{"v2",
			args{[]int{1, 2, 3, 4}, []int{6, 7, 8, 9}},
			5},
		{"v3",
			args{[]int{1, 2, 3, 4}, []int{3, 7, 8, 9}},
			3.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
