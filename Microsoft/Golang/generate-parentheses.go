func generateParenthesis(n int) []string {
    res := make([]string, 0)
    return backtracking(res, n, 0, "")
}

func backtracking(res []string, n, leftCount int, parentheses string) []string {
    if n == 0 {
        res = append(res, parentheses)
        return res
    }
    if leftCount == 0 {
        res = backtracking(res, n, leftCount+1, parentheses + "(")
    } else {
        if leftCount < n {
            res = backtracking(res, n, leftCount+1, parentheses + "(")
        }
        res = backtracking(res, n-1, leftCount-1, parentheses + ")")
    }
    return res
}
