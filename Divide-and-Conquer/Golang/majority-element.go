func majorityElement(nums []int) int {
    return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, lo, hi int) int {
    // base case; the only element in the subarray
    if lo == hi {
        return nums[lo]
    }
    // recurse on left and right halves of this slice
    mid := lo + (hi-lo)/2
    left := helper(nums, lo, mid)
    right := helper(nums, mid+1, hi)
    // if the two halves agree on the majority element
    if left == right {
        return left
    }
    // otherwise, count each element and return the "winner"
    leftCount := countInRange(nums, left, lo, hi)
    rightCount := countInRange(nums, right, lo, hi)
    if leftCount > rightCount {
        return left
    } else {
        return right
    }
}

func countInRange(nums []int, num, lo, hi int) int {
    count := 0
    for i := lo; i <= hi; i++ {
        if nums[i] == num {
            count++
        }
    }
    return count
}
