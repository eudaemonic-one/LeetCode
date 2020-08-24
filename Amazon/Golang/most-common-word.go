// Approach 2: Character Processing in One-Pass

func mostCommonWord(paragraph string, banned []string) string {
    bannedDict := make(map[string]bool)
    for _, word := range banned {
        bannedDict[word] = true
    }
    
    mostCommonWord := ""
    maxCount := 0
    wordCounter := make(map[string]int)
    sb := strings.Builder{}
    
    for i := 0; i < len(paragraph); i++ {
        ch := paragraph[i]
        if isLetter(ch) {
            sb.WriteByte(toLower(ch))
            if i != len(paragraph)-1 {
                continue
            }
        }
        
        if len(sb.String()) > 0 {
            word := sb.String()
            if _, ok := bannedDict[word]; !ok {
                wordCounter[word]++
                if wordCounter[word]> maxCount {
                    mostCommonWord = word
                    maxCount = wordCounter[word]
                }
            }
            sb.Reset()
        }
    }
    
    return mostCommonWord
}

func isLetter(ch byte) bool {
    if ('A' <= ch && ch <= 'Z') || ('a' <= ch && ch <= 'z') {
        return true
    }
    return false
}

func toLower(ch byte) byte {
    if 'A' <= ch && ch <= 'Z' {
        return ch + 'a' - 'A'
    }
    return ch
}

// Approach 1: String Processing in Pipeline

// func mostCommonWord(paragraph string, banned []string) string {
//     regex := regexp.MustCompile("[^a-zA-Z0-9 ]")
//     filteredParagraph := regex.ReplaceAllString(paragraph, " ")
//     lowercasedFilteredParagraph := strings.ToLower(filteredParagraph)
//     words := strings.Split(lowercasedFilteredParagraph, " ")
    
//     bannedDict := make(map[string]bool)
//     for _, word := range banned {
//         bannedDict[word] = true
//     }
    
//     wordCounter := make(map[string]int)
    
//     for _, word := range words {
//         if word == "" {
//             continue
//         }
//         if _, ok := bannedDict[word]; !ok {
//             wordCounter[word]++
//         }
//     }
    
//     maxCnt := 0
//     mostCommonWord := ""
    
//     for word, cnt := range wordCounter {
//         if cnt > maxCnt {
//             mostCommonWord = word
//             maxCnt = cnt
//         }
//     }
    
//     return mostCommonWord
// }
