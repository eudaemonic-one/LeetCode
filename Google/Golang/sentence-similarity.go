func areSentencesSimilar(words1 []string, words2 []string, pairs [][]string) bool {
    if len(words1) != len(words2) {
        return false
    }
    pairset := make(map[[2]string]bool)
    for _, pair := range pairs {
        w1, w2 := pair[0], pair[1]
        pairset[[2]string{w1,w2}] = true
    }
    for i := 0; i < len(words2); i++ {
        w1, w2 := words1[i], words2[i]
        _, ok1 := pairset[[2]string{w1,w2}]
        _, ok2 := pairset[[2]string{w2,w1}]
        if w1 != w2 && !ok1 && !ok2 {
            return false
        }
    }
    return true
}
