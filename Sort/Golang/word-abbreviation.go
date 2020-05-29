func wordsAbbreviation(dict []string) []string {
    res := make([]string, len(dict))
    prefix := make([]int, len(dict))
    for i, word := range dict {
        res[i] = abbreviate(word, 0)
    }
    for i := 0; i < len(res); i++ {
        for true {
            duplicates := make(map[int]bool)
            for j := i+1; j < len(res); j++ {
                if res[i] == res[j] {
                    duplicates[j] = true
                }
            }
            if len(duplicates) == 0 {
                break
            }
            duplicates[i] = true
            for k, _ := range duplicates {
                prefix[k] += 1
                res[k] = abbreviate(dict[k], prefix[k])
            }
        }
    }
    return res
}

func abbreviate(word string, i int) string {
    if len(word) - i <= 3 {
        return word
    }
    return word[:i+1] + strconv.Itoa(len(word)-i-2) + string(word[len(word)-1])
}
