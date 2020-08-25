func reverseWords(s string) string {
    words := make([]string, 0)
    
    l, r := 0, len(s)-1
    for l <= r && s[l] == ' ' {
        l++
    }
    for l <= r && s[r] == ' ' {
        r--
    }
    
    word := ""
    for l <= r {
        if s[l] != ' ' {
            word += string(s[l])
        } else if word != "" {
            words = append(words, word)
            word = ""
        }
        l++
    }
    if word != "" {
        words = append(words, word)
    }
    
    for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
        words[i], words[j] = words[j], words[i]
    }
    
    return strings.Join(words, " ")
}
