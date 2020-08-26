func restoreIpAddresses(s string) []string {
    return dfs(make([]string, 0), s, []string{}, 0)
}

func dfs(res []string, s string, path []string, idx int) []string {
    if idx == len(s) || len(path) > 4 {
        if len(path) == 4 {
            res = append(res, strings.Join(path, "."))
        }
        return res
    }
    for i := idx; i < min(len(s), idx+3); i++ {
        subStr := s[idx:i+1]
        if len(subStr) > 1 && subStr[0] == '0' {
            continue
        }
        digits, _ := strconv.Atoi(subStr)
        if 0 <= digits && digits <= 255 {
            path = append(path, subStr)
            res = dfs(res, s, path, i+1)
            path = path[:len(path)-1]
        }
    }
    return res
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
