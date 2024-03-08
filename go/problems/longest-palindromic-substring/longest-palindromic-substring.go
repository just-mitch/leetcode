package longestpalindromicsubstring

func longestPalindrome(s string) string {

	if len(s) == 0 {
		return ""
	}

	longest := s[0:1]

	for center := 1; center < len(s); center++ {

		iL, iR := center-1, center
		for iL >= 0 && iR < len(s) && s[iL] == s[iR] {
			iL--
			iR++
		}
		iL++
		iR--
		if iR-iL+1 > len(longest) {
			longest = s[iL : iR+1]
		}

		iL, iR = center-1, center+1
		for iL >= 0 && iR < len(s) && s[iL] == s[iR] {
			iL--
			iR++
		}
		iL++
		iR--
		if iR-iL+1 > len(longest) {
			longest = s[iL : iR+1]
		}

	}

	return longest

}
