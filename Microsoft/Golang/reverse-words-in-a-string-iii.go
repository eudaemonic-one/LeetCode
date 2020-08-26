func reverseWords(s string) string {
    bytes := []byte(s)
    l := 0
    for r := 0; r < len(s); r++ {
        if s[r] == ' ' {
            if s[l:r] != "" {
                reverseBytes(bytes, l, r-1)
            }
            l = r+1
        }
    }
    if s[l:] != "" {
        reverseBytes(bytes, l, len(bytes)-1)
    }
    return string(bytes)
}

func reverseBytes(bytes []byte, l, r int) {
    for i, j := l, r; l < r; l, r = l+1, r-1 {
        bytes[i], bytes[j] = bytes[j], bytes[i]
        i++
        j--
    }
}
