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

func validateBST(node *TreeNode, lower, upper int) bool {
    if node == nil {
        return true
    }
    
    val := node.Val
    if val <= lower {
        return false
    }
    if val >= upper {
        return false
    }
    
    if !validateBST(node.Right, val, upper) {
        return false
    }
    if !validateBST(node.Left, lower, val) {
        return false
    }
    
    return true
}
