func longestValidParentheses(s string) int {
    maxans := 0
    left, right := 0, 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            left++
        } else {
            right++
        }
        if left == right {
            maxans = max(maxans, left+right)
        } else if right > left {
            left, right = 0, 0
        }
    }
    left, right = 0, 0
    for i := len(s)-1; i >= 0; i-- {
        if s[i] == '(' {
            left++
        } else {
            right++
        }
        if left == right {
            maxans = max(maxans, left+right)
        } else if left > right {
            left, right = 0, 0
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
