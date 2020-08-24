func partitionLabels(S string) []int {
    res := make([]int, 0)
    last := make([]int, 26)
    for i := 0; i < len(S); i++ {
        last[int(S[i]-'a')] = i
    }
    idx := 0
    for idx < len(S) {
        lastIdx := last[int(S[idx]-'a')]
        for j := idx+1; j < lastIdx; j++ {
            if last[int(S[j]-'a')] > lastIdx {
                lastIdx = last[int(S[j]-'a')]
            }
        }
        res = append(res, lastIdx-idx+1)
        idx = lastIdx + 1
    }
    return res
}
