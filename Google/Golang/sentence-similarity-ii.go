func areSentencesSimilarTwo(words1 []string, words2 []string, pairs [][]string) bool {
    if len(words1) != len(words2) {
        return false
    }
    count := 0
    index := make(map[string]int)
    parent := make(map[int]int)
    buildUnionSet(parent, len(pairs))
    for _, pair := range pairs {
        w1, w2 := pair[0], pair[1]
        if _, ok1 := index[w1]; !ok1 {
            index[w1] = count
            count++
        }
        if _, ok2 := index[w2]; !ok2 {
            index[w2] = count
            count++
        }
        union(parent, index[w1], index[w2])
    }
    for i := 0; i < len(words1); i++ {
        w1, w2 := words1[i], words2[i]
        if w1 == w2 {
            continue
        }
        idx1, ok1 := index[w1]
        idx2, ok2 := index[w2]
        if !ok1 || !ok2 || find(parent, idx1) != find(parent, idx2) {
            return false
        }
    }
    return true
}

func buildUnionSet(parent map[int]int, n int) {
    for i := 0; i < n; i++ {
        parent[i] = i
    }
}

func find(parent map[int]int, target int) int {
    if parent[target] == target {
        return target
    }
    return find(parent, parent[target])
}

func union(parent map[int]int, x, y int) {
    parent[find(parent, x)] = find(parent, y)
}
