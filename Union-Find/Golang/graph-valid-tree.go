func validTree(n int, edges [][]int) bool {
    if len(edges) != n-1 {
        return false
    }
    nums := make([]int, n)
    for i := 0; i < n; i++ {
        nums[i] = i
    }
    for i := 0; i < n-1; i++ {
        x := find(nums, edges[i][0])
        y := find(nums, edges[i][1])
        if x == y {
            return false
        }
        nums[x] = y
    }
    return true
}

func find(parent []int, x int) int {
    for parent[x] != x {
        x = parent[x]
    }
    return x
}
