/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

func inorderSuccessor(node *Node) *Node {
    if node == nil {
        return nil
    }
    if node.Right != nil {
        curr := node.Right
        for curr.Left != nil {
            curr = curr.Left
        }
        return curr
    }
    curr := node
    for curr.Parent != nil && curr == curr.Parent.Right {
        curr = curr.Parent
    }
    return curr.Parent
}
