/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode)  {
    if root == nil {
        return
    }
    var prev, x, y *TreeNode
    deque := make([]*TreeNode, 0)
    for len(deque) > 0 || root != nil {
        for root != nil {
            deque = append(deque, root)
            root = root.Left
        }
        root = deque[len(deque)-1]
        deque = deque[:len(deque)-1]
        if prev != nil && root.Val < prev.Val {
            y = root
            if x == nil {
                x = prev
            } else {
                break
            }
        }
        prev = root
        root = root.Right
    }
    x.Val, y.Val = y.Val, x.Val
}
