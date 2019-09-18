package jewelsandstones

// O(n^2)
func numJewelsInStones1(J string, S string) int {
	var result int
	for _, c1 := range S {
		for _, c2 := range J {
			if c2 == c1 {
				result++
			}
		}
	}
	return result
}

// O(n)
func numJewelsInStones2(J string, S string) int {
	var result int
	jewels := make(map[rune]bool)

	for _, jewelChar := range J {
		jewels[jewelChar] = true
	}

	for _, c := range S {
		if _, ok := jewels[c]; ok {
			result++
		}
	}
	return result
}
