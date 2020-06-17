/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    depth := computeDepth(root)
    if depth == 0 {
        return 1
    }
    l, r := 1, int(math.Pow(2,float64(depth)))-1
    for l <= r {
        m := l + (r-l)/2
        if exists(root, depth, m) {
            l = m + 1
        } else {
            r = m - 1
        }
    }
    return int(math.Pow(2,float64(depth))) - 1 + l
}

func computeDepth(node *TreeNode) int {
    depth := 0
    for node.Left != nil {
        node = node.Left
        depth++
    }
    return depth
}

func exists(node *TreeNode, depth, pivot int) bool {
    l, r := 0, int(math.Pow(2,float64(depth)))-1
    for i := 0; i < depth; i++ {
        m := l + (r-l)/2
        if m >= pivot {
            node = node.Left
            r = m
        } else {
            node = node.Right
            l = m + 1
        }
    }
    return node != nil
}
