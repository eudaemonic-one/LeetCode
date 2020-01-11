type Trie struct {
    Children map[rune]*Trie
    IsEnd bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{make(map[rune]*Trie), false}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    cur := this
    for _, ch := range word {
        if _, ok := cur.Children[ch]; !ok {
            cur.Children[ch] = &Trie{make(map[rune]*Trie), false}
        }
        cur = cur.Children[ch]
    }
    cur.IsEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    cur := this
    for _, ch := range word {
        if _, ok := cur.Children[ch]; !ok {
            return false
        }
        cur = cur.Children[ch]
    }
    if !cur.IsEnd {
        return false
    }
    return true
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    cur := this
    for _, ch := range prefix {
        if _, ok := cur.Children[ch]; !ok {
            return false
        }
        cur = cur.Children[ch]
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
