/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    res := root
    helper(root, p, q, &res)
    return res
}

func helper(root, p, q *TreeNode, res **TreeNode) bool {
    if root == nil {
        return false
    }
    flag := root == p || root == q
    left := helper(root.Left, p, q, res)
    right := helper(root.Right, p, q, res)
    if (left && right) || (flag && left) || (flag && right) {
        *res = root
    }
    return left || right || flag
}
