func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
    res := 0
    tree := make(map[int][]int)
    accumulator := make([]int, n)
    for i := 0; i < n; i++ {
        if manager[i] == -1 {
            continue
        }
        if _, ok := tree[manager[i]]; !ok {
            tree[manager[i]] = make([]int, 0)
        }
        tree[manager[i]] = append(tree[manager[i]], i)
    }
    bfs(tree, headID, accumulator, informTime)
    for _, val := range accumulator {
        res = max(res, val)
    }
    return res
}

func max(x int, y int) int {
    if x > y {
        return x
    }
    return y
}

func bfs(tree map[int][]int, root int, accumulator []int, informTime []int) {
    for _, node := range tree[root] {
        accumulator[node] = accumulator[root] + informTime[root]
        bfs(tree, node, accumulator, informTime)
    }
}