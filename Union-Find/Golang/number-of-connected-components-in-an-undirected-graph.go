func countComponents(n int, edges [][]int) int {
    if n <= 1 {
        return n
    }
    parent := make([]int, n)
    for i := 0; i < n; i++ {
        parent[i] = i
    }
    count := n
    for i := 0; i < len(edges); i++ {
        s, e := edges[i][0], edges[i][1]
        rootS := find(parent, s)
        rootE := find(parent, e)
        if rootS != rootE {
            union(parent, rootS, rootE)
            count--
        }
    }
    return count
}

func find(parent []int, x int) int {
    for x != parent[x] {
        x = parent[x]
    }
    return x
}

func union(parent []int, x, y int) {
    parent[x] = y
}
