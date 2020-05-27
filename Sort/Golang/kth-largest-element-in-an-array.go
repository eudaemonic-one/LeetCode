import "math/rand"

func findKthLargest(nums []int, k int) int {
    l, r := 0, len(nums)-1
    for true {
        pi := partition(nums, l, r)
        if pi == k - 1 {
            return nums[pi]
        } else if pi > k - 1 {
            r = pi - 1
        } else {
            l = pi + 1
        }
    }
    return l
}

func partition(nums []int, l, r int) int {
    randNum := l + rand.Int() % (r-l+1)
    nums[l], nums[randNum] = nums[randNum], nums[l]
    pivot := nums[l]
    i, j := l, r
    for i < j {
        for i < j && nums[j] <= pivot {
            j--
        }
        for i < j && nums[i] >= pivot {
            i++
        }
        if i < j {
            nums[i], nums[j] = nums[j], nums[i]
        }
    }
    nums[i], nums[l] = nums[l], nums[i]
    return i
}
