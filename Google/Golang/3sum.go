func threeSum(nums []int) [][]int {
    res := make([][]int, 0)
    n := len(nums)
    if n < 3 {
        return res
    }
    sort.Ints(nums)
    for i := 0; i < n - 2; i++ {
        if (nums[i] > 0) {
            break
        }
        if (i == 0 || (i > 0 && nums[i] != nums[i-1])) {
            j := i + 1
            k := n - 1
            for j < k {
                a := nums[i]
                b := nums[j]
                c := nums[k]
                total := a + b + c
                if (total == 0) {
                    cand := []int{a, b, c}
                    res = append(res, cand)
                    for j < k && nums[j] == nums[j+1] {
                        j += 1
                    }
                    for j < k && nums[k] == nums[k-1] {
                        k -= 1
                    }
                    j += 1
                    k -= 1
                } else if (total < 0) {
                    j += 1
                } else {
                    k -= 1
                }
            }
        }
    }
    return res
}

