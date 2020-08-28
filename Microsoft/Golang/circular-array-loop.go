func circularArrayLoop(nums []int) bool {
    if len(nums) <= 1 {
        return false
    }
    n := len(nums)
    for i := 0; i < n; i++ {
        if nums[i] == 0 {
            continue
        }
        j := i
        k := getNext(nums, i)
        for nums[k] * nums[i] > 0 && nums[getNext(nums, k)] * nums[i] > 0 {
            if j == k {
                if j != getNext(nums, j) {
                    return true
                }
                break
            }
            j = getNext(nums, j)
            k = getNext(nums, getNext(nums, k))
        }
        j = i
        for nums[j] * nums[i] > 0 {
            next := getNext(nums, j)
            nums[j] = 0
            j = next
        }
    }
    return false
}

func getNext(nums []int, i int) int {
    n := len(nums)
    if nums[i] + i >= 0 {
        return (nums[i] + i) % n
    } else {
        return n + (nums[i] + i) % n
    }
}
