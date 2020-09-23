/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
    return s != nil && (equals(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t))
}

func equals(x, y *TreeNode) bool {
    if x == nil && y == nil {
        return true
    }
    if x == nil || y == nil {
        return false
    }
    return x.Val == y.Val && equals(x.Left, y.Left) && equals(x.Right, y.Right)
}
