// Approach 3: Binary Search

func suggestedProducts(products []string, searchWord string) [][]string {
    sort.Strings(products)

    res := make([][]string, 0)
    
    for i := 0; i < len(searchWord); i++ {
        idx := bisectLeft(products, searchWord[:i+1])
        tmpRes := make([]string, 0)
        rangeLeft, rangeRight := idx, min(len(products), idx+3)
        if rangeLeft >= len(products) {
            res = append(res, tmpRes)
            continue
        }
        for _, word := range products[rangeLeft:rangeRight] {
            if strings.HasPrefix(word, searchWord[:i+1]) {
                tmpRes = append(tmpRes, word)
            }
        }
        res = append(res, tmpRes)
    }
    
    return res
}

func bisectLeft(words []string, target string) int {
    l, r := 0, len(words)
    for l < r {
        m := l + (r-l)/2
        if target > words[m] {
            l = m + 1
        } else {
            r = m
        }
    }
    return l
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

// Approach 2: Trie

// type Trie struct {
//     children [26]*Trie
//     suggestionWords []string
// }

// func buildTrie(words []string) *Trie {
//     root := &Trie{}
//     for _, word := range words {
//         curr := root
//         for i := 0; i < len(word); i++ {
//             idx := int(word[i]-'a')
//             if curr.children[idx] == nil {
//                 curr.children[idx] = &Trie{}
//             }
//             curr = curr.children[idx]
//             if len(curr.suggestionWords) < 3 {
//                 curr.suggestionWords = append(curr.suggestionWords, word)
//             }
//         }
//     }
//     return root
// }

// func suggestedProducts(products []string, searchWord string) [][]string {
//     sort.Strings(products)
    
//     root := buildTrie(products)
    
//     res := make([][]string, 0)
//     for i := 0; i < len(searchWord); i++ {
//         if root == nil {
//             res = append(res, make([]string, 0))
//             continue
//         }
//         root = root.children[int(searchWord[i]-'a')]
//         if root != nil {
//             res = append(res, root.suggestionWords)
//         } else {
//             res = append(res, make([]string, 0))
//         }
//     }
//     return res
// }

// Approach 1: Brute Force

// func suggestedProducts(products []string, searchWord string) [][]string {
//     res := make([][]string, len(searchWord))
    
//     for i := 0; i <len(searchWord); i++ {
//         matchWords := make([]string, 0)
//         for _, product := range products {
//             if i < len(product) && searchWord[:i+1] == product[:i+1] {
//                 matchWords = append(matchWords, product)
//             }
//         }
//         sort.Sort(sort.StringSlice(matchWords))
//         if len(matchWords) > 3 {
//             matchWords = matchWords[0:3]
//         }
//         res[i] = matchWords
//     }
    
//     return res
// }
