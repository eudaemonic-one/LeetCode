func criticalConnections(n int, connections [][]int) [][]int {    
    graph := make(map[int][]int)
    for _, conn := range connections {
        start, end := conn[0], conn[1]
        if _, ok := graph[start]; !ok {
            graph[start] = []int{end}
        } else {
            graph[start] = append(graph[start], end)
        }
        if _, ok := graph[end]; !ok {
            graph[end] = []int{start}
        } else {
            graph[end] = append(graph[end], start)
        }
    }
    criticalConns := make([][]int, 0)
    timestamps := make([]int, n)
    clock = 1
    
    dfs(&criticalConns, timestamps, graph, -1, 0)
    
    return criticalConns
}

var clock int

func dfs(criticalConns *[][]int, timestamps []int, graph map[int][]int, prev, idx int) int {
    if timestamps[idx] != 0 {
        return timestamps[idx]
    }
    timestamps[idx] = clock
    clock++
    
    minTimestamp := math.MaxInt64
    for _, neighbor := range graph[idx] {
        if neighbor == prev {
            continue
        }
        neighborTimestamp := dfs(criticalConns, timestamps, graph, idx, neighbor)
        minTimestamp = min(minTimestamp, neighborTimestamp)
    }
    
    if prev >= 0 && minTimestamp >= timestamps[idx] {
        *criticalConns = append(*criticalConns, []int{prev, idx})
    }
    return min(minTimestamp, timestamps[idx])
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
