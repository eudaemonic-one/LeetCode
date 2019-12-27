func canFinish(numCourses int, prerequisites [][]int) bool {
    if numCourses <= 1 {
        return true
    }
    edges := make(map[int][]int)
    visited := make([]bool, numCourses)
    for i := 0; i < numCourses; i++ {
        edges[i] = make([]int, 0)
    }
    for _, seq := range prerequisites {
        edges[seq[1]] = append(edges[seq[1]], seq[0])
    }
    for i := 0; i < numCourses; i++ {
        if !visited[i] && isCyclic(edges, i, visited) {
            return false
        }
    }
    return true
}

func isCyclic(edges map[int][]int, start int, visited []bool) bool {
    if visited[start] == true {
        return true
    }
    visited[start] = true
    for _, end := range edges[start] {
        if isCyclic(edges, end, visited) {
            return true
        }
    }
    visited[start] = false
    return false
}
