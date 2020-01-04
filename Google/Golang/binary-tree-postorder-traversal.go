/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    stack := []*TreeNode{root}
    output := make([]int, 0)
    for true {
        for root != nil {
            if root.Right != nil {
                stack = append(stack, root.Right)
            }
            stack = append(stack, root)
            root = root.Left
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if root.Right != nil && len(stack) > 0 && root.Right == stack[len(stack)-1] {
            stack = stack[:len(stack)-1]
            stack = append(stack, root)
            root = root.Right
        } else if len(stack) > 0 {
            output = append(output, root.Val)
            root = nil
        }
        if len(stack) == 0 {
            break
        }
    }
    return output
}
