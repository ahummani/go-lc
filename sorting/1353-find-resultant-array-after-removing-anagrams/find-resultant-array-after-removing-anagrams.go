func removeAnagrams(words []string) []string {

	array := []string{}
	array = append(array, words[0])

	for i := 0; i < len(words)-1; i++ {
		if !isAnagram(words[i], words[i+1]) {
			array = append(array, words[i+1])
		}
	}

	return array
}

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	hashTable := make(map[rune]int)

	for _, item := range s {
		hashTable[item]++
	}

	for _, item := range t {
		hashTable[item]--
	}

	for _, item := range hashTable {
		if item != 0 {
			return false
		}
	}

	return true
}