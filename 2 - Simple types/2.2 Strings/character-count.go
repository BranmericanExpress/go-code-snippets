func countCharacters(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, c := range s {
		counts[c]++
	}
	return counts
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
