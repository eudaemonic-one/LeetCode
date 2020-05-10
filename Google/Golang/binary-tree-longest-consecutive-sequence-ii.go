/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestConsecutive(root *TreeNode) int {
    if root == nil {
        return 0
    }
    res := 0
    helper(root, &res)
    return res
}

func helper(root *TreeNode, res *int) (int, int) {
    if root == nil {
        return 0, 0
    }
    inc, dec := 1, 1
    if root.Left != nil {
        leftInc, leftDec := helper(root.Left, res)
        if root.Val == root.Left.Val + 1 {
            dec = leftDec + 1
        } else if root.Val == root.Left.Val -1 {
            inc = leftInc + 1
        }
    }
    if root.Right != nil {
        rightInc, rightDec := helper(root.Right, res)
        if root.Val == root.Right.Val + 1 {
            dec = max(dec, rightDec + 1)
        } else if root.Val == root.Right.Val -1 {
            inc = max(inc, rightInc + 1)
        }
    }
    *res = max(*res, dec + inc - 1)
    return inc, dec
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
