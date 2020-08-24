import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    sb strings.Builder
}

func Constructor() Codec {
    return Codec{strings.Builder{}}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    this.sb.Reset()
    serializeHelper(root, &this.sb)
    return this.sb.String()
}

func serializeHelper(root *TreeNode, sb *strings.Builder) {
    if root == nil {
        sb.WriteString("null,")
    } else {
        sb.WriteString(strconv.Itoa(root.Val)+",")
        serializeHelper(root.Left, sb)
        serializeHelper(root.Right, sb)
    }
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    nodeStrs := strings.Split(data, ",")
    return deserializeHelper(&nodeStrs)
}

func deserializeHelper(nodeStrs *[]string) *TreeNode {
    if len(*nodeStrs) == 0 {
        return nil
    }
    if (*nodeStrs)[0] == "null" {
        (*nodeStrs) = (*nodeStrs)[1:]
        return nil
    }
    val, _ := strconv.Atoi((*nodeStrs)[0])
    (*nodeStrs) = (*nodeStrs)[1:]
    root := &TreeNode{val, nil, nil}
    root.Left = deserializeHelper(nodeStrs)
    root.Right = deserializeHelper(nodeStrs)
    return root
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
