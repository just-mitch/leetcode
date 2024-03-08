package medianoftwosortedarrays

import (
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	/*
	   given sorted arrays nums1 and nums2 of lengths n,m where n<=m
	   we must find the indices p1, p2 in nums1 and nums2 such that
	   the points to the left of p1 and p2 are the (n+m)/2 smallest

	   thus we may select a partition in nums1 that is midway.

	   the partition in nums2 then is about (n+m)/2 - p1

	   thus, p2 is dependent on p1, and we are satisfied when we have found the proper p1

	   our initial range for searching for p1 is [0,n]

	   if the higher point of p1, called v1h is less than v2l, then we must shift our partition p1
	   to be the midpoint of [n/2, n].

	   similarly if the lower point of p1, called v1l is greater than v2h then we shift our partition
	   to be the midpoint of [0, n/2]

	   so it seems in truth that during our iterations we need merely to keep track of
	   the bounds we are considering in our binary search throughout nums1

	   we are finished when the length of our searchable region is <= 0.
	   at this point, we should have v1h >= v2l and v1l <= v2h.

	   then, we should consult the total length, m+n.

	   if that number is even, then we return (max(v1l, v2l) + min(v1h, v2h)) / 2
	   if that number is odd, we return min(v1h,v2h) because we will orchestrate our integer division s.t.
	   the length of our lesser partition is shorter when than the other half when the length is odd.


	   this algorithm should have a time complexity of O(log(min(n,m))) and a space complexity of O(1)
	*/

	// first swap the arrays so nums1 is the shorter

	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	total := m + n

	// define the search domain in terms of left and right points within nums1
	domainL, domainR := 0, m

	for {

		// if m is 5, p1 is 2 which means that v1h is the median of p1
		// if m is 4, p1 is 2 which means that p1 straddles the median
		p1 := (domainR + domainL) / 2

		// if m is 5 and n is 7, p1 is 2 and p2 is 4. that means our entire partition has length 6, which is correct
		// if m is 5 and n is 6, p1 is 2 and p2 is 3. that means our entire partition has length 5, which is correct
		p2 := total/2 - p1

		// now compute v1l, v1h, v2l, v2h
		// but we could be at the ends so be careful.

		v1l := math.MinInt
		if p1-1 >= 0 {
			v1l = nums1[p1-1]
		}

		v1h := math.MaxInt
		if p1 < m {
			v1h = nums1[p1]
		}

		v2l := math.MinInt
		if p2-1 >= 0 {
			v2l = nums2[p2-1]
		}

		v2h := math.MaxInt
		if p2 < n {
			v2h = nums2[p2]
		}

		// now check if we have the correct values

		if v1h < v2l {
			// shift the domain to the right.
			domainL = p1 + 1
		} else if v2h < v1l {
			// shift domain to the left
			domainR = p1 - 1
		} else {
			if total%2 == 0 {
				// we have the correct partitions but which values to return?
				// in this case
				return float64((max(v1l, v2l) + min(v1h, v2h))) / 2.0

			} else {
				return float64(min(v1h, v2h))
			}

		}

	}

}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
