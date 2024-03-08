package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	/*

	   keep track of the start positions of our current substring,
	   and our all time longest substring length

	   then if we update a value in the map that is in our substring
	   we need to update our start position to be the old value plus 1

	   at the end return the all time high

	*/

	ath := 0
	startSubstring := 0
	indices := make(map[rune]int)

	for i, c := range s {
		oldVal, found := indices[c]
		indices[c] = i
		if found && oldVal >= startSubstring {
			startSubstring = oldVal + 1
		}

		if currLength := i - startSubstring + 1; currLength > ath {
			ath = currLength
		}

	}

	return ath

}
