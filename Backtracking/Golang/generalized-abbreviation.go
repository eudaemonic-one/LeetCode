func generateAbbreviations(word string) []string {
    res := make([]string, 0)
    backtrack(word, 0, 0, &[]byte{}, &res)
    return res
}

func backtrack(word string, idx, abbrCnt int, path *[]byte, res *[]string) {
    l := len(*path)
    if idx == len(word) {
        if abbrCnt > 0 {    // append the last abbreviated characters
            *path = append(*path, strconv.Itoa(abbrCnt)...)
        }
        tmp := make([]byte, len(*path))
        copy(tmp, *path)
        *res = append(*res, string(tmp))
    } else {
        // the branch that word[i] is abbreviated
        backtrack(word, idx+1, abbrCnt+1, path, res)
        
        // the branch that word[i] is kept
        if abbrCnt > 0 {
            *path = append(*path, strconv.Itoa(abbrCnt)...)
        }
        *path = append(*path, word[idx])
        backtrack(word, idx+1, 0, path, res)
    }
    *path = (*path)[:l]
}
