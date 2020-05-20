func generateParenthesis(n int) []string {
    res := make([]string, 0)
    backtracking(n, 0, "", &res)
    return res
}

func backtracking(n, leftCnt int, parentheses string, res *[]string) {
    if n == 0 {
        *res = append(*res, parentheses)
        return
    }
    if leftCnt == 0 {
        backtracking(n, leftCnt+1, parentheses + "(", res)      // Append '('
    } else {
        if n > leftCnt {
            backtracking(n, leftCnt+1, parentheses + "(", res)  // Append '('
        }
        backtracking(n-1, leftCnt-1, parentheses + ")", res)    // Append ')'
    }
}
