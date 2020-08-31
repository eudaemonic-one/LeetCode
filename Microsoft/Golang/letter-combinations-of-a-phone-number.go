func letterCombinations(digits string) []string {
    digitToAlpha := map[byte]string{
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }
    
    res := make([]string, 0)
    if len(digits) == 0 {
        return res
    }
    return backtracking(res, digits, digitToAlpha, 0, []rune(""))
}

func backtracking(res []string, digits string, digitToAlpha map[byte]string, idx int, letters []rune) []string {
    if idx == len(digits) {
        res = append(res, string(letters))
        return res
    }
    for _, letter := range digitToAlpha[digits[idx]] {
        letters = append(letters, letter)
        res = backtracking(res, digits, digitToAlpha, idx+1, letters)
        letters = letters[:len(letters)-1]
    }
    return res
}
