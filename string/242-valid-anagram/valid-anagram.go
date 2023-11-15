func isAnagram(s string, t string) bool {
	if len(s) != len(t) { 
		return false 
	}

	characters := map[rune]int{} 
	for _, character := range s { 
		characters[character]++ 
	}

	for _, character := range t {
		value, exists := characters[character]
		if !exists || value <= 0 {
			return false
		}
		characters[character]--
	}

	return true
}