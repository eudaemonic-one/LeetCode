func decodeString(s string) string {
    i := 0
    kStack := make([]int, 0)
    encodedStack := make([]string, 0)
    decoded := ""
    for i < len(s) {
        if isDigit(s[i]) {
            k := 0
            for isDigit(s[i]) {
                k = k*10 + int(s[i]-'0')
                i++
            }
            kStack = append(kStack, k)
        } else if s[i] == '[' {
            encodedStack = append(encodedStack, decoded)
            decoded = ""
            i++
        } else if s[i] == ']' {
            temp := encodedStack[len(encodedStack)-1]
            encodedStack = encodedStack[:len(encodedStack)-1]
            k := kStack[len(kStack)-1]
            kStack = kStack[:len(kStack)-1]
            for j := 0; j < k; j++ {
                temp += decoded
            }
            decoded = temp
            i++
        } else {
            decoded += string(s[i])
            i++
        }
    }
    return decoded
}

func isDigit(ch byte) bool {
    if '0' <= ch && ch <= '9' {
        return true
    }
    return false
}
