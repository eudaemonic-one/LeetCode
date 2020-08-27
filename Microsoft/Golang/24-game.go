func judgePoint24(nums []int) bool {
    if len(nums) != 4 {
        return false
    }
    floatNums := make([]float64, 4)
    for i, num := range nums {
        floatNums[i] = float64(num)
    }
    return dfs(floatNums)
}

func dfs(nums []float64) bool {
    if len(nums) == 0 {
        return false
    }
    if len(nums) == 1 {
        return math.Abs(nums[0] - 24.0) < 1E-6
    }
    for i := 0; i < len(nums); i++ {
        for j := 0; j < len(nums); j++ {
            if i != j {
                nums2 := make([]float64, 0)
                for k := 0; k < len(nums); k++ {
                    if k != i && k != j {
                        nums2 = append(nums2, nums[k])
                    }
                }
                for k := 0; k < 4; k++ {
                    switch k {
                    case 0:
                        nums2 = append(nums2, nums[i]+nums[j])
                    case 1:
                        nums2 = append(nums2, nums[i]-nums[j])
                    case 2:
                        nums2 = append(nums2, nums[i]*nums[j])
                    case 3:
                        if nums[j] == 0 {
                            continue
                        } else {
                            nums2 = append(nums2, nums[i]/nums[j])
                        }
                    }
                    if dfs(nums2) {
                        return true
                    }
                    nums2 = nums2[:len(nums2)-1]
                }
            }
        }
    }
    return false
}
