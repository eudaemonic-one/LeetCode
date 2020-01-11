/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sumEvenGrandparent(root *TreeNode) int {
    sum := 0
    helper(root, nil, nil, &sum)
    return sum
}

func helper(root, father, grandparent *TreeNode, sum *int) {
    if root == nil {
        return
    }
    if grandparent != nil && grandparent.Val % 2 == 0 {
        *sum += root.Val
    }
    newFather := root
    newGrandparent := father
    helper(root.Left, newFather, newGrandparent, sum)
    helper(root.Right, newFather, newGrandparent, sum)
}