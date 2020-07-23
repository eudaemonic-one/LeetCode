func maxNumber(nums1 []int, nums2 []int, k int) []int {
    res := make([]int, k)
    for i := 0; i <= k; i++ {
        j := k - i
        if i > len(nums1) || j > len(nums2) {
            continue
        }
        left := findKGreatestDigitsInOrder(nums1, i)
        right := findKGreatestDigitsInOrder(nums2, j)
        merged := mergeNumsArray(left, right)
        if greaterThan(merged, res) {
            copy(res, merged)
        }
    }
    return res
}

func findKGreatestDigitsInOrder(nums []int, k int) []int {
    if k > len(nums) || k == 0{
        return []int{}
    } else if k == len(nums) {
        return nums
    }
    stack := make([]int, 0)
    for i := 0; i < len(nums); i++ {
        for len(stack) > 0 && nums[i] > stack[len(stack)-1] && len(nums)-i > k-len(stack) {
            stack = stack[:len(stack)-1]
        }
        if len(stack) < k {
            stack = append(stack, nums[i])
        }
    }
    return stack
}

func mergeNumsArray(nums1, nums2 []int) []int {
    arr := make([]int, 0)
    i, j := 0, 0
    for i < len(nums1) && j < len(nums2) {
        if greaterThan(nums1[i:], nums2[j:]) {
            arr = append(arr, nums1[i])
            i++
        } else {
            arr = append(arr, nums2[j])
            j++
        }
    }
    for i < len(nums1) {
        arr = append(arr, nums1[i])
        i++
    }
    for j < len(nums2) {
        arr = append(arr, nums2[j])
        j++
    }
    return arr
}

func greaterThan(nums1, nums2 []int) bool {
    for i := 0; i < len(nums1); i++ {
        if i >= len(nums2) {
            return true
        }
        if nums1[i] > nums2[i] {
            return true
        } else if nums1[i] < nums2[i] {
            return false
        }
    }
    return false
}

// Approach #1: Recursion

// func maxNumber(nums1 []int, nums2 []int, k int) []int {
//     res := make([]int, 0)
//     maxNum := 0
//     helper(&res, &maxNum, []int{}, nums1, nums2, 0, 0, k)
//     return res
// }

// func helper(maxPath *[]int, maxNum *int, path []int, nums1, nums2 []int, i1, i2, k int) {
//     if len(path) >= k {
//         if val := arrToInt(path); val > *maxNum {
//             *maxNum = val
//             tmpPath := make([]int, len(path))
//             copy(tmpPath, path)
//             *maxPath = tmpPath
//         }
//         return
//     }
//     if i1 < len(nums1) {
//         helper(maxPath, maxNum, path, nums1, nums2, i1+1, i2, k) 
//         path = append(path, nums1[i1])
//         helper(maxPath, maxNum, path, nums1, nums2, i1+1, i2, k) 
//         path = path[:len(path)-1]
//     }
//     if i2 < len(nums2) {
//         helper(maxPath, maxNum, path, nums1, nums2, i1, i2+1, k) 
//         path = append(path, nums2[i2])
//         helper(maxPath, maxNum, path, nums1, nums2, i1, i2+1, k) 
//         path = path[:len(path)-1]
//     }
// }

// func arrToInt(nums []int) int {
//     res := 0
//     base := 1
//     for i := len(nums)-1; i >= 0; i-- {
//         res += base * nums[i]
//         base *= 10
//     }
//     return res
// }
