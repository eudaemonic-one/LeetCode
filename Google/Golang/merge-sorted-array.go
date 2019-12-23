func merge(nums1 []int, m int, nums2 []int, n int)  {
    sorted := make([]int, 0)
    i, j := 0, 0
    for i < m && j < n {
        if nums1[i] < nums2[j] {
            sorted = append(sorted, nums1[i])
            i++
        } else {
            sorted = append(sorted, nums2[j])
            j++
        }
    }
    if i < m {
        sorted = append(sorted, nums1[i:m]...)
    } else if j < n {
        sorted = append(sorted, nums2[j:n]...)
    }
    fmt.Println(sorted)
    for k := 0; k < len(sorted); k++ {
        nums1[k] = sorted[k]
    }
    return
}
