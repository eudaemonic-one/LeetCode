func multiply(num1 string, num2 string) string {
    m, n := len(num1), len(num2)
    if m == 0 || n == 0 {
        return "0"
    }
    digits := make([]int, m+n)
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            prod := int(num1[i]-'0') * int(num2[j]-'0')
            sum := digits[i+j+1] + prod
            digits[i+j] += sum / 10
            digits[i+j+1] = sum % 10
        }
    }
    i := 0
    for i < m+n && digits[i] == 0 {
        i++
    }
    res := make([]byte, 0)
    for i < m+n {
        res = append(res, byte('0'+digits[i]))
        i++
    }
    if len(res) == 0 {
        return "0"
    }
    return string(res)
}
