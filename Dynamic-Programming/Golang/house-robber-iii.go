/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Approach #2: Recursion with Tricks

func rob(root *TreeNode) int {
    notRob, rob := helper(root)
    return max(notRob, rob)
}

func helper(root *TreeNode) (int, int) {
    if root == nil {
        return 0, 0
    }
    leftNotRob, leftRob := helper(root.Left)
    rightNotRob, rightRob := helper(root.Right)
    rob := root.Val + leftNotRob + rightNotRob
    notRob := max(leftNotRob, leftRob) + max(rightNotRob, rightRob)
    return notRob, rob
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

// Approach #1: Recursion with Memo

// func rob(root *TreeNode) int {
//     memo := make(map[*TreeNode]int)
//     return helper(memo, root)
// }

// func helper(memo map[*TreeNode]int, root *TreeNode) int {
//     if root == nil {
//         return 0
//     }
//     if v, ok := memo[root]; ok {
//         return v
//     }
//     rob := root.Val
//     if root.Left != nil {
//         rob += helper(memo, root.Left.Left) + helper(memo, root.Left.Right)
//     }
//     if root.Right != nil {
//         rob += helper(memo, root.Right.Left) + helper(memo, root.Right.Right)
//     }
//     notRob := helper(memo, root.Left) + helper(memo, root.Right)
//     res := max(rob, notRob)
//     memo[root] = res
//     return res
// }

// func max(x, y int) int {
//     if x > y {
//         return x
//     }
//     return y
// }
