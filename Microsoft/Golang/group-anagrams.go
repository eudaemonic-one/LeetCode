func groupAnagrams(strs []string) [][]string {
    table := make(map[string][]string)
    for _, word := range strs {
        count := make([]int, 26)
        for i := 0; i < len(word); i++ {
            count[int(word[i]-'a')]++
        }
        bytes := make([]byte, 0)
        for i := 0; i < 26; i++ {
            bytes = append(bytes, '#')
            if count[i] > 0 {
                bytes = append(bytes, []byte(strconv.Itoa(count[i]))...)
            }
        }
        str := string(bytes)
        if _, ok := table[str]; !ok {
            table[str] = []string{word}
        } else {
            table[str] = append(table[str], word)
        }
    }
    res := make([][]string, 0)
    for _, anagrams := range table {
        res = append(res, anagrams)
    }
    return res
}
