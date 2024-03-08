package main

func longestCommonPrefix(strs []string) string {
	/*

	   lets keep a global position of where we are in all strings.

	   we consider the character of the first string in that position.
	   all subsequent strings need to match that character.

	   so we roll through each string at that position.

	   if any string does not have the search character at the given position,
	   then return the slice of the first string from [0:globalIndex)

	   This will have time complexity O(n*m)
	   where n is the number of strings and m is the average length of the strings

	   Space complexity is O(1)
	*/

	for globalIndex := 0; globalIndex < len(strs[0]); globalIndex++ {
		searchChar := strs[0][globalIndex]
		for _, str := range strs[1:] {
			if globalIndex >= len(str) ||
				str[globalIndex] != searchChar {
				return strs[0][:globalIndex]
			}
		}
	}
	return strs[0]

}
