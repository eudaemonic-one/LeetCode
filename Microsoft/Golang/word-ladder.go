type pair struct {
    word  string
    level int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
    wordDict := make(map[string][]string)
    for _, word := range wordList {
        if _, ok := wordDict[word]; !ok {
            wordDict[word] = make([]string, 0)
        }
        for i := 0; i < len(word); i++ {
            newWord := word[:i]+"*"+word[i+1:]
            wordDict[word] = append(wordDict[word], newWord)
            if _, ok := wordDict[newWord]; !ok {
                wordDict[newWord] = make([]string, 0)
            }
            wordDict[newWord] = append(wordDict[newWord], word)
        }
    }
    
    queue := make([]pair, 0)
    queue = append(queue, pair{beginWord, 1})
    
    visited := make(map[string]bool)
    visited[beginWord] = true
    
    for len(queue) > 0 {
        head := queue[0]
        queue = queue[1:]
        word, level := head.word, head.level
        for i := 0; i < len(word); i++ {
            newWord := word[:i] + "*" + word[i+1:]
            for _, adjWord := range wordDict[newWord] {
                if adjWord == endWord {
                    return level + 1
                }
                if _, ok := visited[adjWord]; !ok {
                    visited[adjWord] = true
                    queue = append(queue, pair{adjWord, level+1})
                }
            }
        }
    }
    
    return 0
}
