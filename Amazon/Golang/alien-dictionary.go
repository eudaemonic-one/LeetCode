func alienOrder(words []string) string {
    if len(words) == 0 {
        return ""
    }
    // 0=todo 1=doing 2=done
    visited := make(map[byte]int)
    edges := make(map[[2]byte]bool)
    order := []byte{}
    
    for i, word := range words {
        for j := 0; j < len(word); j++ {
            visited[word[j]] = 0
        }
        if i > 0 {
            if len(words[i-1]) > len(words[i]) && strings.HasPrefix(words[i-1], words[i]) {
                return ""
            }
            for j := 0; j < min(len(words[i-1]), len(words[i])); j++ {
                if words[i-1][j] != words[i][j] {
                    edges[[2]byte{words[i-1][j],words[i][j]}] = true
                    break
                }
            }
        }
    }
    
    for vertice, status := range visited {
        if status == 0 {
            if !dfs(&order, vertice, visited, edges) {
                return ""
            }
        }
    }
    
    reverse(order)
    
    return string(order)
}

func dfs(order *[]byte, start byte, visited map[byte]int, edges map[[2]byte]bool) bool {
    visited[start] = 1
    for end := range visited {
        if edges[[2]byte{start,end}] == true {
            if visited[end] == 1 {
                return false
            } else if visited[end] == 0 {
                if !dfs(order, end, visited, edges) {
                    return false
                }
            }
        }
    }
    visited[start] = 2
    *order = append(*order, start)
    return true
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func reverse(bytes []byte) []byte {
    for l, r := 0, len(bytes)-1; l < r; l, r = l+1, r-1 {
        bytes[l], bytes[r] = bytes[r], bytes[l]
    }
    return bytes
}
