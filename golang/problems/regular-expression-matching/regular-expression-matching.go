package regularexpressionmatching

func isMatch(s string, p string) bool {
	const unknown = byte(0)
	const noMatch = byte(1)
	const match = byte(2)
	cache := make([][]byte, len(s)+1)
	for i := range cache {
		cache[i] = make([]byte, len(p)+1)
		for j := range cache[i] {
			cache[i][j] = unknown
		}
	}

	var dfs func(int, int) bool
	dfs = func(iS, iP int) bool {

		if val := cache[iS][iP]; val != unknown {
			return val == match
		}

		if iS >= len(s) && iP >= len(p) {
			cache[iS][iP] = match
			return true
		} else if iP >= len(p) {
			cache[iS][iP] = noMatch
			return false
		}

		match := iS < len(s) && (s[iS] == p[iP] || p[iP] == '.')
		asterisk := iP+1 < len(p) && p[iP+1] == '*'

		if asterisk {
			return dfs(iS, iP+2) || (match && dfs(iS+1, iP))
		} else if match {
			return dfs(iS+1, iP+1)
		} else {
			cache[iS][iP] = noMatch
			return false
		}

	}

	response := dfs(0, 0)

	return response

}
