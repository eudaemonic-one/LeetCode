/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    res := ^int(^uint32(0) >> 1)
    helper(root, &res)
    return res
}

func helper(root *TreeNode, res *int) int {
    if root == nil {
        return 0
    }
    left := helper(root.Left, res)
    right := helper(root.Right, res)
    pathSum := max(max(root.Val,left+root.Val+right), max(left+root.Val, right+root.Val))
    if pathSum > *res {
        *res = pathSum
    }
    return max(root.Val, max(left+root.Val, right+root.Val))
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
