package palindromenumber

import "fmt"

func isPalindrome(x int) bool {

	s := fmt.Sprint(x)

	head, tail := 0, len(s)-1
	for head < tail {
		if s[head] != s[tail] {
			return false
		}
		head++
		tail--
	}

	return s[head] == s[tail]

}
