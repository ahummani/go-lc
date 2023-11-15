func merge(intervals [][]int) [][]int {
	mp := make(map[int]int)
	var start int
	var end int
	for i := 0; i < len(intervals); i++ {
		start = intervals[i][0]
		end = intervals[i][1]
		mp[start] += 1
		mp[end] -= 1
	}
	var keys []int
	for key := range mp {
		keys = append(keys, key)
	}

	sort.Ints(keys)
	start = 0
	count := 0
	var answer [][]int
	var value int
	for _, key := range keys {
		value = mp[key]
		if count == 0 {
			start = key
		}
		count += value

		if count == 0 {
			answer = append(answer, []int{start, key})
		}
	}
	return answer
}