func findWords(board [][]byte, words []string) []string {
    res := make([]string, 0)
    root := buildTrie(words)
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            dfs(board, root, i, j, &res)
        }
    }
    return res
}

func dfs(board [][]byte, node *TrieNode, i, j int, res *[]string) {
    if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
        return
    }
    
    ch := board[i][j]
    if ch == '#' {
        return
    }
    
    idx := int(ch) - int('a')
    if child := node.Children[idx]; child != nil {
        if child.Word != "" {
            *res = append(*res, child.Word)
            child.Word = ""
        }
        
        board[i][j] = '#'
        dfs(board, child, i, j-1, res)
        dfs(board, child, i, j+1, res)
        dfs(board, child, i-1, j, res)
        dfs(board, child, i+1, j, res)
        board[i][j] = ch
    }
}

const N int = 26

type TrieNode struct {
    Children [N]*TrieNode
    Word string
}

func buildTrie(words []string) *TrieNode {
    root := &TrieNode{}
    for _, word := range words {
        curr := root
        for _, ch := range word {
            i := int(ch - 'a')
            if curr.Children[i] == nil {
                curr.Children[i] = &TrieNode{}
            }
            curr = curr.Children[i]
        }
        curr.Word = word
    }
    return root
}
