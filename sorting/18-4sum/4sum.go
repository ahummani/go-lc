func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)

    res := [][]int{}
    for i:= 0; i< len(nums)-3; i++ {
        if i==0 || nums[i-1] != nums[i] {
            for j := i+1; j< len(nums)-2; j++ {
                if j== i+1 || nums[j-1] != nums[j] {
                    findTwoSum(nums, i, j, target, &res)
                }
            }
        }
    }
    return res
}

func findTwoSum(nums []int, i, j, target int, res *[][]int) {
    left, right := j+1, len(nums)-1

    for left < right {
        sum := nums[i]+nums[j]+nums[left]+nums[right]
        if sum == target {
            (*res) = append(*res, []int{nums[i],nums[j], nums[left], nums[right]})
            left++
            for left<right && nums[left-1] == nums[left] {
                left++
                }
            } else if sum <target {
                left++
            } else {
                right--
            
        }
    }
}