/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)
    helper(root, &res)
    return res
}

func helper(root *TreeNode, res *[]int) {
    if (root == nil) {
        return
    }
    helper(root.Left, res)
    helper(root.Right, res)
    *res = append(*res, root.Val)
}
