func licenseKeyFormatting(S string, K int) string {
    res := make([]byte, 0)
    for i :=  len(S)-1; i >= 0; i-- {
        if S[i] != '-' {
            if len(res)%(K+1) == K {
                res = append([]byte{'-'}, res...)
            }
            res = append([]byte{upper(S[i])}, res...)
        }
    }
    return string(res)
}

func upper(ch byte) byte {
    if 'a' <= ch && ch <= 'z' {
        return ch-32
    }
    return ch
}
