package containerwithmostwater

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func maxArea(height []int) int {

	l, r := 0, len(height)-1
	maxVolume := 0
	for l < r {
		width := r - l
		if height[r] > height[l] {
			maxVolume = max(height[l]*width, maxVolume)
			l++
		} else {
			maxVolume = max(height[r]*width, maxVolume)
			r--
		}
	}
	return maxVolume
}
