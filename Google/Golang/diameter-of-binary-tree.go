/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    diameter := 0
    helper(root, &diameter)
    return diameter
}

func helper(root *TreeNode, diameter *int) int {
    if root == nil {
        return 0
    }
    left := helper(root.Left, diameter)
    right := helper(root.Right, diameter)
    path := left + right
    *diameter = max(*diameter, path)
    return max(left, right) + 1
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
