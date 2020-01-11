type Pair struct {
    Key byte
    Val int
}

func reorganizeString(S string) string {
    counter := make(map[byte]int)
    for i := 0; i < len(S); i++ {
        counter[S[i]]++
        if counter[S[i]] > (len(S)+1)/2 {
            return ""
        }
    }
    pairs := make([]Pair, 0)
    for k, v := range counter {
        pairs = append(pairs, Pair{k,v})
    }
    sort.Slice(pairs, func (i, j int) bool {
        return pairs[i].Val >= pairs[j].Val
    })
    res := make([]byte, len(S))
    j := 0
    for _, pair := range pairs {
        for i := 0; i < pair.Val; i, j = i+1, j+2 {
            if j >= len(S) {
                j = 1
            }
            res[j] = pair.Key
        }        
    }
    return string(res)
}
