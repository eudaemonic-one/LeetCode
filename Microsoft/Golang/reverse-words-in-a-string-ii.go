func reverseWords(s []byte)  {
    reverse(s, 0, len(s)-1)
    reverseEachWord(s)
}

func reverse(s []byte, l, r int) {
    for l < r {
        s[l], s[r] = s[r], s[l]
        l++
        r--
    }
}

func reverseEachWord(s []byte) {
    n := len(s)
    l, r := 0, 0
    for l < n {
        for r < n && s[r] != ' ' {
            r++
        }
        reverse(s, l, r-1)
        l = r + 1
        r = l
    }
}
