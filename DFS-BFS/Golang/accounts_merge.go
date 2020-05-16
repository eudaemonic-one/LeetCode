func accountsMerge(accounts [][]string) [][]string {
    graph := make(map[string][]string)
    emailToName := make(map[string]string)
    
    for _, account := range accounts {
        name := account[0]
        firstEmail := account[1]
        for _, email := range account[1:] {
            if _, ok := graph[account[1]]; !ok {
                graph[firstEmail] = []string{email}
            } else {
                graph[firstEmail] = append(graph[firstEmail], email)
            }
            
            if _, ok := graph[email]; !ok {
                graph[email] = []string{firstEmail}
            } else {
                graph[email] = append(graph[email], firstEmail)
            }
            
            emailToName[email] = name
        }
    }
    
    visited := make(map[string]bool)
    res := make([][]string, 0)
    
    for email := range graph {
        if _, ok := visited[email]; !ok {
            visited[email] = true
            stack := []string{email}
            emails := make([]string, 0)
            for len(stack) > 0 {
                node := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                emails = append(emails, node)
                for _, nei := range graph[node] {
                    if _, ok := visited[nei]; !ok {
                        visited[nei] = true
                        stack = append(stack, nei)
                    }
                }
            }
            sort.Strings(emails)
            res = append(res, append([]string{emailToName[email]}, emails...))
        }
    }
    return res
}
