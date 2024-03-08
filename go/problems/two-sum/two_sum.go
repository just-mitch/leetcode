package twosum

func twoSum(nums []int, target int) []int {
	// faster solution
	// when we see a number, a at index iA,
	// we know what the corresponding number is: b at index iB.
	// the question then is if iB exists at all
	// so if we keep track of a map with nums => index
	// then when we have a new number a at iA, we check if b is in the map,
	//  and if so, we return iB. otherwise we insert a => iA into the map and continue

	m := make(map[int]int)

	for iA, v := range nums {
		if iB, found := m[target-v]; found {
			return []int{iA, iB}
		} else {
			m[v] = iA
		}
	}

	return []int{}
}
