func calculate(s string) int {
    l1, o1 := 0, 1
    l2, o2 := 1, 1
    stack := make([]int, 0)
    for i := 0; i < len(s); i++ {
        switch ch := s[i]; ch {
        case '(':
            stack = append(stack, l1)
            stack = append(stack, o1)
            stack = append(stack, l2)
            stack = append(stack, o2)
            l1, o1 = 0, 1
            l2, o2 = 1, 1
        case ')':
            num := l1 + o1*l2
            o2 = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            l2 = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            o1 = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            l1 = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if o2 == 1 {
                l2 = l2 * num
            } else {
                l2 = l2 / num
            }
        case '*':
            o2 = 1
        case '/':
            o2 = -1
        case '+':
            l1 = l1 + o1*l2
            o1 = 1
            l2, o2 = 1, 1
        case '-':
            l1 = l1 + o1*l2
            o1 = -1
            l2, o2 = 1, 1
        default:
            if isDigit(ch) { 
                num := int(ch-'0')
                for i+1 < len(s) && isDigit(s[i+1]) {
                    num = 10*num + int(s[i+1]-'0')
                    i++
                }
                if o2 == 1 {
                    l2 = l2 * num
                } else {
                    l2 = l2 / num
                }
            }
        }
    }
    return l1 + o1*l2
}

func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}
