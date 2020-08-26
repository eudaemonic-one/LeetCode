
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    res := make([][]int, 0)
    if root == nil {
        return res
    }
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    depth := 0
    for len(queue) > 0 {
        n := len(queue)
        level := make([]int, 0)
        for i := 0; i < n; i++ {
            level = append(level, queue[i].Val)
            if queue[i].Left != nil {
                queue = append(queue, queue[i].Left)
            }
            if queue[i].Right != nil {
                queue = append(queue, queue[i].Right)
            }
        }
        if depth % 2 == 1 {
            for j := 0; j < len(level)/2; j++ {
                level[j], level[len(level)-j-1] = level[len(level)-j-1], level[j]
            }
        }
        res = append(res, level)
        queue = queue[n:]
        depth++
    }
    return res
}
