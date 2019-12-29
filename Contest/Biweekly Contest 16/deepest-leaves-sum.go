/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func deepestLeavesSum(root *TreeNode) int {
    sum := 0
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        n := len(queue)
        tmpSum := 0
        for i := 0; i < n; i++ {
            node := queue[i]
            tmpSum += node.Val
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        queue = queue[n:]
        sum = tmpSum
    }
    return sum
}