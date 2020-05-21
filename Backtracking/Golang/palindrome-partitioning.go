func partition(s string) [][]string {
    res := make([][]string, 0)
    path := make([]string, 0)
    backtrack(s, &path, &res)
    return res
}

func backtrack(s string, path *[]string, res *[][]string) {
    if len(s) == 0 {
        tmp := make([]string, len(*path))
        copy(tmp, *path)
        *res = append(*res, tmp)
        return
    }
    for i := 0; i < len(s); i++ {
        if (isPalindrome(s[:i+1])) {
            *path = append(*path, s[:i+1])
            backtrack(s[i+1:], path, res)
            *path = (*path)[:len(*path)-1]
        }
    }
}

func isPalindrome(s string) bool {
    i, j := 0, len(s)-1
    for i < j {
        if (s[i] != s[j]) {
            return false
        }
        i += 1
        j -= 1
    }
    return true
}
