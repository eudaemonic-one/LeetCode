type TrieNode struct {
    children [26]*TrieNode
    word     string
}

func buildTrie(words []string) *TrieNode {
    root := &TrieNode{}
    for _, word := range words {
        curr := root
        for i := 0; i < len(word); i++ {
            idx := int(word[i]-'a')
            if curr.children[idx] == nil {
                curr.children[idx] = &TrieNode{}
            }
            curr = curr.children[idx]
        }
        curr.word = word
    }
    return root
}

func findWords(board [][]byte, words []string) []string {
    if len(board) == 0 || len(board[0]) == 0 {
        return []string{}
    }
    
    root := buildTrie(words)
    res := make([]string, 0)
    
    for r := 0; r < len(board); r++ {
        for c := 0; c < len(board[0]); c++ {
            dfs(&res, root, board, r, c)
        }
    }
    
    return res
}

func dfs(res *[]string, node *TrieNode, board [][]byte, r, c int) {
    if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) {
        return
    }
    
    ch := board[r][c]
    if ch == '#' {
        return
    }
    
    idx := int(ch-'a')
    if child := node.children[idx]; child != nil {
        if child.word != "" {
            *res = append(*res, child.word)
            child.word = ""
        }
        board[r][c] = '#'
        dfs(res, child, board, r, c-1)
        dfs(res, child, board, r, c+1)
        dfs(res, child, board, r+1, c)
        dfs(res, child, board, r-1, c)
        board[r][c] = ch
    }
}
