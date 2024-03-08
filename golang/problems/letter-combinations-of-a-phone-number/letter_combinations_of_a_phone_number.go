package lettercombinationsofaphonenumber

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	alpha := map[rune][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	chars := make([][]string, len(digits))
	for i, d := range digits {
		chars[i] = alpha[rune(d)]
	}

	head, tail := chars[0], chars[1:]

	for _, ti := range tail {
		newHead := make([]string, len(head)*len(ti))
		for j, hj := range head {
			for k, tik := range ti {
				newHead[j*len(ti)+k] = hj + tik
			}
		}
		head = newHead
	}

	return head

}
