func canFinish(numCourses int, prerequisites [][]int) bool {
    if numCourses <= 1 {
        return true
    }
    edges := make(map[int][]int)
    indegrees := make([]int, numCourses)
    queue := make([]int, 0)
    cnt := 0
    for i := 0; i < numCourses; i++ {
        edges[i] = make([]int, 0)
    }
    for _, seq := range prerequisites {
        start, end := seq[1], seq[0]
        edges[start] = append(edges[start], end)
        indegrees[end]++
    }
    for i := 0; i < numCourses; i++ {
        if indegrees[i] == 0 {
            queue = append(queue, i)
            cnt++
        }
    }
    for len(queue) > 0 {
        currVertex := queue[0]
        queue = queue[1:]
        for j := 0; j < len(edges[currVertex]); j++ {
            nextVertex := edges[currVertex][j]
            indegrees[nextVertex]--
            if indegrees[nextVertex] == 0 {
                queue = append(queue, nextVertex)
                cnt++
            }
        }
    }
    return cnt == numCourses
}
