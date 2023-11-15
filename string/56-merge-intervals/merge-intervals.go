func merge(intervals [][]int) [][]int {
    res := make([][]int, 0)

    sort.Slice(intervals, func(i,j int) bool {
        return intervals[i][0]<intervals[j][0]
    })

    res = append(res, intervals[0])


    for _, interval := range intervals[1:] {
        if res[len(res)-1][1]>=interval[0] {
            if res[len(res)-1][1] <interval[1] {
                res[len(res)-1][1] = interval[1]
            }
            continue
        }
        res = append(res, interval)
    }  
    return res
}