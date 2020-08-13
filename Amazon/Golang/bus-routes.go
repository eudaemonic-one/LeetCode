type pair struct {
    bus   int
    depth int
}

func numBusesToDestination(routes [][]int, S int, T int) int {
    if S == T {
        return 0
    }
    
    N := len(routes)
    intersectMap := make(map[int][]int)
    
    for i := 0; i < N; i++ {
        sort.Ints(routes[i])
    }
    
    for i := 0; i < N; i++ {
        for j := i+1; j < N; j++ {
            if hasIntersection(routes[i], routes[j]) {
                if _, ok := intersectMap[i]; !ok {
                    intersectMap[i] = []int{}
                }
                intersectMap[i] = append(intersectMap[i], j)
                
                if _, ok := intersectMap[j]; !ok {
                    intersectMap[j] = []int{}
                }
                intersectMap[j] = append(intersectMap[j], i)
            }
        }
    }
    
    visited := make(map[int]bool)
    targets := make(map[int]bool)
    queue := make([]pair, 0)
    
    for i := 0; i < N; i++ {
        if idx := sort.Search(len(routes[i]), func (j int) bool { return routes[i][j] >= S }); idx < len(routes[i]) && routes[i][idx] == S {
            visited[i] = true
            queue = append(queue, pair{i, 0})
        }
        if idx := sort.Search(len(routes[i]), func (j int) bool { return routes[i][j] >= T }); idx < len(routes[i]) && routes[i][idx] == T {
            targets[i] = true
        }
    }
    
    for len(queue) > 0 {
        info := queue[0]
        queue = queue[1:]
        bus, depth := info.bus, info.depth
        if _, ok := targets[bus]; ok {
            return depth + 1
        }
        for _, nei := range intersectMap[bus] {
            if _, ok := visited[nei]; !ok {
                visited[nei] = true
                queue = append(queue, pair{nei, depth+1})
            }
        }
    }
    
    return -1
}

func hasIntersection(routeA, routeB []int) bool {
    i, j := 0, 0
    for i < len(routeA) && j < len(routeB) {
        if routeA[i] == routeB[j] {
            return true
        }
        if routeA[i] < routeB[j] {
            i++
        } else {
            j++
        }
    }
    return false
}
