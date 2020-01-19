func printVertically(s string) []string {
    res := make([]string, 0)
    words := make([]string, 0)
    maxlen := 0
    word := ""
    for i := 0; i < len(s); i++ {
        if s[i] == ' ' {
            words = append(words, word)
            word = ""
        } else {
            word += string(s[i])
            if len(word) > maxlen {
                maxlen = len(word)
            }
        }
    }
    if len(word) > 0 {
        words = append(words, word)
    }
    for i := 0; i < maxlen; i++ {
        vword := ""
        for j := 0; j < len(words); j++ {
            if i >= len(words[j]) {
                vword += " "
            } else {
                vword += string(words[j][i])
            }
        }
        vword = strings.TrimRight(vword, " ")
        res = append(res, vword)
    }
    return res
}