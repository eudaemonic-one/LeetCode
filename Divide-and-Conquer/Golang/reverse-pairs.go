func reversePairs(nums []int) int {
    if len(nums) < 2 {
        return 0
    }
    return mergeSortAndCount(nums, 0, len(nums)-1)
}

func mergeSortAndCount(nums []int, l, r int) int {
    if l >= r {
        return 0
    }
    m := l + (r-l)/2
    leftPairs := mergeSortAndCount(nums, l, m)
    rightPairs := mergeSortAndCount(nums, m+1, r)
    crossPairs := 0
    j := m+1
    for i := l; i <= m; i++ {
        for j <= r && nums[i] > nums[j] * 2 {
            j++
        }
        crossPairs += j - (m+1)
    }
    merge(nums, l, m, r)
    return leftPairs + rightPairs + crossPairs
}

func merge(nums []int, l, m, r int) {
    n1 := m - l + 1
    n2 := r - m
    L := make([]int, n1)
    for i := 0; i < n1; i++ {
        L[i] = nums[l+i]
    }
    R := make([]int, n2)
    for j := 0; j < n2; j++ {
        R[j] = nums[m+1+j]
    }
    i, j := 0, 0
    for k := l; k <= r; k++ {
        if j >= n2 || i < n1 && L[i] <= R[j] {
            nums[k] = L[i]
            i++
        } else {
            nums[k] = R[j]
            j++
        }
    }
}
