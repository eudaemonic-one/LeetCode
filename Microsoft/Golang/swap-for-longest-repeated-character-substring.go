type pair struct {
    ch byte
    cnt int
}

func maxRepOpt1(text string) int {
    if len(text) == 0 {
        return 0
    }
    res := 1
    counter := make(map[byte]int)
    groups := make([]pair, 0)
    lastCh := byte(0)
    for i := 0; i < len(text); i++ {
        ch := text[i]
        counter[ch]++
        if ch != lastCh {
            groups = append(groups, pair{ch, 1})
            lastCh = ch
        } else {
            groups[len(groups)-1].cnt++
        }
    }
    
    for i := 0; i < len(groups); i++ {
        ch, cnt := groups[i].ch, groups[i].cnt
        res = max(res, min(cnt+1, counter[ch]))
        if i != 0 && i != len(groups)-1 {
            if cnt == 1 && groups[i-1].ch == groups[i+1].ch {
                res = max(res, min(groups[i-1].cnt+groups[i+1].cnt+1, counter[groups[i+1].ch]))
            }
        }
    }
    
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
