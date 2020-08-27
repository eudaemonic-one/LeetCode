/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {}

func Constructor() *Codec {
    return &Codec{}
}

func (this *Codec) encode(root *Node) *TreeNode {
    if root == nil {
        return nil
    }
    
    newRoot := &TreeNode{root.Val, nil, nil}
    nodeQueue := []*Node{root}
    treeNodeQueue := []*TreeNode{newRoot}
    
    for len(nodeQueue) > 0 {
        treeNode := treeNodeQueue[0]
        treeNodeQueue = treeNodeQueue[1:]
        node := nodeQueue[0]
        nodeQueue = nodeQueue[1:]
        
        var prevTreeNode, headTreeNode *TreeNode
        for _, childNode := range node.Children {
            newTreeNode := &TreeNode{childNode.Val, nil, nil}
            if prevTreeNode == nil {
                headTreeNode = newTreeNode
            } else {
                prevTreeNode.Right = newTreeNode
            }
            prevTreeNode = newTreeNode
            
            nodeQueue = append(nodeQueue, childNode)
            treeNodeQueue = append(treeNodeQueue, newTreeNode)
        }
        treeNode.Left = headTreeNode
    }
    
    return newRoot
}

func (this *Codec) decode(root *TreeNode) *Node {
    if root == nil {
        return nil
    }
    newRoot := &Node{root.Val, make([]*Node, 0)}
    treeNodeQueue := []*TreeNode{root}
    nodeQueue := []*Node{newRoot}
    
    for len(treeNodeQueue) > 0 {
        treeNode := treeNodeQueue[0]
        treeNodeQueue = treeNodeQueue[1:]
        node := nodeQueue[0]
        nodeQueue = nodeQueue[1:]
        
        firstChild := treeNode.Left
        sibling := firstChild
        for sibling != nil {
            childNode := &Node{sibling.Val, make([]*Node, 0)}
            node.Children = append(node.Children, childNode)
            
            nodeQueue = append(nodeQueue, childNode)
            treeNodeQueue = append(treeNodeQueue, sibling)
            
            sibling = sibling.Right
        }
    }
    
    return newRoot
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * bst := obj.encode(root);
 * ans := obj.decode(bst);
 */
