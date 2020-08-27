/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return validateBST(root, math.MinInt64, math.MaxInt64)
}

func validateBST(node *TreeNode, min, max int) bool {
    if node == nil {
        return true
    }
    val := node.Val
    if val <= min || val >= max {
        return false
    }
    return validateBST(node.Left, min, val) && validateBST(node.Right, val, max)
}
