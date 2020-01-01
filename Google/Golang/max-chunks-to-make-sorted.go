func maxChunksToSorted(arr []int) int {
    maxans := 1
    path := [][]int{}
    for i := 2; i <= len(arr); i++ {
        helper(arr, 0, path, i, &maxans)
    }
    return maxans
}

func helper(arr []int, idx int, path [][]int, chunks int, maxans *int) {
    if idx == len(arr) {
        if len(path) == chunks {
            nums := []int{}
            for _, chunk := range path {
                nums = append(nums, chunk...)
            }
            flag := true
            for i := 1; i < len(nums); i++ {
                if nums[i-1] > nums[i] {
                    flag = false
                    break
                }
            }
            if flag {
                *maxans = chunks
            }
        }
        return
    }
    for k := 1; k < len(arr)-idx+1; k++ {
        sorted := make([]int, k)
        copy(sorted, arr[idx:idx+k])
        sort.Ints(sorted)
        path = append(path, sorted)
        helper(arr, idx+k, path, chunks, maxans)
        path = path[:len(path)-1]
    }
}
