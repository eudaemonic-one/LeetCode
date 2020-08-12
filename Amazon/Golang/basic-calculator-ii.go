func calculate(s string) int {
    strings.TrimSpace(s)
    if len(s) == 0 {
        return 0
    }
    
    var stack []int
    var num int
    var sign byte = '+'
    
    for i := 0; i < len(s); i++ {
        if '0' <= s[i] && s[i] <= '9' {
            num = 10*num + int(s[i]-'0')
        }
        if ((s[i] < '0' || s[i] > '9') && s[i] != ' ') || i == len(s)-1 {
            switch sign {
            case '+':
                stack = append(stack, num)
            case '-':
                stack = append(stack, -num)
            case '*':
                top := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                stack = append(stack, num*top)
            case '/':
                top := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                stack = append(stack, top/num)
            }
            sign = s[i]
            num = 0
        }
    }
    
    res := 0
    for _, num := range stack {
        res += num
    }
    return res
}
