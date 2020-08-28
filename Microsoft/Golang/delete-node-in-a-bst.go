/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }
    if key < root.Val {
        root.Left = deleteNode(root.Left, key)
    } else if key > root.Val {
        root.Right = deleteNode(root.Right, key)
    } else {
        if root.Left == nil && root.Right == nil {
            root = nil
        } else if root.Right != nil {
            root.Val = successor(root).Val
            root.Right = deleteNode(root.Right, root.Val)
        } else {
            root.Val = predecessor(root).Val
            root.Left = deleteNode(root.Left, root.Val)
        }
    }
    return root
}

func successor(root *TreeNode) *TreeNode {
    root = root.Right
    for root.Left != nil {
        root = root.Left
    }
    return root
}

func predecessor(root *TreeNode) *TreeNode {
    root = root.Left
    for root.Right != nil {
        root = root.Right
    }
    return root
}
