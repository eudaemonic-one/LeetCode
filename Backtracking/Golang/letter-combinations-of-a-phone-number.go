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
    backtracking(digits, digitToAlpha, "", &res)
    return res
}

func backtracking(digits string, digitToAlpha map[byte]string, letters string, res *[]string) {
    if len(digits) == 0 {
        *res = append(*res, letters)
        return
    }
    for _, letter := range digitToAlpha[digits[0]] {
        backtracking(digits[1:], digitToAlpha, letters + string(letter), res)
    }
}
