func sortColors(nums []int)  {
    zeros := 0
    ones := 0
    for _, x := range nums {
        if x == 0 {
            zeros++
        } else if x == 1 {
            ones++
        } 
    }
    for i, _ := range nums {
        if zeros>0 {
            nums[i]=0
            zeros--
        } else if ones>0 {
            nums[i]=1
            ones--
        } else {
            nums[i]=2
        }
    }

    
}