package threesum

func threeSum(nums []int) [][]int {
	/*
		consider the input
		1 0 -1 2 -2 1

		now lets maintain a map from a "hole" (a value we're looking for)
		to the slices of values that want the hole

		we walk i,j through the input list, and reset the hole map every time we update i.

		We know we can do this because if we were going to satisfy the triple with i, j, and k,
		we would've seen that combination (i.e. 'k') while incrementing j. Since we didn't,
		we don't need to look at that result anymore.

		Complexity:
		O(n^2) time
		O(n) space maintaining the result and hole maps
	*/

	resultMap := make(map[[3]int][]int)

	for i := 0; i < len(nums)-2; i++ {
		holeMap := make(map[int][]int)

		for j := i + 1; j < len(nums); j++ {
			// intentionally set j=i so that we check if
			// our current index satisfies any holes

			// check if j is in our hole map
			numj := nums[j]

			arr, found := holeMap[numj]
			if found {
				// for each arr, add j, and add to result map
				var answer []int

				if numj <= arr[0] {
					answer = []int{
						numj,
						arr[0],
						arr[1],
					}
				} else if numj >= arr[1] {
					answer = []int{
						arr[0],
						arr[1],
						numj,
					}
				} else {
					answer = []int{
						arr[0],
						numj,
						arr[1],
					}
				}

				resultMap[[3]int(answer)] = answer
				delete(holeMap, numj)
			} else {
				// but we don't want to add ourselves as a comparison
				sum := numj + nums[i]

				var sorted []int
				if numj < nums[i] {
					sorted = []int{
						numj,
						nums[i],
					}
				} else {
					sorted = []int{
						nums[i],
						numj,
					}
				}

				// add the negative to our hole map
				holeMap[-sum] = sorted

			}

		}
	}

	results := make([][]int, len(resultMap))

	count := 0
	for _, r := range resultMap {
		results[count] = r
		count++
	}
	return results
}
