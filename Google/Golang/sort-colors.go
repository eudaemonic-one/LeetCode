func sortColors(nums []int)  {
    zeros, ones, twos := 0, 0, 0
    for _, n := range nums {
        switch n {
        case 0:
            zeros++
        case 1:
            ones++
        case 2:
            twos++
        }
    }
    i := 0
    for cnt := 0; cnt < zeros; i, cnt = i+1, cnt+1 {
        nums[i] = 0
    }
    for cnt := 0; cnt < ones; i, cnt = i+1, cnt+1 {
        nums[i] = 1
    }
    for cnt := 0; cnt < twos; i, cnt = i+1, cnt+1 {
        nums[i] = 2
    }
}
