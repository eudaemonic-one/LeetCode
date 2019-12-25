/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    // firstly inorder traverse the binary tree
    stack := make([]*TreeNode, 0)
    sequence := make([]int, 0)
    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        if len(stack) == 0 {
            break
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        sequence = append(sequence, root.Val)
        root = root.Right
    }
    // then judge if the sequence is nondecreasing
    for i := 0; i < len(sequence)-1; i++ {
        if sequence[i] >= sequence[i+1] {
            return false
        }
    }
    return true
}
