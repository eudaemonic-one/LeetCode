func longestValidParentheses(s string) int {
    maxans := 0
    stack := []int{-1}
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            stack = append(stack, i)
        } else {
            stack = stack[:len(stack)-1]
            if len(stack) == 0 {
                stack = append(stack, i)
            } else {
                maxans = max(maxans, i-stack[len(stack)-1])
            }
        }
    }
    return maxans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
