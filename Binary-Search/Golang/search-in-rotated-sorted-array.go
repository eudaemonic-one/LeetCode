func search(nums []int, target int) int {
    n := len(nums)
    if n == 0 {
        return -1
    } else if n == 1 {
        if nums[0] == target {
            return 0
        } else {
            return -1
        }
    }
    rotateIndex := findRotateIndex(nums, 0, n-1)
    if nums[rotateIndex] == target {
        return rotateIndex
    }
    if rotateIndex == 0 {
        return binarySearch(nums, 0, n-1, target)
    }
    if target < nums[0] {
        return binarySearch(nums, rotateIndex, n-1, target)
    } else {
        return binarySearch(nums, 0, rotateIndex-1, target)
    }
}

func findRotateIndex(nums []int, left, right int) int {
    if nums[left] < nums[right] {
        return 0
    }
    for left <= right {
        pivot := left + (right-left) / 2
        if nums[pivot] > nums[pivot+1] {
            return pivot + 1
        } else {
            if nums[pivot] < nums[left] {
                right = pivot - 1
            } else {
                left = pivot + 1
            }
        }
    }
    return 0
}

func binarySearch(nums []int, left, right, target int) int {
    for left <= right {
        pivot := left + (right-left) / 2
        if nums[pivot] == target {
            return pivot
        } else {
            if target < nums[pivot] {
                right = pivot - 1
            } else {
                left = pivot + 1
            }
        }
    }
    return -1
}
