func licenseKeyFormatting(S string, K int) string {
    withoutDashes := make([]byte, 0)
    for i := 0; i < len(S); i++ {
        if S[i] != '-' {
            withoutDashes = append(withoutDashes, upper(S[i]))
        }
    }
    j := len(withoutDashes)
    res := make([]byte, 0)
    for j-K > 0 {
        if j != len(withoutDashes) {
            res = append([]byte{'-'}, res...)
        }
        res = append(withoutDashes[j-K:j], res...)
        j -= K
    }
    if j > 0 {
        if j != len(withoutDashes) {
            res = append([]byte{'-'}, res...)
        }
        res = append(withoutDashes[:j], res...)
    }
    return string(res)
}

func upper(ch byte) byte {
    if 'a' <= ch && ch <= 'z' {
        return ch-32
    }
    return ch
}
