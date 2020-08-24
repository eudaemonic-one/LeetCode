// Approach 2: DFS with Trie

type TrieNode struct {
    children [26]*TrieNode
    isEnd    bool
}

func buildTrie(words []string) *TrieNode {
    if len(words) == 0 {
        return nil
    }
    root := &TrieNode{}
    for _, word := range words {
        if word == "" {
            continue
        }
        curr := root
        for i := 0; i < len(word); i++ {
            idx := int(word[i]-'a')
            if curr.children[idx] == nil {
                curr.children[idx] = &TrieNode{}
            }
            curr = curr.children[idx]
        }
        curr.isEnd = true
    }
    return root
}

func findAllConcatenatedWordsInADict(words []string) []string {
    root := buildTrie(words)
    
    concatenatedWords := make([]string, 0)
    for _, word := range words {
        if word != "" && isConcatenatedWords(root, word, 0, 0) {
            concatenatedWords = append(concatenatedWords, word)
        }
    }
    return concatenatedWords
}

func isConcatenatedWords(root *TrieNode, word string, idx int, count int) bool {
    curr := root
    for i := idx; i < len(word); i++ {
        ch := word[i]
        if curr.children[int(ch-'a')] == nil {
            return false
        }
        curr = curr.children[int(ch-'a')]
        if curr.isEnd {
            if i == len(word)-1 {
                return count > 0
            }
            if isConcatenatedWords(root, word, i+1, count+1) {
                return true
            }
        }
    }
    return false
}

// Approach 1: Dynamic Programming

// func findAllConcatenatedWordsInADict(words []string) []string {
//     sort.Slice(words, func (i, j int) bool {
//         return len(words[i]) < len(words[j])
//     })
    
//     wordDict := make(map[string]bool)
//     concatenatedWords := make([]string, 0)
//     for _, word := range words {
//         if word != "" && isConcatenatedWords(wordDict, word) {
//             concatenatedWords = append(concatenatedWords, word)
//         }
//         wordDict[word] = true
//     }
//     return concatenatedWords
// }

// func isConcatenatedWords(wordDict map[string]bool, word string) bool {
//     dp := make([]bool, len(word))
//     for j := 0; j < len(word); j++ {
//         for i := 0; i <= j; i++ {
//             if _, ok := wordDict[word[i:j+1]]; ok && (i == 0 || dp[i-1]) && !(i == 0 && j == len(word)-1) {
//                 dp[j] = true
//             }
//         }
//     }
//     return dp[len(word)-1]
// }
