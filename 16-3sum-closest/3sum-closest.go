
func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    n := len(nums)
    res := nums[0]+nums[1]+nums[2]
    for i:= 0; i<n-2 ; i++ {
        j, k := i+1, n-1
        for j < k {
            tmp := nums[i]+nums[j]+nums[k]


            if tmp == target {
                return tmp
            }
            if abs(target -res) > abs(target-tmp) {
                res = tmp
            }
            if tmp > target {
                k--
            } else {
                j++
            }
        }
    }
    return res
}

func abs(n int) int {
    if n<0{
        return -n
    }
    return n
}