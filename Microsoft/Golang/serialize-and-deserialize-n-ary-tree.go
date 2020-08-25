/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Codec struct {
    sb strings.Builder
}

func Constructor() *Codec {
    return &Codec{strings.Builder{}}
}

func (this *Codec) serialize(root *Node) string {
    this.sb.Reset()
    this.serializeHelper(root)
    return this.sb.String()
}

func (this *Codec) serializeHelper(root *Node) {
    if root == nil {
        return
    }
    this.sb.WriteString(fmt.Sprintf("%d,%d,", root.Val, len(root.Children)))
    for _, child := range root.Children {
        this.serializeHelper(child)
    }
}

func (this *Codec) deserialize(data string) *Node {
    if data == "" {
        return nil
    }
    tokens := strings.Split(data, ",")
    idx := 0
    return deserializeHelper(tokens, &idx)
}

func deserializeHelper(tokens []string, idx *int) *Node {
    if *idx == len(tokens) {
        return nil
    }
    nodeVal, _ := strconv.Atoi(tokens[*idx])
    node := &Node{nodeVal, make([]*Node, 0)}
    (*idx)++
    childrenNum, _ := strconv.Atoi(tokens[*idx])
    for i := 0; i < childrenNum; i++ {
        (*idx)++
        node.Children = append(node.Children, deserializeHelper(tokens, idx))
    }
    return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
